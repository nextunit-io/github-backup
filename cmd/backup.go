package cmd

import (
	"errors"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/nextunit-io/github-backup/backup"
)

func runBackupCMD(flags PersistentFlags) error {
	err := flags.Validate()

	if err != nil {
		if flags.IsInteractive {
			err = runModelCMD(&flags)
		}

		if err != nil {
			return err
		}
	}

	backupClient, err := backup.NewBackup(flags.Token, WORKING_DIR, flags.Verbose)
	handleError(err)

	handleError(backupClient.StartBackup(backup.BackupInput{
		OutputFile:    flags.OutputFile,
		Organisations: flags.Orgs,
		Users:         flags.Users,
	}))

	return nil
}

func runModelCMD(flags *PersistentFlags) error {
	p := tea.NewProgram(generateModel(flags))
	m, err := p.Run()

	if err != nil {
		return err
	}

	model, ok := m.(CmdModel)
	if !ok {
		return errors.New("cannot type cast interactive model")
	}

	flags.Token = model.inputs[0].input.Value()
	flags.OutputFile = model.inputs[1].input.Value()
	flags.Users = []string{}
	if model.inputs[2].input.Value() != "" {
		flags.Users = strings.Split(model.inputs[2].input.Value(), ",")
	}
	flags.Orgs = []string{}
	if model.inputs[3].input.Value() != "" {
		flags.Orgs = strings.Split(model.inputs[3].input.Value(), ",")
	}

	return flags.Validate()
}
