package main

import (
	"reflect"
	"testing"
)

func TestEqualPlayers(t *testing.T) {
	expected := Player{
		name: "Mat",
		hp:   100,
	}
	have := Player{
		name: "Mat",
		hp:   100,
	}
	if !reflect.DeepEqual(expected, have) {
		t.Errorf("expected %+v and have %+v", expected, have)
	}
}

func TestCalculateValues(t *testing.T) {

	var (
		expected = 10
		a        = 5
		b        = 5
	)

	have := calculateValues(a, b)
	if expected != have {
		t.Errorf("excepted %d and have %d", expected, have)
	}

}

// to test all test cases  -> go test ./...
// to test a single file   ->  go test ./ -v -run TestCalculateValues
// to test a single file and remove cache   ->  go test ./ -v -run TestCalculateValues -count=1 -> count number doesn't matter you can put any number
