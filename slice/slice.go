package main

import "fmt"

// ["chennai", "blore", "delhi", "pune", "blore", "kolkata"]
// cityname -> frequency
// string -> int
func main() {
	cities := []string{"chennai", "blore", "delhi", "pune", "blore", "kolkata"}
	f := freq(cities)

	for i := 0; i < len(cities); i++ {
		if f[cities[i]] > 1 {
			city := cities[i]
			cities = removeIndex(cities, i)
			f[city]--
		}
	}
	fmt.Println(cities)
}

func freq(xs []string) map[string]int {
	m := make(map[string]int)
	for i := 0; i < len(xs); i++ {
		// m[xs[i]]++
		ei := m[xs[i]]
		fmt.Println("ei", ei)
		m[xs[i]] = ei + 1
	}

	return m
}

func removeIndex(s []string, v int) []string {
	//  var rs []int
	//  rs =
	// three dots are called ellipsis - variadic function
	return append(s[:v], s[(v+1):]...)
}
