package cmd

import "github.com/spf13/cobra"

func RootCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "cli-cobra",
		Short: "CLI CLI",
		Long:  `CLI CLI`,
	}
}
