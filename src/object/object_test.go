package object

import "testing"

func TestStringHashKey(t *testing.T) {
	h1 := &String{Value: "H"}
	h2 := &String{Value: "H"}
	d1 := &String{Value: "X"}
	d2 := &String{Value: "X"}

	if h1.HashKey() != h2.HashKey() {
		t.Errorf("Strings with same content have different hash keys")
	}

	if d1.HashKey() != d2.HashKey() {
		t.Errorf("Stings with same content have different hash keys")
	}

	if h1.HashKey() == d1.HashKey() {
		t.Errorf("Stings with different content have same hash keys")
	}
}
