package fizzbuzz_test

import (
	"testing"

	fizzbuzz "github.com/infernalslam/gooooooo"
)

func TestFizzbuzzShouldSayOne(t *testing.T) {
	result := fizzbuzz.Say(1)
	expected := "1"
	if expected != result {
		t.Errorf("it should say %q but get %q", expected, result)
	}
}

func TestFizzbuzzShouldSayTwo(t *testing.T) {
	result := fizzbuzz.Say(2)
	expected := "2"
	if expected != result {
		t.Errorf("it should say %q but get %q", expected, result)
	}
}

func TestFizzbuzzShouldSayFizz(t *testing.T) {
	result := fizzbuzz.Say(3)
	expected := "Fizz"
	if expected != result {
		t.Errorf("it should say %q but get %q", expected, result)
	}
}

func TestFizzbuzzShouldSayBuzz(t *testing.T) {
	result := fizzbuzz.Say(5)
	expected := "Buzz"
	if expected != result {
		t.Errorf("it should say %q but get %q", expected, result)
	}
}

func TestFizzbuzzShouldSayFizzBuzz(t *testing.T) {
	result := fizzbuzz.Say(15)
	expected := "FizzBuzz"
	if expected != result {
		t.Errorf("it should say %q but get %q", expected, result)
	}
}

func TestFizzbuzzShouldSayFizzBuzzFive(t *testing.T) {
	result := fizzbuzz.Say(5)
	expected := "Buzz"
	if expected != result {
		t.Errorf("it should say %q but get %q", expected, result)
	}
}

func TestFizzbuzzShouldSayFizzBuzzNine(t *testing.T) {
	result := fizzbuzz.Say(9)
	expected := "Fizz"
	if expected != result {
		t.Errorf("it should say %q but get %q", expected, result)
	}
}
