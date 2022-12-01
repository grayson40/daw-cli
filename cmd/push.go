/*
Copyright © 2022 Grayson Crozier <graysoncrozier40@gmail.com>
*/
package cmd

import (
	push "github.com/grayson40/daw/pkg/logic"
	"github.com/spf13/cobra"
)

// pushCmd represents the push command
var pushCmd = &cobra.Command{
	Use:   "push",
	Short: "Push staged commits and update local ref",
	Long:  `This command will push the staged commits up the current ref`,
	Run: func(cmd *cobra.Command, args []string) {
		push.ExecutePush()
	},
}

func init() {
	rootCmd.AddCommand(pushCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pushCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pushCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
