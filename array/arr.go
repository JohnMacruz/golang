package main

import "fmt"

func main() {
	var arr = [5]int{10, 20, 30, 40, 50}
	for i := 0; i < 5; i++ {
		fmt.Printf("%d\n", arr[i])
	}
}
