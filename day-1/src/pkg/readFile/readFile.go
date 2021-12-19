package readFile

import (
    "fmt"
    "io"
    "os"
)

func Read(filename string) (nums []int) {
	file, err := os.Open(filename)

	if(err != nil) {
		fmt.Println(err)
		os.Exit(1)
	}
	
	var perline int

	for {

			_, err := fmt.Fscanf(file, "%d\n", &perline) // give a patter to scan

			if err != nil {

					if err == io.EOF {
							break // stop reading the file
					}
					fmt.Println(err)
					os.Exit(1)
			}

			nums = append(nums, perline)
	}
	return nums
}
