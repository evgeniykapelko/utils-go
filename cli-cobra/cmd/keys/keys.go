package keys

import "github.com/spf13/cobra"

var keysCmd = &cobra.Command{
	Use:   "keys",
	Short: "Manage keys",
	Long:  "Manage keys",
}

func Init(rootCmd *cobra.Command) {
	rootCmd.AddCommand(keysCmd)
}
