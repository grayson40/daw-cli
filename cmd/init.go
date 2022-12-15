/*
Copyright © 2022 Grayson Crozier <grayson40@gmail.com>
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	initialize "github.com/grayson40/daw/pkg/logic"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Create an empty Daw repository or reinitialize an existing one",
	Long:  `This command creates an empty Git repository - basically a .daw directory with subdirectories for objects.`,
	Run: func(cmd *cobra.Command, args []string) {
		path, err := os.Getwd()
		if err != nil {
			log.Fatal(err.Error())
		}
		if initialize.IsInitialized() {
			fmt.Println("Reinitialized existing Daw repository in " + path + "\\.daw")
		} else {
			initialize.ExecuteInit()
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
