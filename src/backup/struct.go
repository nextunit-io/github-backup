package backup

import (
	"context"

	"github.com/google/go-github/v80/github"
)

type backup struct {
	token      string
	client     *github.Client
	ctx        context.Context
	user       *github.User
	workingDir string
	verbose    bool
}

func NewBackup(token, workingDir string, verbose bool) (*backup, error) {
	b := &backup{
		ctx:        context.Background(),
		workingDir: workingDir,
		token:      token,
		verbose:    verbose,
	}
	err := b.Login()
	if err != nil {
		return nil, err
	}

	return b, nil
}
