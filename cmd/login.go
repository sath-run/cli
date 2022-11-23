/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"syscall"

	"github.com/spf13/cobra"
	"golang.org/x/term"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: runLogin,
}

func runLogin(cmd *cobra.Command, args []string) {
	reader := bufio.NewReader(os.Stdin)
	var err error
	var username, password string

	for len(username) == 0 {
		fmt.Print("Enter Email: ")
		username, err = reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		} else {
			username = strings.Trim(username, "\n")
		}
	}

	for len(password) == 0 {
		fmt.Print("Enter Password: ")
		bytePassword, err := term.ReadPassword(int(syscall.Stdin))
		if err != nil {
			log.Fatal(err)
		}

		password = string(bytePassword)
		fmt.Println("")
	}

	res, status := sendRequestToEngine(http.MethodPost, "/users/login", map[string]interface{}{
		"email":    username,
		"password": password,
	})
	if status == http.StatusOK {
		fmt.Println("login success")
	} else {
		fmt.Printf("login failed: ")
		if message, ok := res["message"]; ok {
			if str, ok := message.(string); ok {
				fmt.Println(str)
				return
			}
		}
		fmt.Println()
	}
}

func init() {
	rootCmd.AddCommand(loginCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loginCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loginCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
