<<<<<<< HEAD
package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

func main() {
=======
package cli

import (
	"github.com/spf13/cobra"
	"example.com/internet-connection-sharer/goroutecli/commands"

)

func NewRootCommand() *cobra.Command {
>>>>>>> goleft
	rootCmd := &cobra.Command{
		Use:   "gore",
		Short: "Go Router",
	}

<<<<<<< HEAD
	startCmd := &cobra.Command{
		Use:   "goroar",
		Short: "Start the Go Router",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Starting the Go Router...")
			// Code to start the router goes here
		},
	}
	rootCmd.AddCommand(startCmd)

	stopCmd := &cobra.Command{
		Use:   "gohush",
		Short: "Stop the Go Router",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Stopping the Go Router...")
			// Code to stop the router goes here
		},
	}
	rootCmd.AddCommand(stopCmd)

	statusCmd := &cobra.Command{
		Use:   "gostat",
		Short: "Check the status of the Go Router",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Checking status...")
			// Code to check the status goes here
		},
	}
	rootCmd.AddCommand(statusCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
=======
	goroarCmd := commands.NewGoroarCommand()
	rootCmd.AddCommand(goroarCmd)

	return rootCmd
>>>>>>> goleft
}
