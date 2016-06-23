package main

import (
	"log"
	"os"

	"github.com/jjeffery/mechturk"
	"github.com/jjeffery/mechturk/credentials"
	"github.com/spf13/cobra"
)

const programName = "mechturk"

var global = newGlobal()

type globalT struct {
	verbose        bool
	sandbox        bool
	responseGroups []string
	command        *cobra.Command
}

func newGlobal() *globalT {
	g := &globalT{}
	g.command = &cobra.Command{
		Use:              programName,
		Short:            "AWS Mechanical Turk Command Line Interface",
		Run:              g.Run,
		PersistentPreRun: g.PersistentPreRun,
	}
	g.command.PersistentFlags().BoolVarP(&g.verbose, "verbose", "v", false, "verbose mode")
	g.command.PersistentFlags().BoolVarP(&g.sandbox, "sandbox", "s", false, "use sandbox")
	g.command.PersistentFlags().StringSliceVarP(&g.responseGroups,
		"response-groups", "r",
		[]string{"Minimal"}, "response group(s)")
	return g
}

func (g *globalT) PersistentPreRun(cmd *cobra.Command, arg []string) {
	log.SetFlags(0)
	if g.verbose {
		log.SetFlags(log.Ltime | log.Lshortfile)
		logger := log.New(os.Stdout, "", log.Flags())
		mechturk.DefaultConfig.Logger = logger
		credentials.Logger = logger
	}
	if g.sandbox {
		mechturk.DefaultConfig = mechturk.DefaultConfig.WithSandbox(true)
	}
}

func (g *globalT) Run(cmd *cobra.Command, args []string) {
	log.Println(cmd.Short)
	commonShortUsage()
}

func (g *globalT) Execute() error {
	return g.command.Execute()
}
