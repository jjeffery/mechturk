package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws/awsutil"
	mechturk "github.com/jjeffery/mechturk"
	"github.com/spf13/cobra"
)

type searchHitsOpts struct {
	sortDirection string
	sortProperty  string
	pageNumber    int32
	pageSize      int32

	sortDirections stringSelection
	sortProperties stringSelection
}

func (opts *searchHitsOpts) Run(cmd *cobra.Command, args []string) {
	client := mechturk.New()
	resp, err := client.SearchHITs(&mechturk.SearchHITsRequest{
		SortDirection: opts.getSortDirection(),
		SortProperty:  opts.getSortProperty(),
		PageNumber:    opts.pageNumber,
		PageSize:      opts.pageSize,
		ResponseGroup: getResponseGroups(),
	})
	if err != nil {
		log.Fatal("error: ", err)
	}
	result := resp.SearchHITsResult
	fmt.Println(awsutil.Prettify(result))
}

func init() {
	opts := &searchHitsOpts{}
	cmd := &cobra.Command{
		Use:     "search-hits",
		Aliases: []string{"search"},
		Short:   "search HITs",
		Run:     opts.Run,
		PreRun:  applyGlobalOptions,
	}
	cmd.PersistentFlags().StringVar(&opts.sortDirection, "sort-direction", "", "sort direction")
	cmd.PersistentFlags().StringVar(&opts.sortProperty, "sort-property", "", "sort property")
	cmd.PersistentFlags().Int32Var(&opts.pageNumber, "page-number", 0, "page number")
	cmd.PersistentFlags().Int32Var(&opts.pageSize, "page-size", 0, "page size")
	rootCommand.AddCommand(cmd)

	opts.sortDirections = newStringSelection("sort direction",
		mechturk.SortDirectionAscending,
		mechturk.SortDirectionDescending)
	opts.sortProperties = newStringSelection("sort property",
		mechturk.SearchHITsSortPropertyTitle,
		mechturk.SearchHITsSortPropertyReward,
		mechturk.SearchHITsSortPropertyExpiration,
		mechturk.SearchHITsSortPropertyCreationTime,
		mechturk.SearchHITsSortPropertyEnumeration)
}

func (opts *searchHitsOpts) getSortDirection() *mechturk.SortDirection {
	return (*mechturk.SortDirection)(opts.sortDirections.Select(opts.sortDirection))
}

func (opts *searchHitsOpts) getSortProperty() *mechturk.SearchHITsSortProperty {
	return (*mechturk.SearchHITsSortProperty)(opts.sortProperties.Select(opts.sortProperty))
}
