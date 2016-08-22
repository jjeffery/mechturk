package mechturk_test

import (
	"log"

	"github.com/jjeffery/mechturk"
)

func ExampleMechTurk_GetAccountBalance() {
	mt := mechturk.New()
	result, err := mt.GetAccountBalance(&mechturk.GetAccountBalanceRequest{})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(mechturk.Prettify(result))
}
