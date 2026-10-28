package git

import (
	"errors"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/go-git/go-git/v6/config"
)

type gitConfig struct {
	repoCfg      *config.Config
	globalCfg    *config.Config
	hasGlobalCfg bool
}

func readConfig(path string) (*config.Config, error) {
	slog.Debug(fmt.Sprintf("Reading git config from path: '%s'\n", path))
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer func() { _ = file.Close() }()
	return config.ReadConfig(file)
}

// go-git doesn't follow absolute symlinks e.g. ~/.config/git -> ~/dotfiles/git
// Manual reader to workaround
func (cfg *gitConfig) getGlobalConfig() error {
	slog.Debug("Getting global git config")
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	xdg := os.Getenv("XDG_CONFIG_HOME")
	if xdg == "" {
		xdg = filepath.Join(home, ".config")
	}
	paths := []string{
		filepath.Join(home, ".gitconfig"),
		filepath.Join(xdg, "git", "config"),
	}
	for _, path := range paths {
		cfg.globalCfg, err = readConfig(path)
		if err == nil {
			cfg.hasGlobalCfg = true
			return err
		}
	}
	return errors.New("unable to read global config")
}
