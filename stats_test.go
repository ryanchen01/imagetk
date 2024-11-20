package imagetk

import "testing"

func TestMin(t *testing.T) {
	tests := []struct {
		name      string
		pixelData [][]int8
		expect    int8
	}{
		{name: "min of 0", pixelData: [][]int8{{0, 0, 0}, {0, 0, 0}, {1, 1, 1}, {1, 1, 1}}, expect: 0},
		{name: "min of 1", pixelData: [][]int8{{1, 1, 1}, {1, 1, 1}, {10, 10, 10}, {10, 10, 10}}, expect: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			img, err := GetImageFromArray(tt.pixelData)
			if err != nil {
				t.Errorf("Error creating image from array: %v", err)
			}
			minValue := img.Min()
			if minValue.(int8) != tt.expect {
				t.Errorf("Expected min value to be %v, got %v", tt.expect, minValue)
			}
		})
	}
}

func TestMax(t *testing.T) {
	tests := []struct {
		name      string
		pixelData [][]int8
		expect    int8
	}{
		{name: "max of 10", pixelData: [][]int8{{1, 1, 1}, {1, 1, 1}, {10, 10, 10}, {10, 10, 10}}, expect: 10},
		{name: "max of 1", pixelData: [][]int8{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}, {1, 1, 1}}, expect: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			img, err := GetImageFromArray(tt.pixelData)
			if err != nil {
				t.Errorf("Error creating image from array: %v", err)
			}
			maxValue := img.Max()
			if maxValue.(int8) != tt.expect {
				t.Errorf("Expected max value to be %v, got %v", tt.expect, maxValue)
			}
		})
	}
}

func TestMean(t *testing.T) {
	tests := []struct {
		name      string
		pixelData [][]int8
		expect    int8
	}{
		{name: "mean of 0", pixelData: [][]int8{{0, 0, 0}, {0, 0, 0}, {1, 1, 1}, {1, 1, 1}}, expect: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			img, err := GetImageFromArray(tt.pixelData)
			if err != nil {
				t.Errorf("Error creating image from array: %v", err)
			}
			meanValue := img.Mean()
			if meanValue.(int8) != tt.expect {
				t.Errorf("Expected mean value to be %v, got %v", tt.expect, meanValue)
			}
		})
	}
}

func TestExactMean(t *testing.T) {
	tests := []struct {
		name      string
		pixelData [][]int8
		expect    float64
	}{
		{name: "exact mean of 0.5", pixelData: [][]int8{{0, 0, 0}, {0, 0, 0}, {1, 1, 1}, {1, 1, 1}}, expect: 0.5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			img, err := GetImageFromArray(tt.pixelData)
			if err != nil {
				t.Errorf("Error creating image from array: %v", err)
			}
			exactMeanValue := img.ExactMean()
			if exactMeanValue != tt.expect {
				t.Errorf("Expected exact mean value to be %v, got %v", tt.expect, exactMeanValue)
			}
		})
	}
}

func TestMedian(t *testing.T) {
	tests := []struct {
		name      string
		pixelData [][]int8
		expect    float64
	}{
		{name: "median of 0.5", pixelData: [][]int8{{0, 0, 0}, {0, 0, 0}, {1, 1, 1}, {1, 1, 1}}, expect: 0.5},
		{name: "median of 1", pixelData: [][]int8{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}, {1, 1, 1}}, expect: 1},
		{name: "median of 1", pixelData: [][]int8{{1, 1, 1}, {1, 1, 1}, {2, 2, 2}}, expect: 1.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			img, err := GetImageFromArray(tt.pixelData)
			if err != nil {
				t.Errorf("Error creating image from array: %v", err)
			}
			medianValue := img.Median()
			if medianValue != tt.expect {
				t.Errorf("Expected median value to be %v, got %v", tt.expect, medianValue)
			}
		})
	}
}

func TestStd(t *testing.T) {
	tests := []struct {
		name      string
		pixelData [][]int8
		expect    float64
	}{
		{name: "std of 0", pixelData: [][]int8{{0, 0, 0}, {0, 0, 0}, {1, 1, 1}, {1, 1, 1}}, expect: 0.5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			img, err := GetImageFromArray(tt.pixelData)
			if err != nil {
				t.Errorf("Error creating image from array: %v", err)
			}
			stdValue := img.Std()
			if stdValue.(float64) != tt.expect {
				t.Errorf("Expected std value to be %v, got %v", tt.expect, stdValue)
			}
		})
	}
}
