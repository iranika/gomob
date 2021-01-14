package main

import (
	"flag"

	"github.com/iranika/gomob"
)

//gomob-cli tools entry
func main() {
	var (
		getDeck = flag.String("getdeck", "", "")
	)
	flag.Parse()
	gomob.GetBattleResult("Get battle")
}
