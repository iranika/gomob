package main

import (
	"flag"
	"fmt"

	"github.com/iranika/gomob"
)

//gomob-cli tools entry
func main() {
	var (
		//getDeck  = flag.String("getdeck", "", "")
		getSales = flag.String("getsales", "undefined", "")
	)
	flag.Parse()
	if *getSales != "undefined" {
		sales := gomob.GetSalesInfo(*getSales)
		fmt.Println(sales)
	}
	//gomob.GetBattleResult("Get battle")
}
