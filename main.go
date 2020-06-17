package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rcommand = &cobra.Command{
	Use: "kodo",
}

var versionCommand = &cobra.Command{
	Use: "version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Version 2")
	},
}

var (
	cluster   string
	namespace string
	token     string
)

func init() {
	rcommand.AddCommand(versionCommand)
	rcommand.AddCommand(listCommand)
	rcommand.PersistentFlags().StringVarP(&cluster, "server", "s", "myurl", "this is the cluster url")
	rcommand.PersistentFlags().StringVarP(&token, "token", "t", "usertoken", "this is the user token")
	rcommand.PersistentFlags().StringVarP(&namespace, "namespace", "n", "namespace", "this is the namespace")
	rcommand.MarkFlagRequired("server")
}

func main() {
	rcommand.Execute()
}

var listCommand = &cobra.Command{
	Use: "list",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("List All Kubernetes Applications")
		fmt.Printf("Fetching all applications from %s in namespace %s", cluster, namespace)
	},
}
