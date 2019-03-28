package main

import "testing"

func TestFizzBuzzEqualsFizzWhenIsThree(t *testing.T) {
	result := fizzBuzz(3)
	expect := "Fizz"

	if result != expect {
		t.Fail()
	}
}
func TestFizzBuzzEqualsFizzWhenIsNine(t *testing.T) {
	result := fizzBuzz(9)
	expect := "Fizz"

	if result != expect {
		t.Fail()
	}
}

func TestFizzBuzzEqualsBuzzWhenMultipleOfFive(t *testing.T) {
	result := fizzBuzz(5)
	expect := "Buzz"

	if result != expect {
		t.Fail()
	}
}

func TestFizzBuzzEqualsFizzBuzzWhenMultipleOfFiveAndThree(t *testing.T) {
	result := fizzBuzz(15)
	expect := "FizzBuzz"

	if result != expect {
		t.Fail()
	}
}


