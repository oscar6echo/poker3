package main

import "github.com/oscar6echo/poker3/poker"

func main() {

	verbose := true
	poker.Setup(verbose)

	poker.BuildFiveHands(verbose)
	poker.BuildSevenHands(verbose)

}
