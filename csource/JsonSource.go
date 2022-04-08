package csource

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Json struct {
	Path string
}

func (f *Json) Trans(IFace interface{}) error {
	file, err := os.Open(filepath.Clean(f.Path))
	if err != nil {
		return fmt.Errorf("GConf: cannot open json file; err: %v", err)
	}

	if err = json.NewDecoder(file).Decode(IFace); err != nil {
		return fmt.Errorf("GConf: cannot feed struct; err: %v", err)
	}

	return file.Close()
}
