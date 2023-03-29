package main

import (
	"errors"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	APIKey  string `yaml:"apiKey"`
	Weather struct {
		Interval  string     `yaml:"interval"`
		Unit      string     `yaml:"unit"`
		Locations []Location `yaml:"locations"`
	} `yaml:"weather"`
}

type Location struct {
	Name string  `yaml:"name"`
	Lat  float64 `yaml:"lat"`
	Lon  float64 `yaml:"lon"`
	Main bool    `yaml:"main"`
}

func (c *Config) parseConfig() error {
	yamlFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		return err
	}

	return nil
}

func (c *Config) isValid() error {
	cMain := 0
	for _, l := range c.Weather.Locations {
		if l.Main {
			cMain++
		}
	}
	if cMain < 1 {
		return errors.New("You must specify a location with `main: true`")
	}
	if cMain > 1 {
		return errors.New("Only a single location can have `main: true`")
	}
	return nil
}

func (c *Config) getLocations() []Location {
	return c.Weather.Locations
}

func (c *Config) getMainLocation() Location {
	for _, l := range c.Weather.Locations {
		if l.Main {
			return l
		}
	}
	return Location{}
}

func (c *Config) getOtherLocations() []Location {
	var ls []Location
	for _, l := range c.Weather.Locations {
		if !l.Main {
			ls = append(ls, l)
		}
	}
	return ls
}
