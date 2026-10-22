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
	rootCmd.Flags().String("type", defaultType, "commit type")
}

func main(cmd *cobra.Command, args []string) {
	flags := cmd.Flags()
	typeValue, err := flags.GetString("type")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Type %s\n", typeValue)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
