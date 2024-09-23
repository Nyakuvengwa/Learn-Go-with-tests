package iteration

func Repeat(letter string) string {
	result := ""
	for i := 0; i < 5; i++ {
		result = result + letter
	}
	return result
}
