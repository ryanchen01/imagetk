package imagetk

import (
	"math"
	"testing"
)

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

func TestPercentile(t *testing.T) {
	tests := []struct {
		name      string
		pixelData [][]int8
		percent   float64
		expect    float64
	}{
		{name: "percentile of 0.5", pixelData: [][]int8{{0, 0, 0}, {0, 0, 0}, {1, 1, 1}, {1, 1, 1}}, percent: 0.5, expect: 0.5},
		{name: "percentile of 0.99", pixelData: [][]int8{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, percent: 0.99, expect: 8.92},
		{name: "percentile of 0.1", pixelData: [][]int8{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, percent: 0.1, expect: 1.8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			img, err := GetImageFromArray(tt.pixelData)
			if err != nil {
				t.Errorf("Error creating image from array: %v", err)
			}
			percentileValue := img.Percentile(tt.percent)
			if percentileValue != tt.expect {
				t.Errorf("Expected percentile value to be %v, got %v", tt.expect, percentileValue)
			}
		})
	}
}

func TestProduct(t *testing.T) {
	tests := []struct {
		name      string
		pixelData [][]int8
		expect    int64
	}{
		{name: "product of 1", pixelData: [][]int8{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}, {1, 1, 1}}, expect: 1},
		{name: "product of 1000000000", pixelData: [][]int8{{10, 10, 10}, {10, 10, 10}, {10, 10, 10}}, expect: 1000000000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			img, err := GetImageFromArray(tt.pixelData)
			if err != nil {
				t.Errorf("Error creating image from array: %v", err)
			}
			productValue := img.Product()
			if productValue.(int64) != tt.expect {
				t.Errorf("Expected product value to be %v, got %v", tt.expect, productValue)
			}
		})
	}
}

func TestOtsuThreshold(t *testing.T) {
	tests := []struct {
		name      string
		pixelData [][]float64
		expect    float64
	}{
		{name: "otsu threshold of 0", pixelData: [][]float64{{0, 0, 0}, {0, 0, 0}, {1, 1, 1}, {1, 1, 1}}, expect: 0},
		{name: "otsu threshold of 1", pixelData: [][]float64{{0, 0, 0}, {1, 1, 1}, {2, 2, 2}, {3, 3, 3}}, expect: 1},
		{name: "otsu threshold of 5.97", pixelData: [][]float64{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {10, 11, 12}}, expect: 5.976},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			img, err := GetImageFromArray(tt.pixelData)
			if err != nil {
				t.Errorf("Error creating image from array: %v", err)
			}
			otsuThresholdValue := img.OtsuThreshold()
			if !almostEqual(otsuThresholdValue, tt.expect, 1e-3) {
				t.Errorf("Expected otsu threshold value to be %v, got %v", tt.expect, otsuThresholdValue)
			}
		})
	}
}

func almostEqual(a, b float64, tolerance float64) bool {
	diff := math.Abs(a - b)
	return diff <= tolerance
}
