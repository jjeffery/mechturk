package main

import (
	"fmt"
	"log"

	"github.com/jjeffery/mechturk"
	"github.com/spf13/cobra"
)

var balance = newBalance(global)

type balanceT struct {
	dummy   string
	command *cobra.Command
}

func newBalance(g *globalT) *balanceT {
	b := &balanceT{}
	b.command = &cobra.Command{
		Use:     "get-account-balance",
		Aliases: []string{"balance", "bal"},
		Short:   "get account balance",
		Run:     b.Run,
	}
	b.command.PersistentFlags().StringVarP(&b.dummy, "dummy", "d", "", "dummy option")
	g.command.AddCommand(b.command)
	return b
}

func (b *balanceT) Run(cmd *cobra.Command, args []string) {
	client := mechturk.New(nil)
	resp, err := client.GetAccountBalance(&mechturk.GetAccountBalanceRequest{
		ResponseGroup: commonGetResponseGroups(),
	})
	if err != nil {
		log.Fatal("error: ", err)
	}
	result := resp.GetAccountBalanceResult
	fmt.Println(mechturk.Prettify(result))
}
