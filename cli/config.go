package cli

import (
	"errors"
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
	"os"
)

var configPath string
var themesPath string

func setupConfig() error {
	configPath, _ = homedir.Expand("~/.theeman")
	themesPath = ExpandDir(configPath, "themes")
	_, err := os.Stat(configPath)
	if err != nil {
		return errors.New(fmt.Sprintf("Missing config folder at %s", configPath))
	}

	viper.SetConfigType("toml")
	viper.SetConfigName("config")

	viper.AddConfigPath(configPath)

	err = viper.ReadInConfig()
	return err
}
