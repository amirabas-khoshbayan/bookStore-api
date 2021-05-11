package config

import (
	"errors"
	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v3"
	"os"
	"strings"
)

var (
	cfg *Config
   //ErrUnknownFileExtension is returned by the Parse function
   // when the file extension in not allowed for configuration
	ErrUnknownFileExtension = errors.New("unknown file extension")
)

//Parse parses config file into Config
func Parse(path string, cfg *Config) error  {
	switch fileExtension(path) {
	case "yaml":
		return parseYAML(path,cfg)
	default:
		return ErrUnknownFileExtension
	}
}

func parseYAML(path string,cfg *Config) (err error) {
   cfgPath := "./build/config/config.yaml"

  file,err := os.Open(cfgPath)
	if err != nil {
		return err
	}

	defer func() {
		if e := file.Close(); err == nil{
			err = e
		}

	}()
  decoder :=  yaml.NewDecoder(file)
	if err := decoder.Decode(cfg); err != nil{
		return err
	}
    return nil
}

func ReadEnv(cfg *Config) error  {
   return envconfig.Process("",cfg)
}

func SetConfig(c *Config)  {
	cfg = c
}
//fileExtension returns extension of file
func fileExtension(path string) string  {
	s := strings.Split(path,".")
	return s[len(s)-1]
}


