package main

import "testing"

func TestFizzBuzzEqualsFizzWhenMultipleOfThree(t *testing.T) {
	// got := fizzBuzz(3)
	// if got != "Fizz" {
	// 	t.Errorf("fizzBuzz(3) returns %d; want Fizz", got)
	// }
	result := fizzBuzz(3)
	expect := "Fizz"

	if result != expect {
		t.Fail()
	}
}
