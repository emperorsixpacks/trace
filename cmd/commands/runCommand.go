package commands

import (
	"github.com/emperorsixpacks/trace/engine"
	"github.com/spf13/cobra"
)

var RunCmd = &cobra.Command{
	Use:   "run",
	Short: "Run a flow",
	Long:  `Run a flow`,
	Run:   runCommand,
}

func runCommand(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		cmd.Help()
		return
	}
	fileName := args[0]
	engine.Runner(fileName)
}
