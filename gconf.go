package gConf

import (
	"fmt"
)

type ConfSource interface {
	Trans(IFace interface{}) error
}

type Config struct {
	Source ConfSource
	To     interface{}
}

func New() *Config {
	return &Config{}
}

func (c *Config) AddSource(f ConfSource) *Config {
	c.Source = f
	return c
}

func (c *Config) AddTo(f interface{}) *Config {
	c.To = f
	return c
}

func (c *Config) Trans() error {
	if err := (c.Source).Trans(c.To); err != nil {
		return fmt.Errorf("GConf failed trans %v", err)
	}
	return nil

}
