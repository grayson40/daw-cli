/*
Copyright © 2022 Grayson Crozier <grayson40@gmail.com>
*/
package cmd

import (
	handler "github.com/grayson40/daw/pkg/handlers"
	"github.com/spf13/cobra"
)

var stagedFile string

// restoreCmd represents the restore command
var restoreCmd = &cobra.Command{
	Use:   "restore",
	Short: "Restore working tree files",
	Long:  `Restore specified paths in the working tree with some contents from a restore source`,
	Run: func(cmd *cobra.Command, args []string) {
		handler.ExecuteRestore(stagedFile)
	},
}

func init() {
	rootCmd.AddCommand(restoreCmd)

	// Here you will define your flags and configuration settings.
	restoreCmd.Flags().StringVarP(&stagedFile, "staged", "s", "", "Specify the restore location")

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// restoreCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// restoreCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
