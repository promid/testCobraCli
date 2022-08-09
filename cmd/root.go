/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var saySomething string

// rootCmd represents the base command when called without any subcommands

var rootCmd = &cobra.Command{
	Use:   "testCobraCli",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Run:")
		name, _ := cmd.Flags().GetString("name")

		fmt.Printf("%s says %s\n", name, saySomething)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	opts := NewRunOptions()
	opts.AddFlags(rootCmd.Flags())
	// flags will be parsed in rootCmd.Execute(), so pflag.Parse() doesn't need to be called explicitly
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
	fmt.Printf(opts.Test)
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&saySomething, "say", "default words", "what you want say")
	fmt.Printf("init: saySomething = %s\n", saySomething)

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	defaultName := *rootCmd.Flags().StringP("name", "t", "Bei", "Help message for name")
	fmt.Printf("init: name = %s\n", defaultName)
}
