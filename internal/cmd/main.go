package cmd

import (
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
)

func commitTypeDefault() (string, error) {
	branch, err := gitBranch()
	if err != nil {
		return "", err
	}
	for _, commitType := range commitTypes {
		pattern := regexp.MustCompile("^" + commitType)
		if pattern.MatchString(branch) {
			return commitType, nil
		}
	}
	return commitTypeFallback, nil
}

func main(cmd *cobra.Command, args []string) error {
	flags := cmd.Flags()
	verbose, err := flags.GetBool(verboseFlag)
	if err != nil {
		return err
	}
	if verbose {
		logger = log.New(os.Stderr, "", log.LstdFlags)
		logger.Printf("Verbose logging enabled\n")
	} else {
		logger = log.New(io.Discard, "", 0)
	}
	commitType, err := flags.GetString(commitTypeFlag)
	if err != nil {
		return err
	}
	if commitType == "" {
		commitType, err = commitTypeDefault()
		if err != nil {
			return err
		}
	}
	if !isValidCommitType(commitType) {
		return fmt.Errorf("%s is not a valid commit type", commitType)
	}
	subject := strings.Join(args, " ")
	fmt.Printf("%s: %s\n", commitType, subject)
	return nil
}
