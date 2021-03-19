package goget

import (
	"gopkg.in/gcfg.v1"
)

type Config struct {
	Package map[string]*Package
}

func PackagesFromFile(fn string) (Packages, error) {
	var ini Config

	if err := gcfg.ReadFileInto(&ini, fn); err != nil {
		return nil, err
	}

	return ini.Package, nil
}
