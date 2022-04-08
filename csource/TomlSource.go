package csource

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

type Toml struct {
	Path string
}

func (f Toml) Trans(IFace interface{}) error {
	file, err := os.Open(filepath.Clean(f.Path))
	if err != nil {
		return fmt.Errorf("GConf: cannot open toml file; err: %v", err)
	}

	if _, err = toml.NewDecoder(file).Decode(IFace); err != nil {
		return fmt.Errorf("GConf: cannot feed struct; err: %v", err)
	}

	return file.Close()
}
