// Copyright 2019 VinyMeuh. All rights reserved.
// Use of the source code is governed by a MIT-style license that can be found in the LICENSE file.

package sidecar

import (
	"path/filepath"
	"strings"
)

// DOPPath returns the path of a sidecar DOP file eventually existing
func DOPPath(path string) string {
	return path + ".dop"
}

// XMPPath returns the path of a sidecar XMP file eventually existing
func XMPPath(path string) string {
	return strings.TrimSuffix(path, filepath.Ext(path)) + ".xmp"
}
