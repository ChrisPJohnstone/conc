package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var rootCmd *cobra.Command = &cobra.Command{
	Use:   cmdName,
	Short: cmdShortHelp,
	Args:  cobra.MinimumNArgs(1),
	RunE:  main,
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

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
