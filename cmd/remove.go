/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var removeCommand = &cobra.Command{
	Use:   "remove",
	Short: "Removes a package from your gomod",
	Long:  `Removes a package from your go module; short and sweet!`,
	Run: func(cmd *cobra.Command, args []string) {
		var bashcmd *exec.Cmd
		verbose, _ := cmd.Flags().GetBool("verbose")
		url, err := cmd.Flags().GetString("url")
		if err != nil {
			cmd.Println(err)
			return
		}
		if url == "" {
			cmd.Println("You must include the url flag (-u)")
			return
		}
		if strings.HasPrefix(url, "https://") {
			cut, found := strings.CutPrefix(url, "https://")
			if !found && verbose {
				fmt.Println("No https:// prefix")
			}
			url = cut
		}
		bashcmd = exec.Command("go", "get", url+"@none")
		if verbose {
			fmt.Println(bashcmd)
		}
		output, cmderr := bashcmd.CombinedOutput()
		if cmderr != nil {
			panic(cmderr)
		}
		if len(output) > 0 {
			fmt.Println(string(output))
		} else {
			fmt.Println("go get returned no output, it most likely errored")
		}

		fmt.Printf("Successfully removed %v\n", url)
	},
}

func init() {
	removeCommand.Flags().StringP("url", "u", "", "the URL")
	rootCmd.AddCommand(removeCommand)
}
