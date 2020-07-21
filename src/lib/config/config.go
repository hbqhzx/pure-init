package config

import (
	"fmt"
	"github.com/olebedev/config"
	"os"
)

const (
	CONFIG_FILE = "etc/config.yml"
)

var (
	Config *config.Config
)

func LoadCofig() error {
	cfgTemp, err := config.ParseYamlFile(CONFIG_FILE)
	if err != nil {
		fmt.Fprintf(os.Stderr, "parse config from %s failed: %s", CONFIG_FILE, err)
		return err
	}

	mode := "debug"
	if Config, err = cfgTemp.Get(mode); err != nil {
		return err
	}
	return nil
}
