/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/hiroaqii/go-bgg/bgg"
	"github.com/spf13/cobra"
)

// collectionCmd represents the collection command
var collectionCmd = &cobra.Command{
	Use:   "collection",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		username, err := cmd.Flags().GetString("username")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(username)

		collectioItems, err := bgg.Collection(username)
		if err != nil {
			println(err)
			return
		}

		e, _ := json.Marshal(collectioItems)
		fmt.Println(string(e))
	},
}

func init() {
	rootCmd.AddCommand(collectionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// collectionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	collectionCmd.Flags().StringP("username", "u", "", "Name of the user to re	quest the collection for.")
	collectionCmd.MarkFlagRequired("username")
}
