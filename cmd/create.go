/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/black-dev-x/cobra-cli/database"
	"github.com/spf13/cobra"
)

type RunErrorFunction func(cmd *cobra.Command, args []string)

func runCreate(categoryDB *database.CategoryDB) RunErrorFunction {
	return func(cmd *cobra.Command, args []string) {
		fmt.Println("create category called")
		name, _ := cmd.Flags().GetString("name")
		description, _ := cmd.Flags().GetString("description")
		if _, err := categoryDB.Create(name, description); err != nil {
			fmt.Printf("Error creating category: %v\n", err)
		} else {
			fmt.Printf("Category '%s' created successfully!\n", name)
		}
	}
}

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new category",
	Long:  `The create command allows you to create a new category.`,
	Run:   runCreate(GetCategoryDB()),
}

func init() {
	categoryCmd.AddCommand(createCmd)
	createCmd.Flags().StringP("name", "n", "", "Name of the category")
	createCmd.MarkFlagRequired("name")
	createCmd.Flags().StringP("description", "d", "", "Description of the category")

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
