package goget

import (
	"gopkg.in/gcfg.v1"
)

type config struct {
	Package Packages
}

type Config struct {
	config   config
	Filename string
	Renderer Renderer
}

func NewFromFile(fn string, renderer Renderer) (*Config, error) {
	c := &Config{
		Filename: fn,
		Renderer: renderer,
	}

	if err := c.Load(); err != nil {
		return nil, err
	}

	return c, nil
}

// Load()
func (ini *config) load(fn string) error {
	if err := gcfg.ReadFileInto(ini, fn); err != nil {
		return err
	}

	if err := ini.Package.SetDefaults(); err != nil {
		return err
	}

	return nil
}

func (c *Config) Load() error {
	ini := c.config
	return ini.load(c.Filename)
}

// Reload()
func (c *Config) Reload() error {
	ini := config{}
	if err := ini.load(c.Filename); err != nil {
		return err
	}

	c.config = ini
	return nil
}

// Packages
func (c *Config) Packages() Packages {
	return c.config.Package
}

func PackagesFromFile(fn string) (Packages, error) {
	c, err := NewFromFile(fn, nil)
	if err != nil {
		return nil, err
	}

	return c.Packages(), nil
}
