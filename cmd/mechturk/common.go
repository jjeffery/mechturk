package main

import (
	"log"
	"reflect"
	"strings"
)

func commonShortUsage() {
	log.Printf(`type "%s --help" for command line usage`, programName)
}

// getResponseGroups returns a slice of response groups selected
// for the operation. Because the valid response groups depend on
// the type of operation, this
func commonGetResponseGroups(allowable ...string) []string {
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
	for _, rg := range global.responseGroups {
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
func commonGetString(name string, value string, valid ...interface{}) *string {
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
