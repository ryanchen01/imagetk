package imagetk

import (
	"fmt"
	"testing"
)

func TestLinearInterpolation(t *testing.T) {
	pixelData := [][]float32{
		{0, 0, 0},
		{0, 0, 0},
		{1, 1, 1},
	}
	img, err := GetImageFromArray(pixelData)
	if err != nil {
		t.Errorf("Error creating image from array: %v", err)
	}
	interpolator := LinearInterpolation{
		size:      []uint32{6, 6},
		spacing:   []float64{0.5, 0.5},
		origin:    []float64{0, 0},
		direction: [9]float64{1, 0, 0, 0, 1, 0, 0, 0, 1},
	}
	newImg, err := img.Resample(interpolator)
	if err != nil {
		t.Errorf("Error resampling image: %v", err)
	}
	fmt.Println(newImg.pixels.([]float32))
	size := newImg.GetSize()
	if size[0] != 6 || size[1] != 6 {
		t.Errorf("Expected size to be [6, 6], got %v", size)
	}
	pixelValue, err := newImg.GetPixelAsFloat32([]uint32{3, 3})
	if err != nil {
		t.Errorf("Error getting pixel value: %v", err)
	}
	if pixelValue != 0.5 {
		t.Errorf("Expected pixel value to be 0.5, got %v", pixelValue)
	}

}
