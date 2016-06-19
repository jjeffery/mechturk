package main

import (
	"fmt"
	"os"

	"github.com/jjeffery/mturk"
	"github.com/spf13/cobra"
)

type balanceOpts struct {
	dummy string
}

func (opts *balanceOpts) Run(cmd *cobra.Command, args []string) {
	client := mturk.NewClient(nil)
	resp, err := client.GetAccountBalance(&mturk.GetAccountBalanceRequest{
		ResponseGroup: []string{"Minimal", "Request"},
	})
	if err != nil {
		fmt.Println("cannot get account balance:", err)
		os.Exit(1)
	}
	result := resp.GetAccountBalanceResult
	if result.AvailableBalance != nil {
		fmt.Printf("available: %v %s (%s)", result.AvailableBalance.Amount, result.AvailableBalance.CurrencyCode, result.AvailableBalance.FormattedPrice)
	}
	if result.OnHoldBalance != nil {
		fmt.Printf("on hold: %v %s (%s)", result.OnHoldBalance.Amount, result.OnHoldBalance.CurrencyCode, result.OnHoldBalance.FormattedPrice)
	}
}

func init() {
	opts := &balanceOpts{}
	cmd := &cobra.Command{
		Use:     "balance",
		Aliases: []string{"bal"},
		Short:   "get account balance",
		Run:     opts.Run,
		PreRun:  applyGlobalOptions,
	}
	cmd.PersistentFlags().StringVarP(&opts.dummy, "dummy", "d", "", "dummy option")
	rootCommand.AddCommand(cmd)
}
