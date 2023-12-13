package util

import "testing"

func TestReverseInt(t *testing.T) {
	_, err := ReverseInt(123)
	if err != nil {
		t.Fatal("some err with num:", 123)
	}
	_, err = ReverseInt(-649)
	if err != nil {
		t.Fatal("some err with num:", -649)
	}
	_, err = ReverseInt(0)
	if err != nil {
		t.Fatal("some err with num:", 0)
	}
}
