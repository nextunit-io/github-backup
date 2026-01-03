package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const (
	CMD_ROOT    = "github-backup"
	WORKING_DIR = ".github-backup"
)

var rootCMD *cobra.Command

type PersistentFlags struct {
	Token         string `validate:"required,min=1"`
	Verbose       bool
	Users         []string `validate:"required_without=Orgs,omitempty"`
	Orgs          []string `validate:"required_without=Users,omitempty"`
	OutputFile    string   `validate:"required,min=1"`
	IsInteractive bool
}

func handleError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func init() {
	pflags := PersistentFlags{}

	rootCMD = &cobra.Command{
		Use:   fmt.Sprintf("%s", CMD_ROOT),
		Short: fmt.Sprintf("%s is cli operation tool to backup github repositories upon given inputs", CMD_ROOT),
		Long:  fmt.Sprintf("%s is a cli tool that provides multiple functions to backup github repositories", CMD_ROOT),
		Run: func(cmd *cobra.Command, args []string) {
			handleError(runBackupCMD(pflags))
		},
	}

	var nonInteractive bool
	rootCMD.PersistentFlags().StringVarP(&pflags.Token, "token", "t", "", "GitHub personal access token")
	rootCMD.PersistentFlags().BoolVarP(&pflags.Verbose, "verbose", "v", false, "Enable verbose output")
	rootCMD.PersistentFlags().StringSliceVarP(&pflags.Users, "users", "u", []string{}, "GitHub users to backup")
	rootCMD.PersistentFlags().StringSliceVarP(&pflags.Orgs, "orgs", "o", []string{}, "GitHub organizations to backup")
	rootCMD.PersistentFlags().StringVarP(&pflags.OutputFile, "file", "f", "", "Output file for the backup zip")
	rootCMD.PersistentFlags().BoolVarP(&nonInteractive, "non-interactive", "i", false, "Disables interactive mode")
	pflags.IsInteractive = !nonInteractive
}

func Execute() {
	rootCMD.Execute()
}
