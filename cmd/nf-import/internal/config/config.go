// Copyright 2019 VinyMeuh. All rights reserved.
// Use of the source code is governed by a MIT-style license that can be found in the LICENSE file.

package config

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

var Conf conf

type conf struct {
	Path confPath `yaml:"path"`
}

type confPath struct {
	From string `yaml:"from"`
	To   string `yaml:"to"`
}

func Read(path string) error {
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		return fmt.Errorf("error reading configuration file: %w", err)
	}

	err = yaml.Unmarshal(yamlFile, &Conf)
	if err != nil {
		return fmt.Errorf("error unmarshaling yaml from configuration file: %w", err)
	}

	return nil
}
