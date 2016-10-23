package cli

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var applyCommand = &cobra.Command{
	Use:   "apply",
	Short: "Applies a theme",
	RunE:  apply,
}

func apply(cmd *cobra.Command, args []string) (e error) {
	themesPath := ExpandDir(configPath, "themes")
	if len(args) < 1 {
		return errors.New("A theme must be specified")
	}
	if theme_exists(args[0]) {
		os.Remove(ExpandDir(themesPath, "current"))
		err := os.Symlink(ExpandDir(themesPath, args[0]), ExpandDir(themesPath, "current"))
	}
	return nil
}
