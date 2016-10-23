package cli

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var currentCommand = &cobra.Command{
	Use:   "current",
	Short: "Lists current theme",
	Run:   current,
}

func current(cmd *cobra.Command, args []string) {
	currentThemeFolder := configPath + "/themes/current"
	theme, _ := os.Readlink(currentThemeFolder)
	fmt.Println(DirName(theme))
}
