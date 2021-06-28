package main

import "github.com/oscar6echo/poker3/poker"

func main() {

	verbose := true

	poker.SetupKeys()
	poker.ShowKeys()

	poker.BuildEvalFiveTables(verbose)
	poker.BuildEvalSevenTables(verbose)

}
