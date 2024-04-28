package main

import (
	"fmt"
	"strconv"
)

func getMultiplicationTable(n int) [][]int {
	multiplicationTable := make([][]int, n)
	for y := 1; y <= n; y++ {
		row := make([]int, n)
		for x := 1; x <= n; x++ {
			row[x-1] = x * y
		}
		multiplicationTable[y-1] = row
	}
	return multiplicationTable
}

func printMultiplicationTable(n int) {
	multiplicationTable := getMultiplicationTable(n)
	maxNumLen := len(strconv.Itoa(n * n))
	maxMultiplierLen := len(strconv.Itoa(n))
	fmt.Printf("%*s", maxMultiplierLen, "")
	for i := 1; i <= n; i++ {
		fmt.Printf("%*d", maxNumLen+1, i)
	}
	fmt.Println()
	for i, row := range multiplicationTable {
		fmt.Printf("%*d", maxMultiplierLen, i+1)
		for _, value := range row {
			fmt.Printf("%*d", maxNumLen+1, value)
		}
		fmt.Println()
	}
}
func main() {
	var n int
	fmt.Scan(&n)
	printMultiplicationTable(n)
}
