package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// uninstallCmd represents the uninstall command
var uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "Specify the version of golang needed to uninstalled/removed",
	Long: `Specify the version of golang needed to uninstalled/removed

		# To remove a specific version
		$ gvm uninstall 1.19
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("uninstall called")
	},
}

func init() {
	rootCmd.AddCommand(uninstallCmd)
}
