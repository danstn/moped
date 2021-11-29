package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// helloCmd represents the hello command
var helloCmd = &cobra.Command{
	Use:   "hello",
	Short: "Says hello.",
	Long:  `Not much else.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello.")
	},
}

func init() {
	rootCmd.AddCommand(helloCmd)
}
