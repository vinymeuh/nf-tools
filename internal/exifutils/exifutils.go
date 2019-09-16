// Copyright 2019 VinyMeuh. All rights reserved.
// Use of the source code is governed by a MIT-style license that can be found in the LICENSE file.

package exifutils

import (
	"os"
	"time"

	"github.com/vinymeuh/nifuda/pkg/exif"
)

func ReadDateTime(path string) (time.Time, error) {
	osf, err := os.Open(path)
	if err != nil {
		return time.Time{}, err
	}
	defer osf.Close()

	f, err := exif.Read(osf)
	if err != nil {
		return time.Time{}, err
	}

	tags := f.Tags()
	datetime := tags["ifd0"]["DateTime"].Value().String() // TODO check key

	return time.Parse("2006:01:02 15:04:05", datetime)
}
