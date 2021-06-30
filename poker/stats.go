package poker

import "fmt"

type handTypeStatsStruct struct {
	nbHand      int
	minRank     int
	maxRank     int
	nbOccurence int
}
type handStatsStruct map[string]handTypeStatsStruct

var FiveHandTypeStats = make(map[string]*handTypeStatsStruct)
var SevenHandTypeStats = make(map[string]*handTypeStatsStruct)

func BuildFiveHandStats(verbose bool) {

	if verbose {
		fmt.Println(" ")
		defer track(runningtime("BuildFiveHandStats"))
	}

	for k := range FiveHandTypeStats {
		delete(FiveHandTypeStats, k)
	}

	handTypeStats := FiveHandTypeStats
	var rankCount = make(map[int]int)

	var c1, c2, c3, c4, c5 int
	var cards [5]int
	var rank int

	for c1 = 0; c1 < DECK_SIZE; c1++ {
		for c2 = 0; c2 < c1; c2++ {
			for c3 = 0; c3 < c2; c3++ {
				for c4 = 0; c4 < c3; c4++ {
					for c5 = 0; c5 < c4; c5++ {
						cards = [5]int{c1, c2, c3, c4, c5}
						rank = GetRankFive(cards)
						rankCount[rank] += 1
					}
				}
			}
		}
	}

	for rank, nbOccur := range rankCount {
		handType := HAND_TYPE[rank]
		_, present := handTypeStats[handType]
		if !present {
			handTypeStats[handType] = &handTypeStatsStruct{nbHand: 0, minRank: rank, maxRank: rank, nbOccurence: 0}
		}
		obj := handTypeStats[handType]
		obj.nbHand += 1
		obj.nbOccurence += nbOccur
		obj.minRank = min(obj.minRank, rank)
		obj.maxRank = max(obj.maxRank, rank)
	}

	if verbose {
		fmt.Printf("stats five cards\n")
		for k, v := range handTypeStats {
			fmt.Printf("\tstats=%v\thand-type=%s\n", *v, k)
		}
	}

}

func BuildSevenHandStats(verbose bool) {

	if verbose {
		fmt.Println(" ")
		defer track(runningtime("BuildSevenHandStats"))
	}

	for k := range SevenHandTypeStats {
		delete(SevenHandTypeStats, k)
	}

	handTypeStats := SevenHandTypeStats
	var rankCount = make(map[int]int)

	var c1, c2, c3, c4, c5, c6, c7 int
	var cards [7]int
	var rank int

	for c1 = 0; c1 < DECK_SIZE; c1++ {
		for c2 = 0; c2 < c1; c2++ {
			for c3 = 0; c3 < c2; c3++ {
				for c4 = 0; c4 < c3; c4++ {
					for c5 = 0; c5 < c4; c5++ {
						for c6 = 0; c6 < c5; c6++ {
							for c7 = 0; c7 < c6; c7++ {
								cards = [7]int{c1, c2, c3, c4, c5, c6, c7}
								rank = GetRank(cards)
								rankCount[rank] += 1
							}
						}
					}
				}
			}
		}
	}

	for rank, nbOccur := range rankCount {
		handType := HAND_TYPE[rank]
		_, present := handTypeStats[handType]
		if !present {
			handTypeStats[handType] = &handTypeStatsStruct{nbHand: 0, minRank: rank, maxRank: rank, nbOccurence: 0}
		}
		obj := handTypeStats[handType]
		obj.nbHand += 1
		obj.nbOccurence += nbOccur
		obj.minRank = min(obj.minRank, rank)
		obj.maxRank = max(obj.maxRank, rank)
	}

	if verbose {
		fmt.Printf("stats seven cards\n")
		for k, v := range handTypeStats {
			fmt.Printf("\tstats=%v\thand-type=%s\n", *v, k)
		}
	}

}
