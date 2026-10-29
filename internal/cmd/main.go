package cmd

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/ChrisPJohnstone/git-conc/v2/internal/git"
	"github.com/ChrisPJohnstone/git-conc/v2/internal/message"

	"github.com/spf13/cobra"
)

func main(cmd *cobra.Command, args []string) error {
	flags := cmd.Flags()

	var err error
	msg := &message.Message{}
	msg.Subject = strings.Join(args, " ")
	slog.Debug(fmt.Sprintf("Raw subject: '%s'", msg.Subject))
	msg.Body, err = flags.GetString(bodyFlag)
	if err != nil {
		return err
	}
	slog.Debug(fmt.Sprintf("Raw body: '%s'", msg.Body))
	msg.Breaking, err = flags.GetBool(breakingFlag)
	if err != nil {
		return err
	}
	slog.Debug(fmt.Sprintf("Breaking Change: '%t'", msg.Breaking))
	msg.CommitType, err = flags.GetString(commitTypeFlag)
	if err != nil {
		return err
	}
	slog.Debug(fmt.Sprintf("Raw commitType: '%s'", msg.CommitType))
	msg.Footer, err = flags.GetString(footerFlag)
	if err != nil {
		return err
	}
	slog.Debug(fmt.Sprintf("Raw footer: '%s'", msg.Footer))
	msg.Scope, err = flags.GetString(scopeFlag)
	if err != nil {
		return err
	}
	slog.Debug(fmt.Sprintf("Raw scope: '%s'", msg.Scope))

	err = msg.Parse()
	if err != nil {
		return err
	}
	slog.Debug(fmt.Sprintf("Full message: '%s'", msg))

	return git.GitCommit(msg.String())
}
