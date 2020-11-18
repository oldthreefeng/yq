package pkg

import (
	"fmt"
	"github.com/ghodss/yaml"
	"github.com/wonderivan/logger"
	_ "github.com/wonderivan/logger"
)

// format when json is true . yaml is false
func Yj(config string, format bool) error {
	bt, err := GetByteStream(config)
	if err != nil {
		logger.Error(err)
		return err
	}
	// yaml byte , json out
	if !HasJSONPrefix(bt) && format {
		bt, err = yaml.YAMLToJSON(bt)
		if err != nil {
			logger.Error(err)
		}
		fmt.Println(string(bt))
		// json byte , yaml out
	} else if HasJSONPrefix(bt) && !format {
		bt, err = yaml.JSONToYAML(bt)
		if err != nil {
			logger.Error(err)
		}
		fmt.Println(string(bt))
	} else {
		fmt.Println(string(bt))
	}
	return nil
}
