package fizzbuzz_test

import (
	"testing"

	fizzbuzz "github.com/infernalslam/gooooooo"
)

func TestfizzbuzzShouldSayOne(t *testing.T) {
	result := fizzbuzz.Say(1)
	expected := "1"
	if result != expected {
		t.Errorf("it should say %q but get %q", expected, result)
	}
}

func TestfizzbuzzShouldSayTwo(t *testing.T) {
	result := fizzbuzz.Say(2)
	expected := "2"
	if result != expected {
		t.Errorf("it should say %q but get %q", expected, result)
	}
}

func TestfizzbuzzShouldSayFizz(t *testing.T) {
	result := fizzbuzz.Say(3)
	expected := "Fizz"
	if result != expected {
		t.Errorf("it should say %q but get %q", expected, result)
	}
}
