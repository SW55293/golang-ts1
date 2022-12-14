/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "goCli",
	Short: "Basic hello world ",
	Long:  `Testing and Practicing the CLI application from Cobra`,

	Run: func(cmd *cobra.Command, args []string) {

		flagVal, err := cmd.Flags().GetBool("toggle")
		if err != nil {
			return
		}
		if flagVal {
			fmt.Println("Whats up Bitches!! == Not Set (False)")
			return
		}
		fmt.Println("Whats up Bitches!! == Set (True)")
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.goCli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Flip toggle value")
}
