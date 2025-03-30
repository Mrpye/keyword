package cmd

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/securisec/go-keywords"
	"github.com/spf13/cobra"
)

func ExtractTextV2_Command() *cobra.Command {

	//var config_path string
	//var web_port string
	//var web_ip string

	var cmd = &cobra.Command{
		Use:   "text",
		Short: "Extract keywords from text file",
		Long:  "Extract keywords from text file",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				fmt.Println("No URL provided")
				os.Exit(1)
			}
			if len(args) < 2 {
				fmt.Println("No output file provided")
				os.Exit(1)
			}
			// get URL from command line argument

			textfile := args[0]
			filecontents, _ := os.ReadFile(textfile)

			text := string(filecontents)

			k, err := keywords.Extract(text)
			if err != nil {
				return fmt.Errorf("error extracting keywords: %v", err)
			}
			// write to a file
			file, err := os.Create(args[1])
			if err != nil {
				return fmt.Errorf("Error creating file: %v", err)
			}
			defer file.Close()

			for _, phrase := range k {
				_, err = file.WriteString(phrase + "\n")
				if err != nil {
					return fmt.Errorf("error writing to file: %v", err)
				}
				fmt.Println(phrase)
			}
			return nil
		},
	}
	//cmd.Flags().StringVarP(&config_path, "config", "c", "./config", "config path")
	//cmd.Flags().StringVarP(&web_port, "port", "p", "8080", "Listen on Port")
	//cmd.Flags().StringVarP(&web_ip, "ip", "i", "localhost", "Listen on Ip")

	return cmd

}
func ExtractWebpageV2_Command() *cobra.Command {

	var strip_tags bool
	var remove_duplicates bool
	var lowercase bool
	var ignore_pattern string
	//var web_port string
	//var web_ip string

	var cmd = &cobra.Command{
		Use:   "web",
		Short: "Extract keywords from Webpage",
		Long:  "Extract keywords from Webpage",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				fmt.Println("No URL provided")
				os.Exit(1)
			}
			if len(args) < 2 {
				fmt.Println("No output file provided")
				os.Exit(1)
			}
			// get URL from command line argument
			url := args[0]

			req, err := http.Get(url)
			if err != nil {
				return fmt.Errorf("error getting URL: %v", err)
			}
			data, err := io.ReadAll(req.Body)
			if err != nil {
				return fmt.Errorf("error reading response body: %v", err)
			}
			k, err := keywords.Extract(string(data), keywords.ExtractOptions{
				StripTags:        strip_tags,
				RemoveDuplicates: remove_duplicates,
				IgnorePattern:    ignore_pattern,
				Lowercase:        lowercase,
			})
			if err != nil {
				return fmt.Errorf("error extracting keywords: %v", err)
			}

			// write to a file
			file, err := os.Create(args[1])
			if err != nil {
				return fmt.Errorf("error creating file: %v", err)
			}
			defer file.Close()

			for _, phrase := range k {
				_, err = file.WriteString(phrase + "\n")
				if err != nil {
					return fmt.Errorf("error writing to file: %v", err)
				}
				fmt.Println(phrase)
			}
			return nil
		},
	}
	cmd.Flags().BoolVarP(&strip_tags, "tags", "t", false, "Strip HTML tags")
	cmd.Flags().BoolVarP(&remove_duplicates, "duplicates", "d", false, "Remove duplicates")
	cmd.Flags().BoolVarP(&lowercase, "lowercase", "l", false, "Lowercase")
	cmd.Flags().StringVarP(&ignore_pattern, "ignore", "i", "<.+>", "Ignore pattern")

	return cmd

}
func init() {

	method2Cmd.AddCommand(ExtractTextV2_Command())
	method2Cmd.AddCommand(ExtractWebpageV2_Command())

}
