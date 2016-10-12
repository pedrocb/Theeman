package cli

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
	"os"
)

var theemanPath, _ = homedir.Expand("~/.theeman")

func setupConfig() {
	_, err := os.Stat(theemanPath)
	if err != nil {
		fmt.Println("Creating config folder at ~/.theeman ...")
		CopyDir(".theeman", theemanPath)
	}

	setupViperDefaults()
}

func setupViperDefaults() {
	viper.SetConfigName("config")
	viper.AddConfigPath(theemanPath)
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()

	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}
