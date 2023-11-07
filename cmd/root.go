/*
Copyright © 2023 Ajith

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/ajthr/gserve/internal/server"
)

var rootCmd = &cobra.Command{
	Use:   "gserve",
	Short: "CLI tool to create a simple, zero-configuration HTTP file server.",
	Run: server.Init,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringP("port", "p", "7777", "Help message for toggle")
	rootCmd.Flags().StringP("directory", "d", ".", "Help message")
}
