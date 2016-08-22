package disposehit

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"github.com/jjeffery/mechturk"
	"github.com/jjeffery/mechturk/cmd/mechturk/global"
)

func init() {
	command := &cobra.Command{
		Use:     "dispose-hit",
		Aliases: []string{"dispose"},
		Short:   "dispose HIT",
		Run:     run,
	}
	global.Command.AddCommand(command)
}

func run(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		log.Fatalf("usage: %s dispose-hit <HIT-ID>", global.ProgramName)
	}
	client := mechturk.New()
	result, err := client.DisposeHIT(&mechturk.DisposeHITRequest{
		HITId:         args[0],
		ResponseGroup: global.GetResponseGroups(),
	})
	if err != nil {
		log.Fatal("error: ", err)
	}
	fmt.Println(mechturk.Prettify(result))
}
