// Copyright 2019 VinyMeuh. All rights reserved.
// Use of the source code is governed by a MIT-style license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/vinymeuh/nf-tools/cmd/nf-import/internal/config"
	"github.com/vinymeuh/nf-tools/cmd/nf-import/internal/task"
	"github.com/vinymeuh/nf-tools/internal/yesno"
)

func main() {
	// read configuration file if exist
	confDir, err := os.UserConfigDir()
	if err == nil {
		conffile := confDir + string(filepath.Separator) + "nf-tools" + string(filepath.Separator) + "nf-import.yml"
		err := config.Read(conffile)
		if err == nil {
			fmt.Printf("Using configuration file '%s'\n", conffile)
		}
	}

	// parse commande line
	flag.Usage = func() {
		fmt.Printf("Usage: %s [parameters]\n\n", filepath.Base(os.Args[0]))
		fmt.Println("Parameters:")
		flag.PrintDefaults()
	}

	srcpath := flag.String("from", config.Conf.Path.From, "`directory` to import from.")
	tgtpath := flag.String("to", config.Conf.Path.To, "`directory` to import to.")
	flag.Parse()

	if *srcpath == "" || *tgtpath == "" {
		flag.Usage()
		os.Exit(1)
	}

	fmt.Printf("Source directory is '%s'\n", *srcpath)
	fmt.Printf("Target directory is '%s'\n", *tgtpath)

	// prepare tasks
	fmt.Printf("--->  Searching image files\n")
	tasks, err := task.Prepare(*srcpath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	ntasks := len(tasks)
	switch ntasks {
	case 0:
		fmt.Println("Source directory contains no image files")
		os.Exit(0)
	default:
		fmt.Printf("Source directory contains %d image files\n", len(tasks))
	}

	bySubdir := make(map[string]int, 0)
	for _, t := range tasks {
		bySubdir[t.TgtSubdirPath]++
	}
	for k, v := range bySubdir {
		fmt.Printf(" - %d will be imported into %s\n", v, k)
	}

	// wait confirmation to continue
	if !yesno.Ask("Continue [y/n]: ") {
		os.Exit(1)
	}

	// run tasks
	fmt.Printf("--->  Importing image files\n")
	ok := 0
	errs := make([]error, 0, len(tasks))
	for _, t := range tasks {
		err := t.Run(*tgtpath)
		if err == nil {
			ok++
		} else {
			errs = append(errs, err)
		}

	}

	fmt.Printf(" - %d image files imported successfully\n", ok)
	if len(errs) > 0 {
		fmt.Printf(" - %d image files encountered an error\n", len(errs))
		for _, e := range errs {
			fmt.Println(e)
		}
	}
}
