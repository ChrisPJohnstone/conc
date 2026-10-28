package git

import (
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/go-git/go-git/v6"
	"github.com/go-git/go-git/v6/plumbing/object"
)

func gitRepo() (*git.Repository, error) {
	repo, err := git.PlainOpen(".")
	if err != nil {
		return nil, err
	}
	slog.Debug("Repo opened succesfully")
	return repo, nil
}

func gitSignature(repo *git.Repository) (*object.Signature, error) {
	repoCfg, err := repo.Config()
	if err != nil {
		return nil, err
	}
	cfg := &gitConfig{repoCfg: repoCfg}
	name := cfg.repoCfg.User.Name
	if name == "" {
		slog.Debug("Git user name not set in repo config")
		if !cfg.hasGlobalCfg {
			err := cfg.getGlobalConfig()
			if err != nil {
				return nil, err
			}
		}
		name = cfg.globalCfg.User.Name
		if name == "" {
			return nil, errors.New("unable to get committer name")
		}
	}
	email := cfg.repoCfg.User.Email
	if email == "" {
		slog.Debug("Git user email not set in repo config")
		if !cfg.hasGlobalCfg {
			err := cfg.getGlobalConfig()
			if err != nil {
				return nil, err
			}
		}
		email = cfg.globalCfg.User.Email
		if email == "" {
			return nil, errors.New("unable to get committer email")
		}
	}
	return &object.Signature{Name: name, Email: email, When: time.Now()}, nil
}

func GitBranch() (string, error) {
	slog.Debug("Getting git branch")
	repo, err := gitRepo()
	if err != nil {
		return "", err
	}
	ref, err := repo.Head()
	if err != nil {
		return "", err
	}
	branch := ref.Name().Short()
	slog.Debug(fmt.Sprintf("Found branch '%s'", branch))
	return branch, nil
}

func GitCommit(message string) error {
	repo, err := gitRepo()
	if err != nil {
		return err
	}
	slog.Debug("Getting worktree")
	worktree, err := repo.Worktree()
	if err != nil {
		return err
	}
	slog.Debug("Getting author")
	author, err := gitSignature(repo)
	if err != nil {
		return err
	}
	slog.Debug(fmt.Sprintf(
		"Commiting with message '%s' & author '%s'",
		message,
		author,
	))
	commit, err := worktree.Commit(message, &git.CommitOptions{Author: author})
	if err != nil {
		return err
	}
	fmt.Printf("%s %s\n", commit, message)
	return nil
}
