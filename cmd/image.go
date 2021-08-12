/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"github.com/id1893/kubeimg/pkg/client"
	"github.com/id1893/kubeimg/pkg/tools"
	"github.com/spf13/cobra"
	"k8s.io/client-go/1.5/pkg/api"
)

var (
	example = `
			# List all deployment images version in ps output format.
			kubeimg image 
			# List the deployment images version of the specified namespaces
			kubeimg image -n <namespace>
`
)

func Images() error {
	clientset := client.Clientset()
	n, _ := rootCmd.Flags().GetString("namespace")
	deploymentList, _ := clientset.Deployments(n).List(api.ListOptions{})
	var itemSlice []string
	itemSlice = append(itemSlice, "NAMESPACE\tDEPLOY\tCONTAINER\tIMAGE")
	for i := 0; i < len(deploymentList.Items); i++ {
		d := deploymentList.Items[i]
		for i := 0; i < len(d.Spec.Template.Spec.Containers); i++ {
			c := d.Spec.Template.Spec.Containers[i]
			item := d.Namespace + "\t" + d.Name + "\t" + c.Name + "\t" + c.Image
			itemSlice = append(itemSlice, item)
		}
	}
	tools.Fprint(itemSlice)
	return nil

}

// imageCmd represents the image command
var imageCmd = &cobra.Command{
	Use:   "image",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tools to generate the needed files
to quickly create a Cobra application.`,
	Example: example,
	Run: func(cmd *cobra.Command, args []string) {
		if err := Images(); err != nil {
			panic(err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(imageCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// imageCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// imageCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
