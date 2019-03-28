package main

import "testing"

func TestFizzBuzzEqualsFizzWhenMultipleOfThree(t *testing.T) {
	result := fizzBuzz(3)
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
