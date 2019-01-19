package unittest

import (
	"testing"
)

func TestAdd(t *testing.T) {
	var a, b int = 10, 20
	t.Logf("Given two numbers %d and %d.", a, b)
	t.Logf("When calling Add function with the two numbers.")
	c := Add(a, b)
	if c != a+b {
		t.Fail()
	}
	t.Logf("Then result should be %d.", a+b)
}
