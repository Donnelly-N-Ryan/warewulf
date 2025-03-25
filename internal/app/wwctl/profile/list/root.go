package list

import (
	"github.com/spf13/cobra"
	"github.com/warewulf/warewulf/internal/app/wwctl/completions"
)

type variables struct {
	showAll  bool
	showYaml bool
	showJson bool
}

// GetRootCommand returns the root cobra.Command for the application.
func GetCommand() *cobra.Command {
	vars := variables{}
	baseCmd := &cobra.Command{
		DisableFlagsInUseLine: true,
		Use:                   "list [OPTIONS] [PROFILE ...]",
		Short:                 "List profiles and configurations",
		Long:                  "This command will display configurations for PROFILE.",
		RunE:                  CobraRunE(&vars),
		Aliases:               []string{"ls"},
		ValidArgsFunction:     completions.Profiles,
		Args:                  cobra.ArbitraryArgs,
	}
	baseCmd.PersistentFlags().BoolVarP(&vars.showAll, "all", "a", false, "Show all profile configurations")
	baseCmd.PersistentFlags().BoolVarP(&vars.showYaml, "yaml", "y", false, "Show profile configurations via yaml format")
	baseCmd.PersistentFlags().BoolVarP(&vars.showJson, "json", "j", false, "Show profile configurations via json format")

	return baseCmd
}
