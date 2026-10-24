package git

import (
	"fmt"
	"log/slog"

	"github.com/go-git/go-git/v6"
)

func GitBranch() (string, error) {
	slog.Debug("Getting git branch")
	repo, err := git.PlainOpen(".")
	if err != nil {
		return "", err
	}
	slog.Debug("Repo opened succesfully")
	ref, err := repo.Head()
	if err != nil {
		return "", err
	}
	branch := ref.Name().Short()
	slog.Debug(fmt.Sprintf("Found branch '%s'", branch))
	return branch, nil
}
