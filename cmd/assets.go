// +build dev
//go:generate go run assets_generate.go

package cmd

import "net/http"

var (
	_, b, _, _ = runtime.Caller(0)
	projectPath   = filepath.Dir(path.Join(path.Dir(b)))
)

// var assets http.FileSystem = http.Dir("/home/ulli/lang/go/src/github.com/ulranh/saphistory/client/dist")
var assets http.FileSystem = http.Dir(filepath.Join(projectPath,"client","dist")
