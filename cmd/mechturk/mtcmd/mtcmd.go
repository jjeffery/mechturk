// Package mtcmd provide support utilities for mechturk commands.
package mtcmd

import (
	"log"
	"reflect"
	"strings"
)

// getString chooses a string from the list of valid values, which
// might be types based on string.
func GetString(name string, value string, valid ...interface{}) *string {
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
