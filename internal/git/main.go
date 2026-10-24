package git

import (
	"fmt"
	"log/slog"

	"github.com/go-git/go-git/v6"
)

func GitBranch() (string, error) {
	slog.Debug("Getting git branch\n")
	repo, err := git.PlainOpen(".")
	if err != nil {
		return "", err
	}
	slog.Debug("Repo opened succesfully\n")
	ref, err := repo.Head()
	if err != nil {
		return "", err
	}
	branch := ref.Name().Short()
	slog.Debug(fmt.Sprintf("Found branch '%s'\n", branch))
	return branch, nil
}
