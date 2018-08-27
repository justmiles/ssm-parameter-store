package cmd

import (
	"github.com/justmiles/ssm-parameter-store/lib"
	"github.com/spf13/cobra"
)

var ()

func init() {
	rootCmd.AddCommand(pushCmd)
}

// process the list command
var pushCmd = &cobra.Command{
	Use:   "push",
	Short: "push SSM Parameters",
	Run: func(cmd *cobra.Command, args []string) {
		ssmParameterStore.CMDPush(paths, format, directory)
	},
}
