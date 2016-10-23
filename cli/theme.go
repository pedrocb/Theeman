package cli

import (
	"io/ioutil"
	"strings"
)

func get_themes() []string {
	files, _ := ioutil.ReadDir(configPath + "/themes")

	result := make([]string, len(files))
	for i, f := range files {
		result[i] = RemoveExtension(f.Name())
	}
	return result
}

func theme_exists(theme string) bool {
	for _, f := range get_themes() {
		if strings.Compare(f, theme) == 0 {
			return true
		}
	}
	return false
}
