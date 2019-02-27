package cli

import (
	"fmt"
	"github.com/spf13/cobra"
)

var listCommand = &cobra.Command{
	Use:   "list",
	Short: "List themes",
	Long:  "List themes",
	Run:   list,
}

func list(cmd *cobra.Command, args []string) {
	themes := get_themes()
	for _, f := range themes {
		fmt.Println(f)
	}
}
