package main

import (
	"fmt"
	"github.com/CodexNine/advent-of-code/day-1/src/pkg/readFile"
)

func main() {
	sonarScan := readFile.ReadInts("sonar-data.txt")
	fmt.Println(sonarScan)
}