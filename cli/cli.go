package cli

import (
	"github.com/spf13/cobra"
)

var theemanCommand = &cobra.Command{
	Use:   "theeman",
	Short: "theeman is a theme manager",
	Long:  "theeman is the main command to manage your desktop themes",
}

func addCommands() {
	theemanCommand.AddCommand(listCommand)
	theemanCommand.AddCommand(currentCommand)
	theemanCommand.AddCommand(applyCommand)
	theemanCommand.AddCommand(yamlCommand)
}

func Execute() {
	setupConfig()
	addCommands()
	theemanCommand.Execute()
}
