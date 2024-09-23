package iteration

func Repeat(letter string, repeatCount int) string {
	var result string
	for i := 0; i < repeatCount; i++ {
		result = result + letter
	}
	return result
}
