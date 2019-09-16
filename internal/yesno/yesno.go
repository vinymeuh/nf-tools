// Copyright 2019 VinyMeuh. All rights reserved.
// Use of the source code is governed by a MIT-style license that can be found in the LICENSE file.

package yesno

import (
	"bufio"
	"os"
	"strings"

	"github.com/gookit/color"
)

func Ask(msg string) bool {
	reader := bufio.NewReader(os.Stdin)
	for {
		color.Bold.Print(msg)
		response, _ := reader.ReadString('\n')
		response = strings.TrimSuffix(response, "\n")
		switch response {
		case "y":
			return true
		case "n":
			return false
		}
	}
}
