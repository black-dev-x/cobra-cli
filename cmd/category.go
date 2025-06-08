package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// categoryCmd represents the category command
var categoryCmd = &cobra.Command{
	Use:   "category",
	Short: "Able to create or list categories",
	Long:  `The category command allows you to create or list categories.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
	PreRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("Preparing to execute category command...")
	},
	// PostRun: func(cmd *cobra.Command, args []string) {
	// 	fmt.Println("Category command execution completed.")
	// },
	// RunE: func(cmd *cobra.Command, args []string) error {
	// 	// This is where you can handle errors if needed
	// 	fmt.Println("category command executed")
	// 	return nil
	// },
	// PreRunE: func(cmd *cobra.Command, args []string) error {
	// 	// This is where you can handle errors before running the command
	// 	fmt.Println("Preparing to execute category command with error handling...")
	// 	return nil
	// },
	// PostRunE: func(cmd *cobra.Command, args []string) error {
	// 	// This is where you can handle errors after running the command
	// 	fmt.Println("Category command execution completed with error handling.")
	// 	return nil
	// },
}

func init() {
	rootCmd.AddCommand(categoryCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// categoryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// categoryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
