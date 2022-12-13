/*
Copyright © 2022 Grayson Crozier <grayson40@gmail.com>
*/
package cmd

import (
	status "github.com/grayson40/daw/pkg/logic"
	"github.com/spf13/cobra"
)

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show the working tree status",
	Long:  `Displays paths that have differences between the index file and the current HEAD commit`,
	Run: func(cmd *cobra.Command, args []string) {
		status.ExecuteStatus()
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// statusCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// statusCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
