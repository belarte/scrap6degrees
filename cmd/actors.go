package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// actorsCmd represents the actors command.
var actorsCmd = &cobra.Command{
	Use:   "actors",
	Short: "Get actors details",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("actors called")
	},
}

func init() {
	getCmd.AddCommand(actorsCmd)
}
