package backup

import (
	"fmt"

	"github.com/google/go-github/v80/github"
)

func (b *backup) Login() error {
	b.client = github.NewClient(nil).WithAuthToken(b.token)

	user, resp, err := b.client.Users.Get(b.ctx, "")
	if err != nil {
		return err
	}

	if resp.TokenExpiration.IsZero() {
		return fmt.Errorf("token expired")
	}

	b.user = user

	return nil
}
