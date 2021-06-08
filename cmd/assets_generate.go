//
// +build ignore

package main

import (
	"log"
	"net/http"
	"path"
	"path/filepath"
	"runtime"

	"github.com/shurcooL/vfsgen"
)

var (
	_, b, _, _  = runtime.Caller(0)
	projectPath = filepath.Dir(path.Join(path.Dir(b)))
)

func main() {

	// var assets http.FileSystem = http.Dir("/home/ulli/lang/go/src/github.com/ulranh/saphistory/client/dist")

	// adapt for docker build
	var assets http.FileSystem = http.Dir(filepath.Join(projectPath, "client", "dist"))

	err := vfsgen.Generate(assets, vfsgen.Options{
		PackageName:  "cmd",
		BuildTags:    "!dev",
		VariableName: "assets",
	})
	if err != nil {
		log.Fatalln(err)
	}
}
