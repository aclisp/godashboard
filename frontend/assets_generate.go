// +build ignore

package main

import (
	"log"
	"net/http"

	"github.com/shurcooL/vfsgen"
)

func main() {
	var fs http.FileSystem = http.Dir("html/")

	err := vfsgen.Generate(fs, vfsgen.Options{
		Filename:     "bundle/bundle.go",
		PackageName:  "bundle",
		BuildTags:    "!dev",
		VariableName: "Assets",
	})
	if err != nil {
		log.Fatalln(err)
	}
}
