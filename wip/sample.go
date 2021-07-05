package main

import "github.com/oscar6echo/poker3/poker"

func main() {

	poker.Setup(false)

	poker.BuildFiveHandStats(true)
	// poker.BuildSevenHandStats(true)

	poker.TryCalcEquity()
	poker.TryCalcEquityMonteCarlo()

}
