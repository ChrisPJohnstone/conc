package cmd

import (
	"fmt"
	"log"
	"os"
	"slices"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "conc",
	Short: "Conventional Commit Helper",
	Run:   main,
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

func main(cmd *cobra.Command, args []string) {
	flags := cmd.Flags()
	commitType, err := flags.GetString("type")
	if err != nil {
		log.Fatal(err)
	}
	if !isValidCommitType(commitType) {
		log.Fatalf("%s is not a valid commit type", commitType)
	}
	fmt.Printf("Type %s\n", commitType)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
