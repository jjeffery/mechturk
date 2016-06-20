package mechturk_test

import (
	"github.com/jjeffery/mechturk"
	"log"
)

func ExampleMechTurk_GetAccountBalance() {
	mt := mechturk.New()
	resp, err := mt.GetAccountBalance(&mechturk.GetAccountBalanceRequest{})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(mechturk.Prettify(resp))
}
