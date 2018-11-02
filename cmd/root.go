package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use: "Golang Playground",
	Short: "Various Examples and Snippets",
	Long: `A consistent and thorough go snippets built with 
	- Cobra
	`,
	Run: func(cmd *cobra.Command, args []string) {
		playground.Run()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}