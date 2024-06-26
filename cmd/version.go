package cmd

import (
	"github.com/luevano/mangal/meta"
	"github.com/spf13/cobra"
)

var versionArgs = struct {
	Short bool
}{}

func init() {
	subcommands = append(subcommands, versionCmd)

	versionCmd.Flags().BoolVarP(&versionArgs.Short, "short", "s", false, "Only show mangal version number")
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version information",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, _ []string) {
		if versionArgs.Short {
			cmd.Println(meta.Version)
			return
		}

		cmd.Println(meta.PrettyVersion())
	},
}
