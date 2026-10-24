package cmd

import (
	"fmt"
	"strings"

	"github.com/ChrisPJohnstone/conc/internal/message"

	"github.com/spf13/cobra"
)

func main(cmd *cobra.Command, args []string) error {
	flags := cmd.Flags()

	verbose, err := flags.GetBool(verboseFlag)
	if err != nil {
		return err
	}
	logger = newLogger(verbose)
	logger.Printf("Verbose logging enabled\n")

	msg := &message.Message{}
	msg.Subject = strings.Join(args, " ")
	logger.Printf("Raw subject: '%s'\n", msg.Subject)
	msg.CommitType, err = flags.GetString(commitTypeFlag)
	if err != nil {
		return err
	}
	logger.Printf("Raw commitType: '%s'\n", msg.CommitType)
	msg.Scope, err = flags.GetString(scopeFlag)
	if err != nil {
		return err
	}
	logger.Printf("Raw scope: '%s'\n", msg.Scope)
	msg.Body, err = flags.GetString(bodyFlag)
	if err != nil {
		return err
	}
	logger.Printf("Raw body: '%s'\n", msg.Body)
	msg.Footer, err = flags.GetString(footerFlag)
	if err != nil {
		return err
	}
	logger.Printf("Raw footer: '%s'\n", msg.Footer)
	err = msg.Parse()
	if err != nil {
		return err
	}
	logger.Printf("Full message: '%s'\n", msg)

	fmt.Printf("%s\n", msg)
	return nil
}
