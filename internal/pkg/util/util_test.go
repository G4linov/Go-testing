package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

/*
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
*/

func TestReverseInt(t *testing.T) {
	req := require.New(t)
	_, err := ReverseInt(123)
	req.NoError(err)
	_, err = ReverseInt(-649)
	req.NoError(err)
	_, err = ReverseInt(0)
	req.NoError(err)
}
