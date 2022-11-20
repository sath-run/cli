/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/spf13/cobra"
)

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: runStatus,
}

func printStatusResult(result map[string]interface{}) {
	if result == nil {
		fmt.Println("no job is running")
	} else {
		fmt.Printf("id: %s\n", result["id"])
		fmt.Printf("status: %s\n", result["status"])
		fmt.Printf("progress: %f\n", result["progress"])
		fmt.Printf("message: %s\n", result["message"])
		fmt.Println()
	}
}

func runStatus(cmd *cobra.Command, args []string) {
	follow, err := cmd.Flags().GetBool("follow")
	if err != nil {
		log.Fatal(err)
	}
	if follow {
		resp, err := http.Get(origin + "/jobs/current/stream")
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		scanner := bufio.NewScanner(resp.Body)
		for scanner.Scan() {
			response := scanner.Text()
			if strings.HasPrefix(response, "data:") {
				response = strings.TrimPrefix(response, "data:")
				var result map[string]interface{}
				if err := json.Unmarshal([]byte(response), &result); err != nil {
					log.Fatal(err)
				}
				printStatusResult(result)
			}
		}
	} else {
		result := EngineGet("/jobs/current")
		printStatusResult(result)
	}

}

func init() {
	rootCmd.AddCommand(statusCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// statusCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	statusCmd.Flags().BoolP("follow", "f", false, "This option cause sath to stream status")
}
