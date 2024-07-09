package signatures

import "github.com/spf13/cobra"

var signaturesCmd = &cobra.Command{
	Use:   "signatures",
	Short: "Signatures",
	Long:  "Signatures",
}

func Init(rootCmd *cobra.Command) {
	rootCmd.AddCommand(signaturesCmd)
}
