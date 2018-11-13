package cmd

import (
	"github.com/justmiles/ssm-parameter-store/lib"
	"github.com/spf13/cobra"
)

var (
	noInput bool
)

func init() {
	rootCmd.AddCommand(pushCmd)
	pushCmd.PersistentFlags().BoolVar(&noInput, "no-input", false, "not not prompt for input")

}

// process the list command
var pushCmd = &cobra.Command{
	Use:   "push",
	Short: "push SSM Parameters",
	Run: func(cmd *cobra.Command, args []string) {
		ssmParameterStore.CMDPush(paths, format, directory, noInput)
	},
}
