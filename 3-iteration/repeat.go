package iteration

func Repeat(item string, times int) string {
	var result string
	for i := 0; i < times; i++ {
		result += item
	}
	return result
}
