package imagetk

import (
	"fmt"
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
	interpolator := LinearInterpolation{
		size:      []uint32{8, 8},
		spacing:   []float64{0.5, 0.5},
		origin:    []float64{0, 0},
		direction: [9]float64{1, 0, 0, 0, 1, 0, 0, 0, 1},
	}
	newImg, err := img.Resample(interpolator)
	if err != nil {
		t.Errorf("Error resampling image: %v", err)
	}

	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			fmt.Print(newImg.pixels.([]float32)[i*8+j], " ")
		}
		fmt.Println()
	}

	size := newImg.GetSize()
	if size[0] != 8 || size[1] != 8 {
		t.Errorf("Expected size to be [8, 8], got %v", size)
	}
	pixelValue, err := newImg.GetPixelAsFloat32([]uint32{5, 5})
	if err != nil {
		t.Errorf("Error getting pixel value: %v", err)
	}
	if pixelValue != 1 {
		t.Errorf("Expected pixel value to be 1, got %v", pixelValue)
	}

}
