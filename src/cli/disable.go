package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

// disableCmd represents the disable command
var disableCmd = &cobra.Command{
	Use:       fmt.Sprintf(toggleUse, "disable"),
	Short:     "Disable a feature",
	Long:      fmt.Sprintf(toggleLong, "Disable"),
	ValidArgs: toggleArgs,
	Args:      NoArgsOrOneValidArg,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			_ = cmd.Help()
			return
		}
		toggleFeature(cmd, args[0], false)
	},
}

func init() {
	RootCmd.AddCommand(disableCmd)
}
