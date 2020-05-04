package iteration

// Repeat : returns a new string with the given string repeated 5 times.
func Repeat(character string) string {
	var rep string
	for i := 0; i < 5; i++ {
		rep = rep + character
	}
	return rep
}
