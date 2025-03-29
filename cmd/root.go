/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "keyword",
	Short: "extract keywords",
	Long:  `extract keywords`,
}

func GenerateDoc() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "gen_docs",
		Short: "This command will build the documents for the cli",
		Long:  `This command will build the documents for the cli`,
		RunE: func(cmd *cobra.Command, args []string) error {
			os.MkdirAll("./documents", os.ModePerm)
			err := doc.GenMarkdownTree(rootCmd, "./documents")
			if err != nil {
				return err
			}
			fmt.Println("Documents Generated")
			//lib.PrintlnOK("Documents Generated")
			return nil
		},
	}
	return cmd
}

func Execute() {

	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(os.Stderr, "Error:", err)
		os.Exit(1)
	} else {
		fmt.Println("Completed")
	}

}

func init() {

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.PersistentFlags().BoolP("help", "", false, "help for this command")
	rootCmd.AddCommand(GenerateDoc())

}
