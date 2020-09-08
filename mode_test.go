package main

import (
	"testing"
)

type testMode struct {
	list []float64
	mode float64
}

func TestMode(t *testing.T) {
	tests := []testMode{
		testMode{list: []float64{50, 12, 13, 12}, mode: 12},
		testMode{list: []float64{50, 52, 51, 53, 50}, mode: 50},
		testMode{list: []float64{10, 20, 40, 25, 30, 31}, mode: 0},
	}
	for _, test := range tests {
		got := mode(test.list)
		want := test.mode
		if want != got {
			t.Error("\ndata:", test.list, "\nwant:", want, "\ngot:", got)
		}
	}
}
