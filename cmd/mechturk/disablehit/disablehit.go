package disablehit

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"github.com/jjeffery/mechturk"
	"github.com/jjeffery/mechturk/cmd/mechturk/global"
)

func init() {
	command := &cobra.Command{
		Use:     "disable-hit",
		Aliases: []string{"disable"},
		Short:   "disable HIT",
		Run:     run,
	}
	global.Command.AddCommand(command)
}

func run(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		log.Fatalf("usage: %s disable-hit <HIT-ID>", global.ProgramName)
	}
	client := mechturk.New()
	result, err := client.DisableHIT(&mechturk.DisableHITRequest{
		HITId:         args[0],
		ResponseGroup: global.GetResponseGroups(),
	})
	if err != nil {
		log.Fatal("error: ", err)
	}
	fmt.Println(mechturk.Prettify(result))
}
