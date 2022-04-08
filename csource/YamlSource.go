package csource

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type Yaml struct {
	Path string
}

func (f Yaml) Feed(Iface interface{}) error {
	file, err := os.Open(filepath.Clean(f.Path))
	if err != nil {
		return fmt.Errorf("GConf: cannot open yaml file; err: %v", err)
	}

	if err = yaml.NewDecoder(file).Decode(Iface); err != nil {
		return fmt.Errorf("GConf: cannot feed struct; err: %v", err)
	}

	return file.Close()
}
