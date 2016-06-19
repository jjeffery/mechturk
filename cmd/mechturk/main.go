package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jjeffery/mturk"
	"github.com/spf13/cobra"
)

const programName = "mechturk"

var globalOptions struct {
	verbose bool
	sandbox bool
}

var rootCommand = &cobra.Command{
	Use:   programName,
	Short: "AWS Mechanical Turk Command Line Client",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(`type "` + programName + ` --help" for command line usage`)
	},
}

func main() {
	if err := rootCommand.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCommand.PersistentFlags().BoolVarP(&globalOptions.verbose, "verbose", "v", false, "verbose mode")
	rootCommand.PersistentFlags().BoolVarP(&globalOptions.sandbox, "sandbox", "s", false, "use sandbox")
}

func applyGlobalOptions(cmd *cobra.Command, args []string) {
	if globalOptions.verbose {
		mturk.DefaultConfig.Logger = log.New(os.Stdout, "", log.Ltime|log.Lshortfile)
	}
	if globalOptions.sandbox {
		mturk.DefaultConfig = mturk.DefaultConfig.WithSandbox(true)
	}
}
