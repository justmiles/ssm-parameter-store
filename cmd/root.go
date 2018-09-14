package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	directory, format string
	paths             []string
)

func init() {
	dir, err := os.Getwd()
	check(err)
	rootCmd.PersistentFlags().StringVarP(&format, "format", "f", "yaml", "format type")
	rootCmd.PersistentFlags().StringVarP(&directory, "directory", "d", dir, "output directory")
	rootCmd.PersistentFlags().StringSliceVarP(&paths, "path", "p", []string{"/"}, "path")
}

// Configure the root command
var rootCmd = &cobra.Command{
	Use:     "ssm-parameter-store",
	Short:   "Sync SSM Parameter Store",
	Version: "v0.0.4",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

// Execute validates input the Cobra CLI
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// Log errors if exist and exit
func check(err error) {
	if err != nil {
		fmt.Printf("ERROR\t%s", err.Error())
		os.Exit(1)
	}
}
