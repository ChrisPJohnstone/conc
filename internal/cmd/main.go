package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

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
