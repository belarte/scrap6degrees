package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// moviesCmd represents the movies command.
var moviesCmd = &cobra.Command{
	Use:   "movies",
	Short: "Get movies details",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("movies called")
	},
}

func init() {
	getCmd.AddCommand(moviesCmd)
}
