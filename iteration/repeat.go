package iteration

// Repeat the characater given in the parameter
func Repeat(character string, repeatCount int) (repeated string) {
	for i := 0; i < repeatCount; i++ {
		repeated += character
	}

	return repeated
}
