package integers

// Add takes two integers and returns the sum of them.
func Add(x, y int) int {
	return x + y
}

// AddMulti adds all the integers together
func AddMulti(numbersToSum ...int) (sum int) {
	for _, number := range numbersToSum {
		sum += number
	}

	return
}
