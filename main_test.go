package main

import (
	"testing"
)

type testData struct {
	list   []float64
	mean   float64
	median float64
	mode   float64
}

func TestMean(t *testing.T) {
	tests := []testData{
		testData{list: []float64{50, 12, 13, 12}, mean: 21.75, median: 12.5, mode: 12},
		testData{list: []float64{50, 52, 51, 53, 50}, mean: 51.2, median: 51.5, mode: 50},
		testData{list: []float64{10, 20, 40, 25, 30, 31}, mean: 21.75, median: 12.5, mode: 0},
	}
	for _, test := range tests {
		got := mean(test.list)
		want := test.mean
		if want != got {
			t.Error("\ndata:", test.list, "\nwant:", want, "\ngot:", got)
		}
	}
}

func TestMedian(t *testing.T) {
	tests := []testData{
		testData{list: []float64{50, 12, 13, 12}, mean: 21.75, median: 12.5, mode: 12},
		testData{list: []float64{50, 52, 51, 53, 50}, mean: 51.2, median: 51.5, mode: 50},
		testData{list: []float64{10, 20, 40, 25, 30, 31}, mean: 26, median: 27.5, mode: 0},
	}
	for _, test := range tests {
		got := median(test.list)
		want := test.median
		if want != got {
			t.Error("\ndata:", test.list, "\nwant:", want, "\ngot:", got)
		}
	}
}

func TestMode(t *testing.T) {
	tests := []testData{
		testData{list: []float64{50, 12, 13, 12}, mean: 21.75, median: 12.5, mode: 12},
		testData{list: []float64{50, 52, 51, 53, 50}, mean: 51.2, median: 51.5, mode: 50},
		testData{list: []float64{10, 20, 40, 25, 30, 31}, mean: 21.75, median: 12.5, mode: 0},
	}
	for i, test := range tests {
		got := mode(test.list)
		want := test.mode
		if want != got {
			t.Error(i,"\ndata:", test.list, "\nwant:", want, "\ngot:", got)
		}
	}
}
