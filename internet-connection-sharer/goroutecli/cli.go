package cli

import (
	"github.com/spf13/cobra"
	"example.com/internet-connection-sharer/goroutecli/commands"

)

func NewRootCommand() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "gore",
		Short: "Go Router",
	}

	goroarCmd := commands.NewGoroarCommand()
	rootCmd.AddCommand(goroarCmd)

	return rootCmd
}
