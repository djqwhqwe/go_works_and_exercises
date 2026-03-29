package task11111

import "testing"

type sumTestData struct {
	inp [8]int
	out int
}

type averageTestData struct {
	inp [8]int
	out float64
}

type averageFloatTestData struct {
	inp [8]float64
	out float64
}

type reverseTestData struct {
	inp [8]int
	out [8]int
}

func TestSum(t *testing.T) {
	testCases := []sumTestData{
		{inp: [8]int{1, 2, 3, 4, 5, 6, 7, 8}, out: 36},
		{inp: [8]int{3, 2, 5, 1, 234, 12, 4, 34}, out: 295},
		{inp: [8]int{-123, 23, -1, 23, 34, 87, 3, 1}, out: 47},
	}
	for _, tc := range testCases {
		r := sum(tc.inp)
		if r != tc.out {
			t.Errorf("Unexpected result. Input: %v, Expected: %v, Got: %v", tc.inp, tc.out, r)
		}
	}
}

func TestAverage(t *testing.T) {
	testCases := []averageTestData{
		{inp: [8]int{1, 2, 3, 4, 5, 6, 7, 8}, out: 4.5},
		{inp: [8]int{3, 2, 5, 1, 234, 12, 4, 34}, out: 36.875},
		{inp: [8]int{-123, 23, -1, 23, 34, 87, 3, 1}, out: 5.875},
	}
	for _, tc := range testCases {
		r := average(tc.inp)
		if r != tc.out {
			t.Errorf("Unexpected result. Input: %v, Expected: %v, Got: %v", tc.inp, tc.out, r)
		}
	}
}

func TestAverageFloat(t *testing.T) {
	testCases := []averageFloatTestData{
		{inp: [8]float64{-83.2, 84.0, 19.5, 47.8, -55.1, 59.1, 64.9, -81.7}, out: 6.9125},
		{inp: [8]float64{81.1, 78.1, -62.0, -31.8, 68.0, -74.2, 62.7, 86.2}, out: 26.012499999999996},
		{inp: [8]float64{76.1, -62.5, -40.6, -3.1, -29.9, 70.7, -26.9, -9.9}, out: -3.2625},
	}
	for _, tc := range testCases {
		r := averageFloat(tc.inp)
		if r != tc.out {
			t.Errorf("Unexpected result. Input: %v, Expected: %v, Got: %v", tc.inp, tc.out, r)
		}
	}
}

func TestReverse(t *testing.T) {
	testCases := []reverseTestData{
		{inp: [8]int{-63, 68, -49, 60, -7, -38, -15, -10}, out: [8]int{-10, -15, -38, -7, 60, -49, 68, -63}},
		{inp: [8]int{-45, -8, 45, 84, 71, 6, -25, 24}, out: [8]int{24, -25, 6, 71, 84, 45, -8, -45}},
		{inp: [8]int{48, -35, 62, 47, -7, -12, 92, 93}, out: [8]int{93, 92, -12, -7, 47, 62, -35, 48}},
	}
	for _, tc := range testCases {
		r := reverse(tc.inp)
		if r != tc.out {
			t.Errorf("Unexpected result. Input: %v, Expected: %v, Got: %v", tc.inp, tc.out, r)
		}
	}
}
