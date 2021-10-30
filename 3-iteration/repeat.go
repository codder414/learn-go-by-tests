package iteration

func Repeat(item string) string {
	var result string
	for i := 0; i < 5; i++ {
		result += item
	}
	return result
}
