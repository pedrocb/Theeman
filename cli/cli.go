package cli

import (
	"github.com/spf13/cobra"
)

var TheemanCommand = &cobra.Command{
	Use:   "theeman",
	Short: "theeman is a theme manager",
	Long:  "theeman is the main command to manage your desktop themes",
}

func AddCommands() {
	TheemanCommand.AddCommand(listCommand)
	TheemanCommand.AddCommand(currentCommand)
}

func Execute() {
	AddCommands()
	setupConfig()

	TheemanCommand.Execute()
}
