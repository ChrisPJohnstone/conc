package cmd

import (
	"slices"

	"github.com/go-git/go-git/v6"
)

func isValidCommitType(commitType string) bool {
	return slices.Contains(commitTypes, commitType)
}

func gitBranch() (string, error) {
	logger.Printf("Getting git branch\n")
	repo, err := git.PlainOpen(".")
	if err != nil {
		return "", err
	}
	logger.Printf("Repo opened succesfully\n")
	ref, err := repo.Head()
	if err != nil {
		return "", err
	}
	branch := ref.Name().Short()
	logger.Printf("Found branch '%s'\n", branch)
	return branch, nil
}
