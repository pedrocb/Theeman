package cli

import (
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"math/rand"
	"os"
	"os/exec"
	"strings"
)

func get_themes() []string {
	files, _ := ioutil.ReadDir(themesPath)

	result := make([]string, len(files)-1)
	var themeName string
	i := 0
	for _, f := range files {
		themeName = RemoveExtension(f.Name())
		if strings.Compare(themeName, "current") != 0 {
			result[i] = themeName
			i++
		}
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

func getCurrentTheme() string {
	currentThemePath := ExpandDir(themesPath, "current")
	theme, _ := os.Readlink(currentThemePath)
	return DirName(theme)
}

func load_theme(path string) error {
	// i3
	var err error
	if _, err = os.Stat(ExpandDir(path, "i3/config")); err == nil {
		i3Path := viper.GetString("i3.path")
		if err = ForceSymlink(ExpandDir(path, "i3/config"), ExpandDir(i3Path, "config")); err == nil {
			err = exec.Command("i3-msg", "reload").Run()
		}
	}

	barPath := ExpandDir(path, "bar/bar.sh")
	if _, err = os.Stat(barPath); err == nil {
		exec.Command("sh", barPath).Run()
	}

	if _, err = os.Stat(ExpandDir(path, "rofi/config")); err == nil {
	}

	wallpapersPath := ExpandDir(path, "wallpaper")
	if _, err = os.Stat(wallpapersPath); err == nil {
		wallpapers, _ := ioutil.ReadDir(wallpapersPath)
		wallpaper := wallpapers[rand.Intn(len(wallpapers))].Name()
		err = exec.Command("feh", "--bg-scale", ExpandDir(wallpapersPath, wallpaper)).Run()
	}
	return err
}
