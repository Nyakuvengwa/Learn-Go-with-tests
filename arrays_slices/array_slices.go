package arraysslices

func Sum(numbers []int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersArray ...[]int) (result []int) {
	for _, numbers := range numbersArray {
		result = append(result, Sum(numbers))
	}
	return
}

func SumAllTails(numbers ...[]int) (result []int) {
	for _, numbersArray := range numbers {
		if len(numbersArray) == 0 {
			result = append(result, 0)
		} else {
			tail := numbersArray[1:]
			result = append(result, Sum(tail))
		}
	}
	return
}
