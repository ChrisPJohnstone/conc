package cmd

import "fmt"

type message struct {
	subject    string
	commitType string
	scope      string
	body       string
	footer     string
}

func (msg *message) string() string {
	logger.Printf("Starting message with commit type '%s'\n", msg.scope)
	message := msg.commitType
	logger.Printf("Message: '%s'\n", message)
	if msg.scope != "" && msg.scope != scopeFlagSentinel {
		logger.Printf("Adding scope '%s' to message\n", msg.scope)
		message = fmt.Sprintf("%s(%s)", message, msg.scope)
		logger.Printf("Message: '%s'\n", message)
	}
	logger.Printf("Adding subject '%s' to message\n", msg.subject)
	message = fmt.Sprintf("%s: %s", message, msg.subject)
	logger.Printf("Message: '%s'\n", message)
	if msg.body != "" {
		logger.Printf("Adding body '%s' to message\n", msg.body)
		message = fmt.Sprintf("%s\n\n%s", message, msg.body)
		logger.Printf("Message: '%s'\n", message)
	}
	if msg.footer != "" {
		logger.Printf("Adding footer '%s' to message\n", msg.footer)
		message = fmt.Sprintf("%s\n\n%s", message, msg.footer)
		logger.Printf("Message: '%s'\n", message)
	}
	return message
}
