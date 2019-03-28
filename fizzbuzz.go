package main

import "fmt"

func fizzBuzz(number int) string {
	fizz := (number%3 == 0)
	buzz := (number%5 == 0)

	if fizz && buzz {
		return "FizzBuzz"
	}
	if fizz {
		return "Fizz"
	}
	if buzz {
		return "Buzz"
	}
	return fmt.Sprintf("%d", number)
}


func main() {
}
