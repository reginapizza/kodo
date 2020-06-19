package main

import (
	"fmt"

	"github.com/cli-playground/kodo/pkg/kodo/cmd"
	"github.com/spf13/cobra"
)

var envVar = new(cmd.EnvironmentVariables)   // creating instance of struct EnvironmentVariables from cmd.openshiftclient.go
var deployVar = new(cmd.DeploymentVariables) // creating instance of struct EnvironmentVariables from cmd.deploy.go

var rcommand = &cobra.Command{
	Use: "kodo",
}

var versionCommand = &cobra.Command{
	Use: "version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Version 2")
	},
}

func init() {
	rcommand.AddCommand(versionCommand)
	rcommand.AddCommand(countCommand)
	countCommand.AddCommand(podCommand)
	rcommand.AddCommand(deployCommand)

	deployCommand.Flags().StringVarP(&envVar.Host, "server", "s", "", "this is the cluster url")
	deployCommand.Flags().StringVarP(&envVar.Bearertoken, "token", "t", "", "this is the user token")
	deployCommand.Flags().StringVarP(&envVar.Namespace, "namespace", "n", "", "this is the namespace")

	/* Changed default namespace from 'namespace' to "" because
	'namespace' causes some internal error and returns 0 pods
	even if login credentials correct but namespace not explicitly set
	whereas "" works fine */

	//Flag for individual command

	deployCommand.Flags().StringVarP(&deployVar.Image, "image", "i", "myimage", "tagged image for deployment")
	deployCommand.Flags().Int32VarP(&deployVar.Replicas, "replicas", "r", 1, "number of replicas")
	deployCommand.Flags().Int32VarP(&deployVar.Port, "port", "p", 8000, "port at which app should run")

	//required attributes for deploy command
	deployCommand.MarkFlagRequired("server")
	deployCommand.MarkFlagRequired("token")
	deployCommand.MarkFlagRequired("image")

	podCommand.Flags().StringVarP(&envVar.Host, "server", "s", "", "this is the cluster url")
	podCommand.Flags().StringVarP(&envVar.Bearertoken, "token", "t", "", "this is the user token")
	podCommand.Flags().StringVarP(&envVar.Namespace, "namespace", "n", "", "this is the namespace")
	//required parameters for count pods command
	podCommand.MarkFlagRequired("server")
	podCommand.MarkFlagRequired("token")

}

func main() {

	rcommand.Execute()
}

var countCommand = &cobra.Command{
	Use:   "count",
	Short: "Command to count <resources>",
}

var podCommand = &cobra.Command{
	Use: "pods",
	Run: func(cm *cobra.Command, args []string) {
		fmt.Println("List All Kubernetes Applications")
		fmt.Printf("\nFetching all applications from %s in namespace %s", envVar.Host, envVar.Namespace)
		cmd.List(envVar)
	},
}

var deployCommand = &cobra.Command{
	Use:   "deploy",
	Short: "Command to deploy an image",
	Run: func(cm *cobra.Command, args []string) {
		fmt.Println("Creating image deployment")
		cmd.Deploy(deployVar, envVar)
	},
}
