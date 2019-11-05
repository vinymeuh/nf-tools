// Copyright 2019 VinyMeuh. All rights reserved.
// Use of the source code is governed by a MIT-style license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"reflect"

	"github.com/vinymeuh/nifuda"
)

func main() {
	flag.Usage = func() {
		fmt.Printf("Usage: %s EXIFFILE\n", filepath.Base(os.Args[0]))
		flag.PrintDefaults()
	}

	// parse commande line
	flag.Parse()
	args := flag.Args()

	if len(args) != 1 {
		flag.Usage()
		os.Exit(1)
	}

	err := task(args[0], os.Stdout)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}
}

func task(path string, out io.Writer) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	x, err := nifuda.Read(f)
	if err != nil {
		return err
	}

	v := reflect.ValueOf(x.Image)
	vt := v.Type()
	for i := 0; i < vt.NumField(); i++ {
		fmt.Printf("Image.%-30s   %v\n", vt.Field(i).Name, v.Field(i).Interface())
	}

	v = reflect.ValueOf(x.Photo)
	vt = v.Type()
	for i := 0; i < vt.NumField(); i++ {
		fmt.Printf("Photo.%-30s   %v\n", vt.Field(i).Name, v.Field(i).Interface())
	}

	v = reflect.ValueOf(x.Gps)
	vt = v.Type()
	for i := 0; i < vt.NumField(); i++ {
		fmt.Printf("GPS.%-30s     %v\n", vt.Field(i).Name, v.Field(i).Interface())
	}

	return nil
}
