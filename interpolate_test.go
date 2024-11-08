package imagetk

import (
	"testing"
)

func TestLinearInterpolation(t *testing.T) {
	pixelData := [][]float32{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{1, 1, 1, 1},
		{1, 1, 1, 1},
	}
	img, err := GetImageFromArray(pixelData)
	if err != nil {
		t.Errorf("Error creating image from array: %v", err)
	}
	img.SetOrigin([]float64{0.5, 0.5})

	interpolator := LinearInterpolation{
		size:      []uint32{8, 8},
		spacing:   []float64{0.5, 0.5},
		origin:    []float64{0.25, 0.25},
		direction: [9]float64{1, 0, 0, 0, 1, 0, 0, 0, 1},
	}
	newImg, err := img.Resample(interpolator)
	if err != nil {
		t.Errorf("Error resampling image: %v", err)
	}

	size := newImg.GetSize()
	if size[0] != 8 || size[1] != 8 {
		t.Errorf("Expected size to be [8, 8], got %v", size)
	}
	pixelValue, err := newImg.GetPixelAsFloat32([]uint32{5, 5})
	if err != nil {
		t.Errorf("Error getting pixel value: %v", err)
	}

	expected := float32(1)
	if pixelValue != expected {
		t.Errorf("Expected pixel value to be %v, got %v", expected, pixelValue)
	}

}

func TestLinearInterpolation2(t *testing.T) {
	pixelData := [][]float32{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}
	img, err := GetImageFromArray(pixelData)
	if err != nil {
		t.Errorf("Error creating image from array: %v", err)
	}
	img.SetOrigin([]float64{0.5, 0.5})

	interpolator := LinearInterpolation{
		size:      []uint32{5, 5},
		spacing:   []float64{3.0 / 5.0, 3.0 / 5.0},
		origin:    []float64{3.0 / 5.0 / 2, 3.0 / 5.0 / 2},
		direction: [9]float64{1, 0, 0, 0, 1, 0, 0, 0, 1},
	}
	newImg, err := img.Resample(interpolator)
	if err != nil {
		t.Errorf("Error resampling image: %v", err)
	}

	size := newImg.GetSize()
	if size[0] != 5 || size[1] != 5 {
		t.Errorf("Expected size to be [5, 5], got %v", size)
	}

	pixelValue, err := newImg.GetPixelAsFloat32([]uint32{2, 2})
	if err != nil {
		t.Errorf("Error getting pixel value: %v", err)
	}

	expected := float32(1)
	if pixelValue != expected {
		t.Errorf("Expected pixel value to be %v, got %v", expected, pixelValue)
	}

}
