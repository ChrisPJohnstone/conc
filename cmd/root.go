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
	flags := rootCmd.Flags()
	flags.StringP(
		typeFlag,
		string(typeFlag[0]),
		typeFlagDefault,
		typeFlagDescription,
	)
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
