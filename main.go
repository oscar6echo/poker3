package main

import (
	"syscall/js"

	"github.com/oscar6echo/poker3/poker"
)

func main() {
	c := make(chan struct{}, 0)

	println("Go WASM initialized")

	js.Global().Set("getRank", js.FuncOf(getRank))
	js.Global().Set("getRankFive", js.FuncOf(getRankFive))
	js.Global().Set("calcEquity", js.FuncOf(calcEquity))
	js.Global().Set("calcEquityMonteCarlo", js.FuncOf(calcEquityMonteCarlo))
	js.Global().Set("buildFiveHandStats", js.FuncOf(buildFiveHandStats))
	js.Global().Set("buildSevenHandStats", js.FuncOf(buildSevenHandStats))

	js.Global().Set("wasmDoc", js.FuncOf(wasmDoc))

	<-c
}

func wasmDoc(this js.Value, args []js.Value) interface{} {
	// return name of functions set to global scope in main()
	result := make([]interface{}, 0)
	var funcNames = []string{"getRank", "getRankFive", "buildFiveHandStats", "buildSevenHandStats"}
	for _, e := range funcNames {
		result = append(result, e)
	}
	return js.ValueOf(result)
}

func getRank(this js.Value, args []js.Value) interface{} {
	arrIn := args[0]

	var arr [7]int
	for i := 0; i < 7; i++ {
		arr[i] = arrIn.Index(i).Int()
	}

	rank := poker.GetRank(arr)
	result := rank

	return js.ValueOf(result)
}

func getRankFive(this js.Value, args []js.Value) interface{} {
	arrIn := args[0]

	var arr [5]int
	for i := 0; i < 7; i++ {
		arr[i] = arrIn.Index(i).Int()
	}

	rank := poker.GetRankFive(arr)
	result := rank

	return js.ValueOf(result)
}

func calcEquity(this js.Value, args []js.Value) interface{} {
	playerCardsIn := args[0]
	tableCardsIn := args[1]

	P := playerCardsIn.Length()
	var playerCards = make([][2]int, P)

	var i int

	for i = 0; i < P; i++ {
		onePlayerCardIn := playerCardsIn.Index(i)
		for j := 0; j < 2; j++ {
			playerCards[i][j] = onePlayerCardIn.Index(j).Int()
		}
	}

	T := tableCardsIn.Length()
	var tableCards = make([]int, T)
	for i = 0; i < T; i++ {
		tableCards[i] = tableCardsIn.Index(i).Int()
	}

	equity := poker.CalcEquity(playerCards, tableCards)

	result := make([]interface{}, 0)
	for _, e := range equity {
		var obj = map[string]interface{}{
			"win": e.Win,
			"tie": e.Tie,
		}
		result = append(result, obj)
	}
	return js.ValueOf(result)
}

func calcEquityMonteCarlo(this js.Value, args []js.Value) interface{} {
	playerCardsIn := args[0]
	tableCardsIn := args[1]
	nbPlayer := args[2].Int()
	nbGame := args[3].Int()

	var i int

	var playerCards [2]int
	for i = 0; i < 2; i++ {
		playerCards[i] = playerCardsIn.Index(i).Int()
	}

	T := tableCardsIn.Length()
	var tableCards = make([]int, T)
	for i = 0; i < T; i++ {
		tableCards[i] = tableCardsIn.Index(i).Int()
	}

	equity := poker.CalcEquityMonteCarlo(playerCards, tableCards, nbPlayer, nbGame)

	var result = map[string]interface{}{
		"win": equity.Win,
		"tie": equity.Tie,
	}
	return js.ValueOf(result)
}

func buildFiveHandStats(this js.Value, args []js.Value) interface{} {

	poker.BuildFiveHandStats(false)
	res := poker.FiveHandTypeStats

	result := make(map[string]interface{})
	for k, v := range res {
		var obj map[string]interface{}
		obj["nbHand"] = v.NbHand
		obj["minRank"] = v.MinRank
		obj["maxRank"] = v.MaxRank
		obj["nbOccur"] = v.NbOccur
		result[k] = obj
	}
	return js.ValueOf(result)
}

func buildSevenHandStats(this js.Value, args []js.Value) interface{} {

	poker.BuildSevenHandStats(false)
	res := poker.SevenHandTypeStats

	result := make(map[string]interface{})
	for k, v := range res {
		var obj map[string]interface{}
		obj["nbHand"] = v.NbHand
		obj["minRank"] = v.MinRank
		obj["maxRank"] = v.MaxRank
		obj["nbOccur"] = v.NbOccur
		result[k] = obj
	}
	return js.ValueOf(result)
}
