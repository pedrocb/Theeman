package cli

import (
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
)

var listCommand = &cobra.Command{
	Use:   "list",
	Short: "List themes",
	Long:  "List themes",
	Run:   list,
}

func list(cmd *cobra.Command, args []string) {
	files, _ := ioutil.ReadDir(theemanPath + "/themes")
	for _, f := range files {
		fmt.Println(RemoveExtension(f.Name()))
	}
}
