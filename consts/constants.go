package consts

import (
	"../utils"
	"github.com/mitchellh/go-homedir"
)

var CONFIG_FOLDER = ".theeman"
var CONFIG_PATH, _ = homedir.Expand("~/" + CONFIG_FOLDER)
var THEMES_PATH = utils.ExpandDir(CONFIG_PATH, "themes")
