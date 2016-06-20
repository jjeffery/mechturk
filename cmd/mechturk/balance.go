package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws/awsutil"
	mechturk "github.com/jjeffery/mechturk"
	"github.com/spf13/cobra"
)

type balanceOpts struct {
	dummy string
}

func (opts *balanceOpts) Run(cmd *cobra.Command, args []string) {
	client := mechturk.New(nil)
	resp, err := client.GetAccountBalance(&mechturk.GetAccountBalanceRequest{
		ResponseGroup: getResponseGroups(),
	})
	if err != nil {
		log.Fatal("error: ", err)
	}
	result := resp.GetAccountBalanceResult
	fmt.Println(awsutil.Prettify(result))
}

func init() {
	opts := &balanceOpts{}
	cmd := &cobra.Command{
		Use:     "get-account-balance",
		Aliases: []string{"balance", "bal"},
		Short:   "get account balance",
		Run:     opts.Run,
		PreRun:  applyGlobalOptions,
	}
	cmd.PersistentFlags().StringVarP(&opts.dummy, "dummy", "d", "", "dummy option")
	rootCommand.AddCommand(cmd)
}
