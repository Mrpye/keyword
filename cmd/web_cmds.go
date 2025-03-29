package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/gelembjuk/articletext"
	"github.com/gelembjuk/keyphrases"
	"github.com/spf13/cobra"
)

func ExtractText_Command() *cobra.Command {

	//var config_path string
	//var web_port string
	//var web_ip string

	var cmd = &cobra.Command{
		Use:   "text",
		Short: "Extract text file",
		Long:  "Extract text file",
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
			filecontents, _ := ioutil.ReadFile(textfile)

			text := string(filecontents)

			analyser := keyphrases.TextPhrases{Language: "english"}

			analyser.Init()

			phrases := analyser.GetKeyWords(text)

			// write to a file
			file, err := os.Create(args[1])
			if err != nil {
				fmt.Println(err.Error())
				os.Exit(3)
			}
			defer file.Close()

			for _, phrase := range phrases {
				_, err = file.WriteString(phrase + "\n")
				if err != nil {
					fmt.Println(err.Error())
					os.Exit(4)
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
func ExtractWebpage_Command() *cobra.Command {

	//var config_path string
	//var web_port string
	//var web_ip string

	var cmd = &cobra.Command{
		Use:   "web",
		Short: "Extract Webpage",
		Long:  "Extract Webpage",
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

			// get text from this web page
			text, err := articletext.GetArticleTextFromUrl(url)

			if err != nil {
				fmt.Println(err.Error())
				os.Exit(2)
			}
			// print a text to a console
			//fmt.Println(text)

			// Create a text analyser object. It requires a path to WordNet dictionary directory

			analyser := keyphrases.TextPhrases{Language: "english",
				LanguageOptions: map[string]string{"wordnetdirectory": "./WordNet/dict"}}

			// this is required procedure to initialise analyser
			analyser.Init()

			// get key phrases
			phrases := analyser.GetKeyWords(text)
			// write to a file
			file, err := os.Create(args[1])
			if err != nil {
				fmt.Println(err.Error())
				os.Exit(3)
			}
			defer file.Close()

			for _, phrase := range phrases {
				_, err = file.WriteString(phrase + "\n")
				if err != nil {
					fmt.Println(err.Error())
					os.Exit(4)
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
func init() {

	rootCmd.AddCommand(ExtractText_Command())
	rootCmd.AddCommand(ExtractWebpage_Command())

}
