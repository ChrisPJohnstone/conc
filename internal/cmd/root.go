package cmd

import (
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/ChrisPJohnstone/conc/internal/globals"

	"github.com/spf13/cobra"
)

var rootCmd *cobra.Command = &cobra.Command{
	Use:               cmdName,
	Short:             cmdShortHelp,
	Args:              cobra.MinimumNArgs(1),
	PersistentPreRunE: setupLogging,
	RunE:              main,
}

func init() {
	rootCmd.PersistentFlags().BoolP(
		verboseFlag,
		string(verboseFlag[0]),
		false,
		"enable verbose logging",
	)
	flags := rootCmd.Flags()
	flags.StringP(
		commitTypeFlag,
		string(commitTypeFlag[0]),
		"",
		fmt.Sprintf(
			"commit type\noptions: %s\nif unset will attempt to detect from branch\nfallback is '%s'",
			strings.Join(globals.CommitTypes, ", "),
			globals.CommitTypeFallback,
		),
	)
	flags.StringP(
		scopeFlag,
		string(scopeFlag[0]),
		scopeFlagDefault,
		fmt.Sprintf(
			"scope for commit message\nif unset will attempt to detect from branch\nuse '%s' to disable",
			globals.ScopeFlagSentinel,
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

func setupLogging(cmd *cobra.Command, args []string) error {
	flags := cmd.Flags()
	verbose, err := flags.GetBool(verboseFlag)
	if err != nil {
		return err
	}
	if verbose {
		opts := &slog.HandlerOptions{Level: slog.LevelDebug}
		handler := slog.NewJSONHandler(os.Stderr, opts)
		logger := slog.New(handler)
		slog.SetDefault(logger)
	}
	slog.Debug("Verbose logging enabled")
	return nil
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
