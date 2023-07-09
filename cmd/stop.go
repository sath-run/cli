/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

// stopCmd represents the stop command
var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop SATH engine",
	Long: `Stop SATH engine.
After stopping, engine will cancel current running job
and will no longer start new jobs`,
	Run: runStop,
}

func runStop(cmd *cobra.Command, args []string) {
	wait, err := cmd.Flags().GetBool("wait")
	if err != nil {
		log.Fatal(err)
	}
	resp := EnginePost("/services/stop", map[string]interface{}{"wait": wait})
	fmt.Println(resp["message"])
}

func init() {
	rootCmd.AddCommand(stopCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// stopCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// stopCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	stopCmd.Flags().BoolP("wait", "w", false, "Wait for job completion before exit")
}
