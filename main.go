package main

import "github.com/oscar6echo/poker3/poker"

func main() {

	poker.Init()
	poker.ShowKeys()

	poker.BuildEvalFiveTables()
	poker.BuildEvalSevenTables()

}
