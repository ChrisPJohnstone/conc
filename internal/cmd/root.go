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
	flags.BoolP(
		verboseFlag,
		string(verboseFlag[0]),
		false,
		"enable verbose logging",
	)
	flags.StringP(
		commitTypeFlag,
		string(commitTypeFlag[0]),
		"",
		fmt.Sprintf(
			"commit type\noptions: %s\nif unset will attempt to detect from branch\nfallback is '%s'",
			strings.Join(commitTypes, ", "),
			commitTypeFallback,
		),
	)
	flags.StringP(
		scopeFlag,
		string(scopeFlag[0]),
		scopeFlagDefault,
		fmt.Sprintf(
			"scope for commit message\nif unset will attempt to detect from branch\nuse '%s' to disable",
			scopeFlagSentinel,
		),
	)
	flags.StringP(
		bodyFlag,
		string(bodyFlag[0]),
		bodyFlagDefault,
		"add a body to the commit message",
	)
	flags.StringP(
		footerFlag,
		string(footerFlag[0]),
		footerFlagDefault,
		"add a footer to the commit message",
	)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
