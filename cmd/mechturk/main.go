package main

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"

	"github.com/jjeffery/mechturk"
	"github.com/jjeffery/mechturk/credentials"
	"github.com/spf13/cobra"
)

const programName = "mechturk"

var globalOptions struct {
	verbose        bool
	sandbox        bool
	responseGroups []string
}

var rootCommand = &cobra.Command{
	Use:   programName,
	Short: "AWS Mechanical Turk Command Line Client",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(`type "` + programName + ` --help" for command line usage`)
	},
}

func main() {
	log.SetFlags(0)
	if err := rootCommand.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCommand.PersistentFlags().BoolVarP(&globalOptions.verbose, "verbose", "v", false, "verbose mode")
	rootCommand.PersistentFlags().BoolVarP(&globalOptions.sandbox, "sandbox", "s", false, "use sandbox")
	rootCommand.PersistentFlags().StringSliceVarP(&globalOptions.responseGroups, "response-groups", "r", []string{"Minimal"}, "response group(s)")
}

func applyGlobalOptions(cmd *cobra.Command, args []string) {
	if globalOptions.verbose {
		log.SetFlags(log.Ltime | log.Lshortfile)
		logger := log.New(os.Stdout, "", log.Flags())
		mechturk.DefaultConfig.Logger = logger
		credentials.Logger = logger
	}
	if globalOptions.sandbox {
		mechturk.DefaultConfig = mechturk.DefaultConfig.WithSandbox(true)
	}
}

// getResponseGroups returns a slice of response groups selected
// for the operation. Because the valid response groups depend on
// the type of operation, this
func getResponseGroups(allowable ...string) []string {
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
	for _, rg := range globalOptions.responseGroups {
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

// getString chooses a string from the list of valid values, which
// might be types based on string.
func getString(name string, value string, valid ...interface{}) *string {
	value = strings.TrimSpace(value)
	if value == "" {
		// empty string means not specified
		return nil
	}

	validList := []string{}
	for _, v := range valid {
		vs := reflect.ValueOf(v).String()
		if vs == value {
			// valid value
			return &value
		}
		validList = append(validList, vs)
	}
	log.Fatalf("error: invalid value for %s: %q\n       valid values: %q",
		name, value, strings.Join(validList, ","))
	panic("not reached")
}
