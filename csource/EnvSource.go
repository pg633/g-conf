package csource

import (
	"fmt"

	"github.com/golobby/env/v2"
)

type Env struct{}

func (f Env) Feed(Iface interface{}) error {
	if err := env.Feed(Iface); err != nil {
		return fmt.Errorf("GConf: cannot feed struct; err: %v", err)
	}

	return nil
}
