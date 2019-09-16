// Copyright 2019 VinyMeuh. All rights reserved.
// Use of the source code is governed by a MIT-style license that can be found in the LICENSE file.

package fileutils

import (
	"io"
	"os"
	"path/filepath"
)

// Copy copies from src to dst, creating the parents directories for dst if necessary
func Copy(src string, dst string) error {
	from, err := os.Open(src)
	if err != nil {
		return err
	}
	defer from.Close()

	to, err := os.OpenFile(dst, os.O_RDWR|os.O_CREATE|os.O_EXCL, os.ModePerm)
	if err != nil {
		if os.IsNotExist(err) {
			if err := os.MkdirAll(filepath.Dir(dst), os.ModePerm); err != nil {
				return err
			}
			to, err = os.OpenFile(dst, os.O_RDWR|os.O_CREATE|os.O_EXCL, os.ModePerm)
			if err != nil {
				return err
			}
		} else {
			return err
		}
	}
	defer to.Close()

	_, err = io.Copy(to, from)
	return err
}
