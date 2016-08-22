package getassignmentsforhit

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
	pageNumber    int32
	pageSize      int32

	sortDirections mtcmd.StringSelection
	sortProperties mtcmd.StringSelection
)

func init() {
	command := &cobra.Command{
		Use:     "get-assignments-for-hit",
		Aliases: []string{"assignments"},
		Short:   "get assignments for HIT",
		Run:     run,
	}
	command.PersistentFlags().StringVar(&sortDirection, "sort-direction", "", "sort direction")
	command.PersistentFlags().StringVar(&sortProperty, "sort-property", "", "sort property")
	command.PersistentFlags().Int32Var(&pageNumber, "page-number", 0, "page number")
	command.PersistentFlags().Int32Var(&pageSize, "page-size", 0, "page size")
	global.Command.AddCommand(command)

	sortDirections = mtcmd.NewStringSelection("sort direction",
		mechturk.SortDirectionAscending,
		mechturk.SortDirectionDescending)
	sortProperties = mtcmd.NewStringSelection("sort property",
		mechturk.GetAssignmentsForHITSortPropertyAcceptTime,
		mechturk.GetAssignmentsForHITSortPropertyAnswer,
		mechturk.GetAssignmentsForHITSortPropertyAssignmentStatus,
		mechturk.GetAssignmentsForHITSortPropertySubmitTime,
	)
}

func run(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		log.Fatalf("usage: %s get-assignments-for-hit <HIT-ID>", global.ProgramName)
	}
	hitId := args[0]
	client := mechturk.New()
	resp, err := client.GetAssignmentsForHIT(&mechturk.GetAssignmentsForHITRequest{
		SortDirection: getSortDirection(),
		SortProperty:  getSortProperty(),
		PageNumber:    pageNumber,
		PageSize:      pageSize,
		ResponseGroup: global.GetResponseGroups(
			"AssignmentFeedback",
		),
		HITId: hitId,
	})
	if err != nil {
		log.Fatal("error: ", err)
	}
	result := resp.GetAssignmentsForHITResult
	fmt.Println(mechturk.Prettify(result))
}

func getSortDirection() *mechturk.SortDirection {
	return (*mechturk.SortDirection)(sortDirections.Select(sortDirection))
}

func getSortProperty() *mechturk.GetAssignmentsForHITSortProperty {
	return (*mechturk.GetAssignmentsForHITSortProperty)(sortProperties.Select(sortProperty))
}
