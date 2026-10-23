package cmd

import (
	"slices"

	"github.com/go-git/go-git/v6"
)

func isValidCommitType(commitType string) bool {
	return slices.Contains(commitTypes, commitType)
}

func gitBranch() (string, error) {
	repo, err := git.PlainOpen(".")
	if err != nil {
		return "", err
	}
	ref, err := repo.Head()
	if err != nil {
		return "", err
	}
	branch := ref.Name().Short()
	return branch, nil
}
