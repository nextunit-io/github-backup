package backup

import (
	"log"
	"os"

	"github.com/google/go-github/v80/github"
)

type BackupInput struct {
	OutputFile    string
	Organisations []string
	Users         []string
}

func (b backup) StartBackup(input BackupInput) error {
	repos := []*github.Repository{}

	for _, user := range input.Users {
		userRepos, err := b.getUserRepositories(user)
		if err != nil {
			return err
		}
		repos = append(repos, userRepos...)
	}

	for _, org := range input.Organisations {
		orgRepos, err := b.getOrgainizationRepositories(org)
		if err != nil {
			return err
		}
		repos = append(repos, orgRepos...)
	}

	log.Printf("Total repositories to back up: %d, start downloading...\n", len(repos))

	for _, repo := range repos {
		err := b.cloneRepository(repo)
		if err != nil {
			return err
		}
	}

	log.Println("Cloning completed successfully, start zipping...")

	err := b.zip(input.OutputFile)
	if err != nil {
		return err
	}

	log.Println("Zipping completed successfully, start cleaning up...")

	os.RemoveAll(b.workingDir)

	log.Println("Backup completed successfully.")

	return nil
}

func (b backup) getUserRepositories(user string) ([]*github.Repository, error) {
	repos, _, err := b.client.Repositories.ListByUser(b.ctx, user, nil)
	if err != nil {
		return nil, err
	}

	for _, repo := range repos {
		if b.verbose {
			log.Printf("Backing up repository: %s\n", repo.GetFullName())
		}
	}

	return repos, nil
}

func (b backup) getOrgainizationRepositories(org string) ([]*github.Repository, error) {
	repos, _, err := b.client.Repositories.ListByOrg(b.ctx, org, nil)
	if err != nil {
		return nil, err
	}

	for _, repo := range repos {
		if b.verbose {
			log.Printf("Backing up repository: %s\n", repo.GetFullName())
		}
	}

	return repos, nil
}
