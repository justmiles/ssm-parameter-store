package cmd

import (
	"github.com/justmiles/ssm-parameter-store/lib"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(diffCmd)
}

// process the list command
var diffCmd = &cobra.Command{
	Use:   "diff",
	Short: "diff SSM Parameters with those on disk",
	Run: func(cmd *cobra.Command, args []string) {
		ssmParameterStore.CMDDiff(paths, format, directory)
	},
}
