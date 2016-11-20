package cli

import (
	"errors"
	"github.com/spf13/cobra"
)

var applyCommand = &cobra.Command{
	Use:   "apply",
	Short: "Applies a theme",
	RunE:  apply,
}

func apply(cmd *cobra.Command, args []string) (e error) {
	if len(args) < 1 {
		return errors.New("A theme must be specified")
	}
	if theme_exists(args[0]) {
		var currentThemePath = ExpandDir(themesPath, "current")
		load_theme(ExpandDir(themesPath, args[0]))
		err := ForceSymlink(ExpandDir(themesPath, args[0]), currentThemePath)
		return err
	}
	return nil
}
