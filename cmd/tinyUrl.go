/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"log"
	"net/http"
)

// tinyUrlCmd represents the tinyUrl command
var tinyUrlCmd = &cobra.Command{
	Use:   "tinyUrl",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		uri, _ := cmd.Flags().GetString("uri")
		getTinyUrl(uri)
	},
}

func init() {
	rootCmd.AddCommand(tinyUrlCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// tinyUrlCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	tinyUrlCmd.Flags().StringP("uri", "u", "", "uri to make tiny")

}

func getTinyUrl(uri string) {
	tinyUrlApi := fmt.Sprintf("http://tinyurl.com/api-create.php?url=%s", uri)
	resp, err := http.Get(tinyUrlApi)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(resp.Body)
	bodyString := string(body)
	fmt.Println(bodyString)
}
