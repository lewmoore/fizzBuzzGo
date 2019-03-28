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

func TestFizzBuzzEqualsBuzzWhenisFive(t *testing.T) {
	result := fizzBuzz(5)
	expect := "Buzz"

	if result != expect {
		t.Fail()
	}
}
func TestFizzBuzzEqualsBuzzWhenIsOfTen(t *testing.T) {
	result := fizzBuzz(10)
	expect := "Buzz"

	if result != expect {
		t.Fail()
	}
}

func TestFizzBuzzEqualsFizzBuzzWhenIsFifteen(t *testing.T) {
	result := fizzBuzz(15)
	expect := "FizzBuzz"

	if result != expect {
		t.Fail()
	}
}
func TestFizzBuzzEqualsFizzBuzzWhenIsThirty(t *testing.T) {
	result := fizzBuzz(30)
	expect := "FizzBuzz"

	if result != expect {
		t.Fail()
	}
}
