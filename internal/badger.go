package internal

import (
	"os"

	badger "github.com/dgraph-io/badger/v3"
	"github.com/pkg/errors"
)

// Badger - needed infos for one badger db
type Badger struct {
	Path    string // Filepath to the data file.
	Options badger.Options
	db      *badger.DB
}

// Open - open and initialize the store.
func (s *Badger) Open() error {
	db, err := badger.Open(s.Options)
	if err != nil {
		return errors.Wrap(err, "Open")
	}
	s.db = db

	return nil
}

// Close - close and shutdown the store.
func (s *Badger) Close() error {
	return s.db.Close()
}

// GetStoreRo - return readonly store
func GetStoreRo(filePath string) (*Badger, error) {
	store, err := getStore(filePath, true)
	if err != nil {
		return nil, err
	}
	return store, nil
}

// GetStoreRw - return read/write store
func GetStoreRw(filePath string) (*Badger, error) {
	store, err := getStore(filePath, false)
	if err != nil {
		return nil, err
	}
	return store, nil
}

// getStore - return a store
func getStore(filePath string, ro bool) (*Badger, error) {

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return nil, err
	}
	opts := badger.DefaultOptions(filePath)
	// opts.WithCompression(0) // 2: zstd, 0: off
	opts.WithCompression(2) // 2: zstd, 0: off
	opts.WithZSTDCompressionLevel(2)
	opts.Logger = nil
	if ro {
		opts.ReadOnly = true
	}

	store := Badger{
		Path:    filePath,
		Options: opts,
	}
	err := store.Open()
	if err != nil {
		return nil, err
	}
	return &store, nil
}

// GetValue - return a single value
func (s *Badger) GetValue(key []byte) ([]byte, error) {
	var value []byte
	err := s.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get(key)
		if err != nil {
			switch err {
			case badger.ErrKeyNotFound:
				return badger.ErrKeyNotFound
			default:
				return err
			}
		}
		value, err = item.ValueCopy(value)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return value, nil
}

// SetValue - insert a single key with value
func (s *Badger) SetValue(key, value []byte) error {

	// Start a writable transaction.
	txn := s.db.NewTransaction(true)
	defer txn.Discard()

	// Use the transaction...
	err := txn.Set(key, value)
	if err != nil {
		return err
	}

	// Commit the transaction and check for error.
	if err := txn.Commit(); err != nil {
		return err
	}
	return nil
}

// DeleteValue - delete a single value
func (s *Badger) DeleteValue(key []byte) error {

	// Start a writable transaction.
	txn := s.db.NewTransaction(true)
	defer txn.Discard()

	// Use the transaction...
	err := txn.Delete(key)
	if err != nil {
		return err
	}

	// Commit the transaction and check for error.
	if err := txn.Commit(); err != nil {
		return err
	}
	return nil
}

// GetValues - return all key,value pairs of a database as map
func (s *Badger) GetValues() (map[string][]byte, error) {
	values := make(map[string][]byte)

	err := s.db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchSize = 10
		it := txn.NewIterator(opts)
		defer it.Close()
		for it.Rewind(); it.Valid(); it.Next() {
			item := it.Item()
			k := item.Key()
			err := item.Value(func(v []byte) error {
				value, err := item.ValueCopy(v)
				if err != nil {
					return err
				}
				values[string(k)] = value
				return nil
			})
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return values, nil
}

// FindMatchingKey - find correct key depending on timestamp and direction
func (s *Badger) FindMatchingKey(ts string, direction int32) ([]byte, error) {
	var mKey []byte
	var it *badger.Iterator

	err := s.db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchValues = false
		if direction <= 0 {
			opts.Reverse = true
		}
		it = txn.NewIterator(opts)
		defer it.Close()

		it.Seek([]byte(ts))
		if direction != 0 {
			it.Next()
		}
		if !it.Valid() {
			return badger.ErrKeyNotFound
		}
		item := it.Item()
		mKey = item.Key()
		return nil
	})
	if err != nil {
		return nil, err
	}
	return mKey, nil
}
