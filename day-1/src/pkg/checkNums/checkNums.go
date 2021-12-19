package checkNums

func ReturnNum(data []int) int {
	count := 0
	for index, number := range data {
		if(index < 1) {
			continue
		} else if(number > data[index - 1]) {
			count++
		}
	}
	return count
}

func GroupNum(data []int) int {
	count := 0
	prevValue := 0
	for index,_ := range data {
		if(index < 2) {
			continue
		} else {
			groupValue := (data[index-2] + data[index-1] + data[index])	
			if(groupValue > prevValue && prevValue > 0) {
				count++
			}
			prevValue = groupValue
		}
	}
	return count
}