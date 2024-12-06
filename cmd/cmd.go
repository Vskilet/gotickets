package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func main() {

	var version = &cobra.Command{
		Use:   "version",
		Short: "Get a version api of our tool",
		Long: `More fun with a long description that point to the masterclass 
				website of the universe `,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("v0.0.1")
		},
	}
	cmd := &cobra.Command{}
	cmd.AddCommand(version)
	cmd.Execute()
}
