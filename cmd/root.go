package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "conc",
	Short: "Conventional Commit Helper",
	Run:   main,
}

func init() {
}

func main(cmd *cobra.Command, args []string) {
	fmt.Printf("Hello World!\n")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
