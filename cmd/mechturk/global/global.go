package global

import (
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"github.com/jjeffery/mechturk"
	"github.com/jjeffery/mechturk/credentials"
)

const ProgramName = "mechturk"

var (
	Verbose        bool
	Sandbox        bool
	ResponseGroups []string
	Command        *cobra.Command
)

func init() {
	Command = &cobra.Command{
		Use:              ProgramName,
		Short:            "AWS Mechanical Turk Command Line Interface",
		Run:              run,
		PersistentPreRun: persistentPreRun,
	}
	Command.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose mode")
	Command.PersistentFlags().BoolVarP(&Sandbox, "sandbox", "s", false, "use sandbox")
	Command.PersistentFlags().StringSliceVarP(&ResponseGroups,
		"response-groups", "r",
		[]string{"Minimal"}, "response group(s)")
}

func persistentPreRun(cmd *cobra.Command, arg []string) {
	if Verbose {
		log.SetFlags(log.Flags() | log.Lshortfile)
		logger := log.New(os.Stdout, "", log.Flags())
		mechturk.DefaultConfig.Logger = logger
		credentials.Logger = logger
	}
	if Sandbox {
		mechturk.DefaultConfig = mechturk.DefaultConfig.WithSandbox(true)
	}
}

func run(cmd *cobra.Command, args []string) {
	log.Println(cmd.Short)
	ShortUsage()
}

func ShortUsage() {
	log.Printf(`type "%s --help" for command line usage`, ProgramName)
}

// GetResponseGroups returns a slice of response groups selected
// for the operation. Because the valid response groups depend on
// the type of operation, this
func GetResponseGroups(allowable ...string) []string {
	// Set of valid response groups, initialized with
	// values that are valid for all operations.
	validSet := map[string]bool{
		"Minimal": true,
		"Request": true,
	}
	validList := []string{
		"Minimal",
		"Request",
	}
	selectedSet := map[string]bool{}
	selectedList := []string{}
	for _, rg := range allowable {
		if !validSet[rg] {
			validSet[rg] = true
			validList = append(validList, rg)
		}
	}
	for _, rg := range ResponseGroups {
		rg = strings.TrimSpace(rg)
		if !validSet[rg] {
			log.Fatalf("error: invalid response group %q: valid values are %q",
				rg, strings.Join(validList, ","))
		}
		if !selectedSet[rg] {
			selectedSet[rg] = true
			selectedList = append(selectedList, rg)
		}
	}
	if len(selectedList) == 0 {
		selectedList = []string{"Minimal"}
	}
	return selectedList
}
