package pkg

import (
	"fmt"
	"github.com/ghodss/yaml"
	"github.com/wonderivan/logger"
	_ "github.com/wonderivan/logger"
)

func Yj(config string) error {
	bt, err := GetByteStream(config)
	if err != nil {
		logger.Error(err)
		return err
	}
	if !HasJSONPrefix(bt) {
		bt, err = yaml.YAMLToJSON(bt)
		if err != nil {
			logger.Error(err)
			return err
		}
		fmt.Println(string(bt))
	} else {
		bt, err = yaml.JSONToYAML(bt)
		if err != nil {
			logger.Error(err)
			return err
		}
		fmt.Println(string(bt))
	}
	return nil
}
