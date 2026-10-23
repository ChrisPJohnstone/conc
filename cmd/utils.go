package cmd

import (
	"io"
	"log"
	"os"

	"github.com/go-git/go-git/v6"
)

func newLogger(verbose bool) *log.Logger {
	if verbose {
		return log.New(os.Stderr, "", log.LstdFlags)
	} else {
		return log.New(io.Discard, "", 0)
	}
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
