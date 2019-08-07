package msp

import (
	"testing"
)

var testcases = [...]struct {
	n       []int
	emaxSum int
	emi     int
	emj     int
}{
	{[]int{1, -2, 2, 3, -4, -5}, 5, 2, 3},
	{[]int{1, 2, 3, 4}, 10, 0, 3},
	{[]int{-1, -2, -3}, -1, 0, 0},
	{[]int{-3, -2, -1}, -1, 2, 2},
	{[]int{0}, 0, 0, 0},
	{[]int{-1}, -1, 0, 0},
	{[]int{}, -1, -1, -1},
	{[]int{1, 2, 3, 4}, 10, 0, 3},
	{[]int{}, 10000, 9999, 9999},
}

const largeN = 2000

func populateLargeData() {
	if len(testcases[8].n) == 0 {
		for i := 1; i < largeN; i++ {
			testcases[8].n = append(testcases[8].n, i*-1)
		}
		testcases[8].n = append(testcases[8].n, largeN)
	}
	testcases[8].emaxSum = largeN
	testcases[8].emi = largeN - 1
	testcases[8].emj = largeN - 1
}

func TestMspN1(t *testing.T) {
	populateLargeData()

	for _, tt := range testcases {
		maxSum, mi, mj := mspN1(tt.n)
		if maxSum != tt.emaxSum || mi != tt.emi || mj != tt.emj {
			t.Errorf("expected %v %v %v, got %v, %v, %v", tt.emaxSum, tt.emi, tt.emj, maxSum, mi, mj)
		}
	}
}

func TestMspN2(t *testing.T) {
	populateLargeData()

	for _, tt := range testcases {
		maxSum, mi, mj := mspN2(tt.n)
		if maxSum != tt.emaxSum || mi != tt.emi || mj != tt.emj {
			t.Errorf("expected %v %v %v, got %v, %v, %v", tt.emaxSum, tt.emi, tt.emj, maxSum, mi, mj)
		}
	}
}

func TestMspN3(t *testing.T) {
	populateLargeData()

	for _, tt := range testcases {
		maxSum, mi, mj := mspN3(tt.n)
		if maxSum != tt.emaxSum || mi != tt.emi || mj != tt.emj {
			t.Errorf("expected %v %v %v, got %v, %v, %v", tt.emaxSum, tt.emi, tt.emj, maxSum, mi, mj)
		}
	}
}
