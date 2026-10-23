package cmd

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
)

func parseCommitType(branch string, branchErr error) (string, error) {
	logger.Printf("Attempting to parse commit type\n")
	if branchErr != nil {
		logger.Printf(
			"Error getting branch: '%s'. Falling back to '%s'\n",
			branchErr,
			commitTypeFallback,
		)
		return commitTypeFallback, nil
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

func parseScope(branch string, branchErr error) (string, error) {
	logger.Printf("Attempting to parse scope\n")
	if branchErr != nil {
		logger.Printf(
			"Error getting branch: '%s'. Falling back to '%s'\n",
			branchErr,
			scopeFallback,
		)
		return scopeFallback, nil
	}
	fmt.Printf("TODO: Parse branch\n")
	return scopeFallback, nil
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
		logger.Printf("Type flag not specified")
		commitType, err = parseCommitType(branch, branchErr)
		if err != nil {
			return err
		}
	}
	if !isValidCommitType(commitType) {
		return fmt.Errorf("'%s' is not a valid commit type", commitType)
	}

	scope, err := flags.GetString(scopeFlag)
	if err != nil {
		return err
	}
	if scope == "" {
		scope, err = parseScope(branch, branchErr)
		if err != nil {
			return err
		}
	}

	subject := strings.Join(args, " ")
	logger.Printf("Subject: '%s'\n", subject)

	message := commitType
	if scope != "" && scope != scopeFlagSentinel {
		message = fmt.Sprintf("%s(%s)", message, scope)
	}
	message = fmt.Sprintf("%s: %s", message, subject)
	fmt.Printf("%s\n", message)
	return nil
}
