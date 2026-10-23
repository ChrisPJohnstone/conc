package cmd

import (
	"fmt"
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
