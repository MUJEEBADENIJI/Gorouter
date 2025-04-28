import (
	"github.com/spf13/cobra"
)

func NewGoroarCommand() *cobra.Command {
	goroarCmd := &cobra.Command{
		Use:   "goroar",
		Short: "Start the Go Router",
		Run: func(cmd *cobra.Command, args []string) {
			// Call the router's start function
			// For now, let's just print a message
			fmt.Println("Starting the Go Router...")
		},
	}
	return goroarCmd
}
