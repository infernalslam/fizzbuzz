package cap_test

import (
	"testing"

	"github.com/infernalslam/gooooooo/cap"
)

func TestCap1111(t *testing.T) {
	result := cap.Cap(1, 1, 1, 1)
	expected := "1 + one"
	if expected != result {
		t.Errorf("it should say %q but get %q", expected, result)
	}
}

func TestCap2111(t *testing.T) {
	result := cap.Cap(2, 1, 1, 1)
	expected := "one + 1"
	if expected != result {
		t.Errorf("it should say %q but get %q", expected, result)
	}
}

func TestCap2244(t *testing.T) {
	result := cap.Cap(2, 2, 4, 4)
	expected := "two * 4"
	if expected != result {
		t.Errorf("it should say %q but get %q", expected, result)
	}
}

func TestCap2341(t *testing.T) {
	result := cap.Cap(2, 3, 4, 1)
	expected := "three * 1"
	if expected != result {
		t.Errorf("it should say %q but get %q", expected, result)
	}
}
