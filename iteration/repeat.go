package iteration

// Repeat : returns a new string with the given string repeated 5 times.
func Repeat(character string, times int) string {
	var rep string
	for i := 0; i < times; i++ {
		rep += character
	}
	return rep
}
