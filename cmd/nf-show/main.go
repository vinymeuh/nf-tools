package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/vinymeuh/nifuda/pkg/exif"
)

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) != 1 {
		fmt.Fprintln(os.Stderr, "usage: nf-show EXIFFILE")
		os.Exit(2)
	}
	filepath := args[0]

	osf, err := os.Open(filepath)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}
	defer osf.Close()

	f, err := exif.Read(osf)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}

	for namespace, tags := range f.Tags() {
		for _, tag := range tags {
			fmt.Printf("%-6s   %-30s   %s\n", namespace, tag.Name(), tag.Value().String())
		}
	}
}
