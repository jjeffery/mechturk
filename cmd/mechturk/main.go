package main

import (
	"log"

	"github.com/jjeffery/mechturk/cmd/mechturk/global"

	_ "github.com/jjeffery/mechturk/cmd/mechturk/balance"
	_ "github.com/jjeffery/mechturk/cmd/mechturk/disablehit"
	_ "github.com/jjeffery/mechturk/cmd/mechturk/searchhits"
)

func main() {
	log.SetFlags(0)
	if err := global.Command.Execute(); err != nil {
		log.Fatal(err)
	}
}
