package gotest

import "testing"

func TestAdd(t *testing.T) {
	actual := Add(1, 1)

	if actual != 2 {
		t.Fatal("expected 1+1 to equal 2")
	}
}
