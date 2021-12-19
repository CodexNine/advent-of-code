package main

import (
	"os"
	"fmt"
	"encoding/csv"
	"log"
	"strconv"
) 

func main() {
	openfile, err := os.Open("./data/motion.csv")
	if err != nil {
        log.Println(err)
        return
    }

	filedata, err := csv.NewReader(openfile).ReadAll()
    if err != nil {
        log.Println(err)
        return
    }
	
	slicedData := filedata[0]

	var x_pos int = 0
	var y_pos int = 0

	var aim int = 0


	for index, direction := range slicedData {
		switch direction {
		case "forward":
			val, _ := strconv.Atoi(slicedData[index+1])
			x_pos += val
			y_pos += (aim * val)
		case "down":
			val, _ := strconv.Atoi(slicedData[index+1])
			// y_pos += val
			aim += val
		case "up":
			val, _ := strconv.Atoi(slicedData[index+1])
			// y_pos -= val
			aim -= val
		}
	}

	fmt.Printf("horizontal displacement: %d\n", x_pos)
	fmt.Printf("depth displacement: %d\n", y_pos)
	fmt.Printf("Answer: %d\n", y_pos*x_pos)
}