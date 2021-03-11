package main

import "fmt"

func main() {
	statePopulations := map[string]int{ //MAP EXAMPLE
		"California":   39250017,
		"Texas":        27862596,
		"Florida":      20612439,
		"New York":     19745289,
		"Pennsylvania": 12802503,
		"Illinois":     12801539,
		"Ohio":         11614373,
	}
	m := map[[3]int]string{}
	fmt.Println(statePopulations, m)
	//can use m["k1"]
}
