package cmd

import (
	"fmt"
	"os"
	"slices"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "conc",
	Short: "Conventional Commit Helper",
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
		commitTypeFlagDescription,
	)
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
	fmt.Printf("Type %s\n", commitType)
	fmt.Printf("%s\n", args)
	return nil
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
