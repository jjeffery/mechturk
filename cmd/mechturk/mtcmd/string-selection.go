package mtcmd

import (
	"log"
	"reflect"
	"strings"
)

type StringSelection struct {
	Name  string
	Valid []string
}

func NewStringSelection(name string, valid ...interface{}) StringSelection {
	ss := StringSelection{
		Name:  name,
		Valid: make([]string, len(valid)),
	}
	for i, vv := range valid {
		ss.Valid[i] = reflect.ValueOf(vv).String()
	}
	return ss
}

func (ss StringSelection) Select(value string) *string {
	value = strings.TrimSpace(value)
	if value == "" {
		// empty string means not specified
		return nil
	}

	for _, v := range ss.Valid {
		if value == v {
			return &value
		}
	}

	log.Fatalf("error: invalid value for %s: %q\n       valid values: %q",
		ss.Name, value, strings.Join(ss.Valid, ","))
	panic("not reached")
}
