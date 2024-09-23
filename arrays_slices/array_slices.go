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
