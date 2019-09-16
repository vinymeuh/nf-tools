// Copyright 2019 VinyMeuh. All rights reserved.
// Use of the source code is governed by a MIT-style license that can be found in the LICENSE file.

package task

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/karrick/godirwalk"
	"github.com/vinymeuh/nf-tools/internal/exifutils"
	"github.com/vinymeuh/nf-tools/internal/fileutils"
	"github.com/vinymeuh/nf-tools/internal/sidecar"
)

type Task struct {
	SrcFilePath   string
	TgtSubdirPath string
	TgtFilePath   string
}

func Prepare(imgpath string) ([]Task, error) {
	tasks := make([]Task, 0)

	err := godirwalk.Walk(imgpath, &godirwalk.Options{
		Callback: func(path string, de *godirwalk.Dirent) error {

			fileExt := filepath.Ext(path)
			switch strings.ToUpper(fileExt) {
			case ".JPG", ".JPEG", ".NEF":
			default:
				return nil
			}

			datetime, err := exifutils.ReadDateTime(path)
			if err != nil {
				return err
			}

			// Format YYYY/YYYY-MM-DD/EXT
			tgtSubdirPath := fmt.Sprintf("%d/%s/%s", datetime.Year(), datetime.Format("2006-01-02"), strings.ToUpper(fileExt[1:]))

			// Format YYYY-MM-DD_HHMISS_ORIGINAL_NAME.EXT
			prefix := datetime.Format("2006-01-02_150405")
			tgtFilePath := fmt.Sprintf("%s_%s", prefix, strings.TrimPrefix(filepath.Base(path), prefix+"_"))

			t := Task{
				SrcFilePath:   path,
				TgtSubdirPath: tgtSubdirPath,
				TgtFilePath:   tgtFilePath,
			}
			tasks = append(tasks, t)

			return nil
		},
	})
	return tasks, err
}

func (t Task) Run(tgtroot string) error {
	tgt := tgtroot + string(filepath.Separator) + t.TgtSubdirPath + string(filepath.Separator) + t.TgtFilePath
	err := fileutils.Copy(t.SrcFilePath, tgt)
	if err != nil {
		return err
	}

	srcDOP := sidecar.DOPPath(t.SrcFilePath)
	_, err = os.Stat(srcDOP)
	switch err {
	case nil:
		tgtDOP := sidecar.DOPPath(tgt)
		err := fileutils.Copy(srcDOP, tgtDOP)
		if err != nil {
			return err
		}
	default:
		if !os.IsNotExist(err) {
			return err
		}
	}

	srcXMP := sidecar.XMPPath(t.SrcFilePath)
	_, err = os.Stat(srcXMP)
	switch err {
	case nil:
		tgtXMP := sidecar.XMPPath(tgt)
		err := fileutils.Copy(srcXMP, tgtXMP)
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
