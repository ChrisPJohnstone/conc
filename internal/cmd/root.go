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
		"",
		fmt.Sprintf(
			"commit type (valid: %s). If unset will attempt to detect from branch. Fallback is %s",
			strings.Join(commitTypes, ", "),
			commitTypeFallback,
		),
	)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
