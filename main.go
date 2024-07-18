/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"tini-cli/cmd"
	"tini-cli/config"
)

func main() {
	// Initialize the
	commonConfig := config.NewConfig()

	// Run the command
	cmd.Execute(commonConfig)
}
