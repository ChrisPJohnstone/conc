package cmd

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
)

func parseCommitType(
	input string,
	branch string,
	branchErr error,
) (string, error) {
	logger.Printf("Attempting to parse commit type\n")
	if input != "" {
		logger.Printf("Type flag specified: '%s'\n", input)
		var matches []string
		for _, commitType := range commitTypes {
			if strings.HasPrefix(commitType, input) {
				logger.Printf("Flag matches '%s'\n", commitType)
				matches = append(matches, commitType)
			}
		}
		switch len(matches) {
		case 0:
			return commitTypeFallback, fmt.Errorf(
				"'%s' is not a valid commit type",
				input,
			)
		case 1:
			return matches[0], nil
		default:
			return commitTypeFallback, fmt.Errorf(
				"'%s' is ambiguous: %s. Use a longer prefix",
				input,
				strings.Join(matches, ", "),
			)
		}
	}
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
		if strings.HasPrefix(branch, commitType) {
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
	pattern := regexp.MustCompile(scopeBranchPattern)
	matches := pattern.FindStringSubmatch(branch)
	if len(matches) == 2 {
		scope := matches[1]
		logger.Printf("Branch name matches pattern, found scope: '%s'\n", scope)
		return scope, nil
	}
	return scopeFallback, nil
}

func parseMessage(
	commitType string,
	scope string,
	subject string,
	body string,
	footer string,
) string {
	message := commitType
	if scope != "" && scope != scopeFlagSentinel {
		message = fmt.Sprintf("%s(%s)", message, scope)
	}
	message = fmt.Sprintf("%s: %s", message, subject)
	if body != "" {
		message = fmt.Sprintf("%s\n\n%s", message, body)
	}
	if footer != "" {
		message = fmt.Sprintf("%s\n\n%s", message, footer)
	}
	return message
}

func main(cmd *cobra.Command, args []string) error {
	flags := cmd.Flags()

	verbose, err := flags.GetBool(verboseFlag)
	if err != nil {
		return err
	}
	logger = newLogger(verbose)
	logger.Printf("Verbose logging enabled\n")

	subject := strings.Join(args, " ")
	commitTypeFlagVal, err := flags.GetString(commitTypeFlag)
	if err != nil {
		return err
	}
	scope, err := flags.GetString(scopeFlag)
	if err != nil {
		return err
	}
	body, err := flags.GetString(bodyFlag)
	if err != nil {
		return err
	}
	footer, err := flags.GetString(footerFlag)
	if err != nil {
		return err
	}

	branch, branchErr := gitBranch()
	commitType, err := parseCommitType(commitTypeFlagVal, branch, branchErr)
	if err != nil {
		return err
	}
	if scope == "" {
		logger.Printf("Scope flag not provided, attempting to parse")
		scope, err = parseScope(branch, branchErr)
		if err != nil {
			return err
		}
	}

	message := parseMessage(commitType, scope, subject, body, footer)
	fmt.Printf("%s\n", message)
	return nil
}
