package searchhits

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"github.com/jjeffery/mechturk"
	"github.com/jjeffery/mechturk/cmd/mechturk/global"
	"github.com/jjeffery/mechturk/cmd/mechturk/mtcmd"
)

var (
	sortDirection string
	sortProperty  string
	pageNumber    int
	pageSize      int

	sortDirections mtcmd.StringSelection
	sortProperties mtcmd.StringSelection
)

func init() {
	command := &cobra.Command{
		Use:     "search-hits",
		Aliases: []string{"search"},
		Short:   "search HITs",
		Run:     run,
	}
	command.PersistentFlags().StringVar(&sortDirection, "sort-direction", "", "sort direction")
	command.PersistentFlags().StringVar(&sortProperty, "sort-property", "", "sort property")
	command.PersistentFlags().IntVar(&pageNumber, "page-number", 0, "page number")
	command.PersistentFlags().IntVar(&pageSize, "page-size", 0, "page size")
	global.Command.AddCommand(command)

	sortDirections = mtcmd.NewStringSelection("sort direction",
		mechturk.SortDirectionAscending,
		mechturk.SortDirectionDescending)
	sortProperties = mtcmd.NewStringSelection("sort property",
		mechturk.SearchHITsSortPropertyTitle,
		mechturk.SearchHITsSortPropertyReward,
		mechturk.SearchHITsSortPropertyExpiration,
		mechturk.SearchHITsSortPropertyCreationTime,
		mechturk.SearchHITsSortPropertyEnumeration)
}

func run(cmd *cobra.Command, args []string) {
	client := mechturk.New()
	result, err := client.SearchHITs(&mechturk.SearchHITsRequest{
		SortDirection: getSortDirection(),
		SortProperty:  getSortProperty(),
		PageNumber:    pageNumber,
		PageSize:      pageSize,
		ResponseGroup: global.GetResponseGroups(
			"HITDetail",
			"HITQuestion",
			"HITAssignmentSummary",
		),
	})
	if err != nil {
		log.Fatal("error: ", err)
	}
	fmt.Println(mechturk.Prettify(result))
}

func getSortDirection() *mechturk.SortDirection {
	return (*mechturk.SortDirection)(sortDirections.Select(sortDirection))
}

func getSortProperty() *mechturk.SearchHITsSortProperty {
	return (*mechturk.SearchHITsSortProperty)(sortProperties.Select(sortProperty))
}
