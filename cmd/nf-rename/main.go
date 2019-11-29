// Copyright 2019 VinyMeuh. All rights reserved.
// Use of the source code is governed by a MIT-style license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/vinymeuh/nf-tools/internal/fileutils"
	"github.com/vinymeuh/nf-tools/internal/sidecar"
	"github.com/vinymeuh/nifuda"
)

func main() {
	flag.Usage = func() {
		fmt.Printf("Usage: %s DIRECTORY\n", filepath.Base(os.Args[0]))
		flag.PrintDefaults()
	}
	flag.Parse()
	args := flag.Args()

	if len(args) != 1 {
		flag.Usage()
		os.Exit(1)
	}
	rootPath := args[0]

	files, err := ioutil.ReadDir(rootPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	var errorsOccured int
	for _, file := range files {
		path := fmt.Sprintf("%s/%s", rootPath, file.Name())
		pathExt := filepath.Ext(path)
		switch strings.ToUpper(pathExt) {
		case ".JPG", ".JPEG", ".NEF":
		default:
			continue
		}

		// copy image file to new path
		newPath, err := newNameForImageFile(path)
		if err != nil {
			fmt.Println(err)
			continue
		}

		err = fileutils.Copy(path, newPath)
		if err != nil {
			fmt.Println(err)
			errorsOccured++
			continue
		}

		// sidecar files
		srcDOP := sidecar.DOPPath(path)
		tgtDOP := sidecar.DOPPath(newPath)
		err = copySidecarFile(srcDOP, tgtDOP)
		if err != nil {
			fmt.Println(err)
			errorsOccured++
			continue
		}

		srcXMP := sidecar.XMPPath(path)
		tgtXMP := sidecar.XMPPath(newPath)
		err = copySidecarFile(srcXMP, tgtXMP)
		if err != nil {
			fmt.Println(err)
			errorsOccured++
			continue
		}
	}

	if errorsOccured > 0 {
		os.Exit(1)
	}
}

func newNameForImageFile(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()

	x, err := nifuda.Read(f)
	if err != nil {
		return "", err
	}

	datetime, err := time.Parse("2006:01:02 15:04:05", x.Image.DateTime)
	if err != nil {
		return "", err
	}

	return newName(path, datetime), nil
}

// new path YYYY-MM-DD/YYYY-MM-DD_HHMISS_ORIGINAL_NAME.EXT
func newName(path string, datetime time.Time) string {
	pathDir := filepath.Dir(path)
	pathExt := filepath.Ext(path) // file extension including the dot
	pathBase := filepath.Base(path)[0 : len(filepath.Base(path))-len(pathExt)]

	newSubDir := datetime.Format("2006-01-02")
	datePrefix := datetime.Format("2006-01-02_150405")

	pathPostfix := strings.TrimPrefix(pathBase, datePrefix+"_")

	if pathPostfix == datePrefix {
		return fmt.Sprintf("%s/%s/%s%s", pathDir, newSubDir, datePrefix, pathExt)
	}
	return fmt.Sprintf("%s/%s/%s_%s%s", pathDir, newSubDir, datePrefix, pathPostfix, pathExt)

}

func copySidecarFile(path string, newPath string) error {
	_, err := os.Stat(path)
	switch err {
	case nil:
		err := fileutils.Copy(path, newPath)
		if err != nil {
			return err
		}
	default:
		if !os.IsNotExist(err) {
			return err
		}
	}
	return nil
}
