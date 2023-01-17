/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

const (
	RESOURCE_SERVICES = "services"
	RESOURCE_API      = "api"
)

type Service struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Services struct {
	Items []Service `json:"services"`
}

var validArgs = []string{RESOURCE_SERVICES, RESOURCE_API}

// Todos
// - Print the 0+n'th args in help and their information

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists API services by resource",
	Args: func(cmd *cobra.Command, args []string) error {

		err := cobra.MinimumNArgs(1)(cmd, args)
		if err != nil {
			fmt.Printf("Please specify a resource from %v\n", validArgs)
			return err
		}

		err = cobra.OnlyValidArgs(cmd, args)
		if err != nil {
			fmt.Printf("Please specify a resource from %v\n", validArgs)
			return err
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {

		resourceName := args[0]
		if resourceName == RESOURCE_SERVICES {
			services, err := getServices()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(services)

		} else if resourceName == RESOURCE_API {
			if len(args) < 2 {
				log.Fatal("Please specify an api.")
			}
			api, err := getApi(args[1])
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(api)

		} else {
			fmt.Printf("Please specify a resource from [%v]", RESOURCE_SERVICES)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func getServices() (string, error) {
	url := "https://api.openshift.com/api"

	response, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	if response.StatusCode == 200 {
		return string(body), nil
	}
	return "", nil
}

func getApi(api string) (string, error) {

	base_url := "https://api.openshift.com/api"
	url := fmt.Sprintf("%s/%s/v1/openapi", base_url, api)

	response, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	if response.StatusCode == 200 {
		return string(body), nil
	}
	return "", nil
}
