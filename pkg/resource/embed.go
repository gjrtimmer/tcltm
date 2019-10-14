// +build ignore

package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/format"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

const (
	blob = "resources.go"
)

var (
	directory string
)

func formatByteSlice(sl []byte) string {
	builder := strings.Builder{}
	for _, v := range sl {
		builder.WriteString(fmt.Sprintf("%d,", int(v)))
	}
	return builder.String()
}

func init() {
	flag.StringVar(&directory, "dir", "", "Resource Directory")
}

func main() {
	// Parse Flags
	flag.Parse()

	log.Printf("Embedding resources: %s\n", directory)

	if _, err := os.Stat(directory); os.IsNotExist(err) {
		log.Fatal("Resources directory does not exists")
	}

	resources := make(map[string][]byte)
	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Println("Error :", err)
			return err
		}

		relativePath := filepath.ToSlash(strings.TrimPrefix(path, directory))
		if info.IsDir() {
			// NOOP
			return nil
		} else {
			log.Printf("Embedding %s\n", path)
			b, err := ioutil.ReadFile(path)
			if err != nil {
				log.Printf("Error reading %s: %s", path, err)
				return err
			}
			resources[relativePath] = b
		}
		return nil
	})

	if err != nil {
		log.Fatal("Error walking through resources directory:", err)
	}

	f, err := os.Create(blob)
	if err != nil {
		log.Fatal("Error creating blob file:", err)
	}
	defer f.Close()

	builder := &bytes.Buffer{}

	t, err := template.New("_template.go").Funcs(template.FuncMap{
		"conv": formatByteSlice,
	}).ParseFiles("_template.go")
	if err != nil {
		log.Fatal("Error creating template", err)
	}

	err = t.Execute(builder, resources)
	if err != nil {
		log.Fatal("Error executing template", err)
	}

	data, err := format.Source(builder.Bytes())
	if err != nil {
		log.Fatal("Error formatting generated code", err)
	}
	err = ioutil.WriteFile(blob, data, os.ModePerm)
	if err != nil {
		log.Fatal("Error writing blob file", err)
	}

	log.Println("Embedding resources done")
}

// EOF
