package poker

import "fmt"

var FIVE_HAND [2598960][5]uint8    // combinations (5 among 52)
var SEVEN_HAND [133784560][7]uint8 // combinations (7 among 52)

func BuildFiveHands(verbose bool) {

	if verbose {
		fmt.Println(" ")
		defer track(runningtime("BuildFiveHands"))
	}

	var c1, c2, c3, c4, c5 uint8
	var cards [5]uint8
	c := 0

	for c1 = 0; c1 < DECK_SIZE; c1++ {
		for c2 = 0; c2 < c1; c2++ {
			for c3 = 0; c3 < c2; c3++ {
				for c4 = 0; c4 < c3; c4++ {
					for c5 = 0; c5 < c4; c5++ {
						cards = [5]uint8{c1, c2, c3, c4, c5}
						FIVE_HAND[c] = cards
						c += 1
					}
				}
			}
		}
	}
}

func BuildSevenHands(verbose bool) {

	if verbose {
		fmt.Println(" ")
		defer track(runningtime("BuildSevenHands"))
	}

	var c1, c2, c3, c4, c5, c6, c7 uint8
	var cards [7]uint8
	c := 0

	for c1 = 0; c1 < DECK_SIZE; c1++ {
		for c2 = 0; c2 < c1; c2++ {
			for c3 = 0; c3 < c2; c3++ {
				for c4 = 0; c4 < c3; c4++ {
					for c5 = 0; c5 < c4; c5++ {
						for c6 = 0; c6 < c5; c6++ {
							for c7 = 0; c7 < c6; c7++ {
								cards = [7]uint8{c1, c2, c3, c4, c5}
								SEVEN_HAND[c] = cards
								c += 1
							}
						}
					}
				}
			}
		}
	}
}
