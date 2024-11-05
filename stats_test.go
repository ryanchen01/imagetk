package imagetk

import "testing"

func TestMinMaxMean(t *testing.T) {
	pixelData := [][]int8{
		{0, 0, 0},
		{0, 0, 0},
		{1, 1, 1},
		{1, 1, 1},
	}
	img, err := GetImageFromArray(pixelData)
	if err != nil {
		t.Errorf("Error creating image from array: %v", err)
	}

	minValue := img.Min()
	if minValue.(int8) != int8(0) {
		t.Errorf("Expected min value to be 0, got %v", minValue)
	}

	maxValue := img.Max()
	if maxValue.(int8) != int8(1) {
		t.Errorf("Expected max value to be 1, got %v", maxValue)
	}

	meanValue := img.Mean()
	if meanValue.(int8) != int8(0) {
		t.Errorf("Expected mean value to be 0, got %v", meanValue)
	}

	exactMeanValue := img.ExactMean()
	if exactMeanValue.(float64) != float64(0.5) {
		t.Errorf("Expected exact mean value to be 0.5, got %v", exactMeanValue)
	}
}
