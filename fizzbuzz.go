package main

func fizzBuzz(number int) string {
	if number == 15 {
		return "FizzBuzz"
	}
	if number%3 == 0 {
		return "Fizz"
	}
	if number%5 == 0 {
		return "Buzz"
	}
	return ""
}

func main() {
}
