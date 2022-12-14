/*
Copyright © 2022 Grayson Crozier <grayson40@gmail.com>
*/
package cmd

import (
	"log"

	config "github.com/grayson40/daw/pkg/logic"
	"github.com/spf13/cobra"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure user credentials",
	Long:  `This command will configure the credentials of the user.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Get flags
		username, _ := cmd.Flags().GetString("username")
		email, _ := cmd.Flags().GetString("email")

		// Check for errors
		if username == "" && email == "" {
			log.Fatal("Error: requires at least 2 args, received 0")
		}
		if username == "" {
			log.Fatal("Error: flag needs an argument: --username")
		}
		if email == "" {
			log.Fatal("Error: flag needs an argument: --email")
		}

		config.ExecuteConfig(username, email)
	},
}

func init() {
	rootCmd.AddCommand(configCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	configCmd.PersistentFlags().String("email", "", "User email")
	configCmd.PersistentFlags().String("username", "", "User name")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
