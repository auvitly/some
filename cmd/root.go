/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

//go:generate env

var INTERFACE1 interface{ String() error }
var INTERFACE2 = interface{ String() error }(nil)
var INTERFACE3 = INTERFACE2
var INT1 int
var INT2 = 0
var CHAN chan struct{}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "some",
	Short: "A brief description of your application",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
	Run: func(cmd *cobra.Command, args []string) {
		return
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.CompletionOptions.DisableDescriptions = true

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}
