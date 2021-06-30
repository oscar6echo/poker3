package main

import "github.com/oscar6echo/poker3/poker"

func main() {

	poker.Setup(true)

	// poker.BuildFiveHands(true)
	// poker.BuildSevenHands(true)

	poker.BuildFiveHandStats(true)
	poker.BuildSevenHandStats(true)

}
