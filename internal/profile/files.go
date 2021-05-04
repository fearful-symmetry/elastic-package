// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package profile

import (
	"io/ioutil"

	"github.com/pkg/errors"
)

// NewConfig is a generic function type to return a new Managed config
type NewConfig = func(profileName string, profilePath string) (*simpleFile, error)

// simpleFile defines a file that's managed by the profile system
// and doesn't require any  rendering
type simpleFile struct {
	Name string
	Path string
	Body string
}

// ConfigfilesDiffer checks to see if a local configItem differs from the one it knows.
func (cfg simpleFile) ConfigfilesDiffer() (bool, error) {
	changes, err := ioutil.ReadFile(cfg.Path)
	if err != nil {
		return false, errors.Wrapf(err, "error reading %s", KibanaConfigFile)
	}
	if string(changes) != cfg.Body {
		return true, nil
	}
	return false, nil
}

// WriteConfig writes the config item
func (cfg simpleFile) WriteConfig() error {
	err := ioutil.WriteFile(cfg.Path, []byte(cfg.Body), 0644)
	if err != nil {
		return errors.Wrapf(err, "writing file failed (path: %s)", cfg.Path)
	}
	return nil
}

// ReadConfig reads the config item, overwriting whatever exists in the fileBody.
func (cfg *simpleFile) ReadConfig() error {
	body, err := ioutil.ReadFile(cfg.Path)
	if err != nil {
		return errors.Wrapf(err, "reading filed failed (path: %s)", cfg.Path)
	}
	cfg.Body = string(body)
	return nil
}
