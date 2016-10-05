package main

import (
	"./consts"
	"./utils"
	"fmt"
	"github.com/mitchellh/cli"
	"io/ioutil"
	"os"
)

type ListCommand struct {
	Ui cli.Ui
}

func (c ListCommand) Run(_ []string) int {
	_, err := os.Stat(consts.CONFIG_PATH)
	if err != nil {
		c.Ui.Output("Creating config folder at ~/.theeman...")
		utils.CopyDir(consts.CONFIG_FOLDER, consts.CONFIG_PATH)
	}

	files, err := ioutil.ReadDir(consts.THEMES_PATH)
	for _, f := range files {
		fmt.Println(utils.RemoveExtension(f.Name()))
	}
	return 0
}

func (c ListCommand) Help() string {
	return "HELP"
}

func (c ListCommand) Synopsis() string {
	return "Run as list"
}
