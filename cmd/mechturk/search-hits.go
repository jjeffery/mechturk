package main

import (
	"fmt"
	"log"

	"github.com/jjeffery/mechturk"
	"github.com/spf13/cobra"
)

var searchHits = newSearchHits(global)

type searchHitsT struct {
	command *cobra.Command

	sortDirection string
	sortProperty  string
	pageNumber    int32
	pageSize      int32

	sortDirections stringSelection
	sortProperties stringSelection
}

func newSearchHits(g *globalT) *searchHitsT {
	s := &searchHitsT{}
	s.command = &cobra.Command{
		Use:     "search-hits",
		Aliases: []string{"search"},
		Short:   "search HITs",
		Run:     s.Run,
	}
	s.command.PersistentFlags().StringVar(&s.sortDirection, "sort-direction", "", "sort direction")
	s.command.PersistentFlags().StringVar(&s.sortProperty, "sort-property", "", "sort property")
	s.command.PersistentFlags().Int32Var(&s.pageNumber, "page-number", 0, "page number")
	s.command.PersistentFlags().Int32Var(&s.pageSize, "page-size", 0, "page size")
	g.command.AddCommand(s.command)

	s.sortDirections = newStringSelection("sort direction",
		mechturk.SortDirectionAscending,
		mechturk.SortDirectionDescending)
	s.sortProperties = newStringSelection("sort property",
		mechturk.SearchHITsSortPropertyTitle,
		mechturk.SearchHITsSortPropertyReward,
		mechturk.SearchHITsSortPropertyExpiration,
		mechturk.SearchHITsSortPropertyCreationTime,
		mechturk.SearchHITsSortPropertyEnumeration)
	return s
}

func (s *searchHitsT) Run(cmd *cobra.Command, args []string) {
	client := mechturk.New()
	resp, err := client.SearchHITs(&mechturk.SearchHITsRequest{
		SortDirection: s.getSortDirection(),
		SortProperty:  s.getSortProperty(),
		PageNumber:    s.pageNumber,
		PageSize:      s.pageSize,
		ResponseGroup: commonGetResponseGroups(),
	})
	if err != nil {
		log.Fatal("error: ", err)
	}
	result := resp.SearchHITsResult
	fmt.Println(mechturk.Prettify(result))
}

func (s *searchHitsT) getSortDirection() *mechturk.SortDirection {
	return (*mechturk.SortDirection)(s.sortDirections.Select(s.sortDirection))
}

func (s *searchHitsT) getSortProperty() *mechturk.SearchHITsSortProperty {
	return (*mechturk.SearchHITsSortProperty)(s.sortProperties.Select(s.sortProperty))
}
