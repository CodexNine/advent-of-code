package main

import (
	"fmt"
	"readFile"
	"checkNums"
)

func main() {
	fileName := "./data/sonar-data.txt"
	sonarData := readFile.Read(fileName)
	
	fmt.Printf("%d", checkNums.GroupNum(sonarData))

}