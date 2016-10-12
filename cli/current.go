package cli

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var currentCommand = &cobra.Command{
	Use:   "current",
	Short: "Lists current theme",
	Run:   current,
}

func current(cmd *cobra.Command, args []string) {
	fmt.Println(viper.GetString("themes.current"))
}
