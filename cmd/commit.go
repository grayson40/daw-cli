/*
Copyright © 2022 Grayson Crozier <grayson40@gmail.com>
*/
package cmd

import (
	commit "github.com/grayson40/daw/pkg/logic"
	"github.com/spf13/cobra"
)

// commitCmd represents the commit command
var commitCmd = &cobra.Command{
	Use:   "commit [message]",
	Short: "Record changes to the project file",
	Long:  `This command will commit staged project file(s) with a specified message.`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		commit.ExecuteCommit(args[0])
	},
}

func init() {
	rootCmd.AddCommand(commitCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// commitCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// commitCmd.Flags().StringP("message", "m", "", "Commit message")
}
