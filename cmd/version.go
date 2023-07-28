/*
Copyright © 2023 tochiman development@tochiman.com
*/
package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print GoogleDriveController version information",
	Long: `Print GoogleDriveController version information`,
	Run: func(cmd *cobra.Command, args []string) {
		version := os.Getenv("VERSION")
		fmt.Println(version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// versionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// versionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
