// +build ignore

package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/hexiao04/radius/dictionary"
)

func main() {
	outputFile := flag.String("o", "-", "Output filename")
	flag.Parse()

	parser := &dictionary.Parser{
		Opener: &dictionary.FileSystemOpener{},
	}

	dict := &dictionary.Dictionary{}

	for _, filename := range flag.Args() {
		nextDict, err := parser.ParseFile(filename)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		dict, err = dictionary.Merge(dict, nextDict)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}

	var w io.Writer
	if *outputFile == "-" {
		w = os.Stdout
	} else {
		f, err := os.OpenFile(*outputFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		defer f.Close()
		w = f
	}

	fmt.Fprintln(w, "// Code generated by generate_main.go. DO NOT EDIT.")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "package debug")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, `import "layeh.com/radius/dictionary"`)
	fmt.Fprintln(w, "")
	fmt.Fprintf(w, "var IncludedDictionary = %#v\n", dict)
}
