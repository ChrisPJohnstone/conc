package cmd

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/ChrisPJohnstone/conc/internal/message"

	"github.com/spf13/cobra"
)

func main(cmd *cobra.Command, args []string) error {
	flags := cmd.Flags()

	var err error
	msg := &message.Message{}
	msg.Subject = strings.Join(args, " ")
	slog.Debug(fmt.Sprintf("Raw subject: '%s'", msg.Subject))
	msg.CommitType, err = flags.GetString(commitTypeFlag)
	if err != nil {
		return err
	}
	slog.Debug(fmt.Sprintf("Raw commitType: '%s'", msg.CommitType))
	msg.Scope, err = flags.GetString(scopeFlag)
	if err != nil {
		return err
	}
	slog.Debug(fmt.Sprintf("Raw scope: '%s'", msg.Scope))
	msg.Body, err = flags.GetString(bodyFlag)
	if err != nil {
		return err
	}
	slog.Debug(fmt.Sprintf("Raw body: '%s'", msg.Body))
	msg.Footer, err = flags.GetString(footerFlag)
	if err != nil {
		return err
	}
	slog.Debug(fmt.Sprintf("Raw footer: '%s'", msg.Footer))
	err = msg.Parse()
	if err != nil {
		return err
	}
	slog.Debug(fmt.Sprintf("Full message: '%s'", msg))

	fmt.Printf("%s\n", msg)
	return nil
}
