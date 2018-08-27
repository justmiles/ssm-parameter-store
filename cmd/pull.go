package cmd

import (
	"github.com/justmiles/ssm-parameter-store/lib"
	"github.com/spf13/cobra"
)

var ()

func init() {
	rootCmd.AddCommand(pullCmd)
}

// process the list command
var pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "pull SSM Parameters",
	Run: func(cmd *cobra.Command, args []string) {
		ssmParameterStore.CMDPull(paths, format, directory)
	},
}
