package cmd

import (
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/go-git/go-git/v6"
	"github.com/spf13/cobra"
)

var rootCmd *cobra.Command = &cobra.Command{
	Use:   cmdName,
	Short: cmdShortHelp,
	Args:  cobra.MinimumNArgs(1),
	RunE:  main,
}

func isValidCommitType(commitType string) bool {
	return slices.Contains(commitTypes, commitType)
}

func init() {
	flags := rootCmd.Flags()
	flags.StringP(
		commitTypeFlag,
		string(commitTypeFlag[0]),
		commitTypeFlagDefault,
		fmt.Sprintf("commit type (valid: %s)", strings.Join(commitTypes, ", ")),
	)
}

func gitBranch() (string, error) {
	repo, err := git.PlainOpen(".")
	if err != nil {
		return "", err
	}
	ref, err := repo.Head()
	if err != nil {
		return "", err
	}
	branch := ref.Name().Short()
	return branch, nil
}

func main(cmd *cobra.Command, args []string) error {
	flags := cmd.Flags()
	commitType, err := flags.GetString("type")
	if err != nil {
		return err
	}
	if !isValidCommitType(commitType) {
		return fmt.Errorf("%s is not a valid commit type", commitType)
	}
	subject := strings.Join(args, " ")
	fmt.Printf("%s: %s\n", commitType, subject)
	branch, err := gitBranch()
	if err != nil {
		return err
	}
	fmt.Printf("%s\n", branch)
	return nil
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
