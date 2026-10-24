package git

import (
	"log"

	"github.com/go-git/go-git/v6"
)

func GitBranch() (string, error) {
	log.Printf("Getting git branch\n")
	repo, err := git.PlainOpen(".")
	if err != nil {
		return "", err
	}
	log.Printf("Repo opened succesfully\n")
	ref, err := repo.Head()
	if err != nil {
		return "", err
	}
	branch := ref.Name().Short()
	log.Printf("Found branch '%s'\n", branch)
	return branch, nil
}
