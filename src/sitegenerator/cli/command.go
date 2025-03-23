package cli

import (
	"github.com/spf13/cobra"
)

func CreateRootCommand(appVersion string) *cobra.Command {
	return &cobra.Command{
		Use:     "sitegenerator",
		Version: appVersion,
		Short:   "Static site generator",
		Run: func(cmd *cobra.Command, args []string) {
			generate(cmd, args)
		},
	}
}
