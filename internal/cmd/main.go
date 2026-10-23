package cmd

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
)

func parseCommitType(branch string, branchErr error) (string, error) {
	if branchErr != nil {
		logger.Printf(
			"Error getting branch: '%s'. Falling back to '%s'\n",
			branchErr,
			commitTypeFallback,
		)
		return commitTypeFallback, branchErr
	}
	logger.Printf("Checking branch name '%s' for type prefix\n", branch)
	for _, commitType := range commitTypes {
		pattern := regexp.MustCompile("^" + commitType)
		if pattern.MatchString(branch) {
			return commitType, nil
		}
	}
	logger.Printf("No type matched, falling back to '%s'\n", commitTypeFallback)
	return commitTypeFallback, nil
}

func main(cmd *cobra.Command, args []string) error {
	flags := cmd.Flags()

	verbose, err := flags.GetBool(verboseFlag)
	if err != nil {
		return err
	}
	logger = newLogger(verbose)
	logger.Printf("Verbose logging enabled\n")

	branch, branchErr := gitBranch()

	commitType, err := flags.GetString(commitTypeFlag)
	logger.Printf("Type flag value: '%s'\n", commitType)
	if err != nil {
		return err
	}
	if commitType == "" {
		logger.Printf("Type flag not specified, determining default")
		commitType, err = parseCommitType(branch, branchErr)
		if err != nil {
			return err
		}
	}
	if !isValidCommitType(commitType) {
		return fmt.Errorf("'%s' is not a valid commit type", commitType)
	}

	subject := strings.Join(args, " ")
	logger.Printf("Subject: '%s'\n", subject)

	fmt.Printf("%s: %s\n", commitType, subject)
	return nil
}
