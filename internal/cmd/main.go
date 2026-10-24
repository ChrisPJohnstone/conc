package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/ChrisPJohnstone/conc/internal/message"

	"github.com/spf13/cobra"
)

func main(cmd *cobra.Command, args []string) error {
	flags := cmd.Flags()

	var err error
	msg := &message.Message{}
	msg.Subject = strings.Join(args, " ")
	log.Printf("Raw subject: '%s'\n", msg.Subject)
	msg.CommitType, err = flags.GetString(commitTypeFlag)
	if err != nil {
		return err
	}
	log.Printf("Raw commitType: '%s'\n", msg.CommitType)
	msg.Scope, err = flags.GetString(scopeFlag)
	if err != nil {
		return err
	}
	log.Printf("Raw scope: '%s'\n", msg.Scope)
	msg.Body, err = flags.GetString(bodyFlag)
	if err != nil {
		return err
	}
	log.Printf("Raw body: '%s'\n", msg.Body)
	msg.Footer, err = flags.GetString(footerFlag)
	if err != nil {
		return err
	}
	log.Printf("Raw footer: '%s'\n", msg.Footer)
	err = msg.Parse()
	if err != nil {
		return err
	}
	log.Printf("Full message: '%s'\n", msg)

	fmt.Printf("%s\n", msg)
	return nil
}
