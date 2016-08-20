package balance

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"github.com/jjeffery/mechturk"
	"github.com/jjeffery/mechturk/cmd/mechturk/global"
)

func init() {
	command := &cobra.Command{
		Use:     "get-account-balance",
		Aliases: []string{"balance", "bal"},
		Short:   "get account balance",
		Run:     run,
	}
	global.Command.AddCommand(command)
}

func run(cmd *cobra.Command, args []string) {
	client := mechturk.New(nil)
	resp, err := client.GetAccountBalance(&mechturk.GetAccountBalanceRequest{
		ResponseGroup: global.GetResponseGroups(),
	})
	if err != nil {
		log.Fatal("error: ", err)
	}
	result := resp.GetAccountBalanceResult
	fmt.Println(mechturk.Prettify(result))
}
