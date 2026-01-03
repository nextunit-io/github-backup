package backup

import (
	"log"

	"github.com/google/go-github/v80/github"
	git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/http"
)

func (b *backup) cloneRepository(repo *github.Repository) error {
	if b.verbose {
		log.Printf("Cloning repository: %s\n", repo.GetFullName())
	}

	_, err := git.PlainClone(b.workingDir+"/"+repo.GetName(), false, &git.CloneOptions{
		URL: repo.GetCloneURL(),
		Auth: &http.BasicAuth{
			Username: *b.user.Login,
			Password: b.token,
		},
		Depth: 5,
	})
	if err != nil {
		return err
	}

	return nil
}
