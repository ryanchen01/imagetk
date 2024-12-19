package imagetk

import (
	"testing"
)

func TestLinearInterpolator(t *testing.T) {
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

	interpolator := LinearInterpolator{
		Size:      []uint32{8, 8},
		Spacing:   []float64{0.5, 0.5},
		Origin:    []float64{0.25, 0.25},
		Direction: [9]float64{1, 0, 0, 0, 1, 0, 0, 0, 1},
		FillType:  FillTypeNearest,
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

func TestLinearInterpolator2(t *testing.T) {
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

	interpolator := LinearInterpolator{
		Size:      []uint32{5, 5},
		Spacing:   []float64{3.0 / 5.0, 3.0 / 5.0},
		Origin:    []float64{3.0 / 5.0 / 2, 3.0 / 5.0 / 2},
		Direction: [9]float64{1, 0, 0, 0, 1, 0, 0, 0, 1},
		FillType:  FillTypeNearest,
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

func TestLinearInterpolator3(t *testing.T) {
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

	interpolator := LinearInterpolator{
		Size:      []uint32{4, 4},
		Spacing:   []float64{1, 1},
		Origin:    []float64{0.5, 3.5},
		Direction: [9]float64{1, 0, 0, 0, -1, 0, 0, 0, 1},
		FillType:  FillTypeNearest,
	}
	newImg, err := img.Resample(interpolator)
	if err != nil {
		t.Errorf("Error resampling image: %v", err)
	}

	size := newImg.GetSize()
	if size[0] != 4 || size[1] != 4 {
		t.Errorf("Expected size to be [4, 4], got %v", size)
	}

	expected := [][]float32{
		{1, 1, 1, 1},
		{1, 1, 1, 1},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}

	for y := 0; y < 4; y++ {
		for x := 0; x < 4; x++ {
			pixelValue, err := newImg.GetPixelAsFloat32([]uint32{uint32(x), uint32(y)})
			if err != nil {
				t.Errorf("Error getting pixel value: %v", err)
			}

			if pixelValue != expected[y][x] {
				t.Errorf("Expected pixel value to be %v, got %v", expected[y][x], pixelValue)
			}
		}
	}

}

func TestNearestInterpolator(t *testing.T) {
	pixelData := [][]float32{
		{0, 0, 0, 0},
		{0, 1, 1, 0},
		{0, 1, 1, 0},
		{0, 0, 0, 0},
	}
	img, err := GetImageFromArray(pixelData)
	if err != nil {
		t.Errorf("Error creating image from array: %v", err)
	}
	img.SetOrigin([]float64{0.5, 0.5})

	interpolator := NearestInterpolator{
		Size:      []uint32{8, 8},
		Spacing:   []float64{0.5, 0.5},
		Origin:    []float64{0.25, 0.25},
		Direction: [9]float64{1, 0, 0, 0, 1, 0, 0, 0, 1},
	}
	newImg, err := img.Resample(interpolator)
	if err != nil {
		t.Errorf("Error resampling image: %v", err)
	}

	size := newImg.GetSize()
	if size[0] != 8 || size[1] != 8 {
		t.Errorf("Expected size to be [8, 8], got %v", size)
	}

	// Test center pixel value
	pixelValue, err := newImg.GetPixelAsFloat32([]uint32{3, 3})
	if err != nil {
		t.Errorf("Error getting pixel value: %v", err)
	}

	expected := float32(1)
	if pixelValue != expected {
		t.Errorf("Expected pixel value to be %v, got %v", expected, pixelValue)
	}

	// Test corner pixel value (should be 0)
	cornerValue, err := newImg.GetPixelAsFloat32([]uint32{0, 0})
	if err != nil {
		t.Errorf("Error getting corner pixel value: %v", err)
	}

	expectedCorner := float32(0)
	if cornerValue != expectedCorner {
		t.Errorf("Expected corner pixel value to be %v, got %v", expectedCorner, cornerValue)
	}
}

func TestLinearInterpolatorBoundaryPixels(t *testing.T) {
	pixelData := [][]float32{
		{1, 1},
		{1, 1},
	}
	img, err := GetImageFromArray(pixelData)
	if err != nil {
		t.Fatalf("Failed to create image: %v", err)
	}

	interpolator := LinearInterpolator{
		Size:      []uint32{4, 4},
		Spacing:   []float64{1, 1},
		Origin:    []float64{0, 0},
		Direction: [9]float64{1, 0, 0, 0, 1, 0, 0, 0, 1},
		FillType:  FillTypeZero,
	}
	newImg, err := img.Resample(interpolator)
	if err != nil {
		t.Fatalf("Resample failed: %v", err)
	}

	// Test out-of-bounds pixel
	pixelValue, err := newImg.GetPixelAsFloat32([]uint32{3, 3})
	if err != nil {
		t.Fatalf("Failed to get pixel: %v", err)
	}
	if pixelValue != 0 {
		t.Errorf("Expected 0 for out-of-bounds pixel, got %v", pixelValue)
	}
}

func TestLinearInterpolatorDirectionMatrix(t *testing.T) {
	pixelData := [][]float32{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	img, err := GetImageFromArray(pixelData)
	if err != nil {
		t.Fatalf("Failed to create image: %v", err)
	}

	interpolator := LinearInterpolator{
		Size:      []uint32{3, 3},
		Spacing:   []float64{1, 1},
		Origin:    []float64{2, 0},
		Direction: [9]float64{0, -1, 0, 1, 0, 0, 0, 0, 1}, // 90-degree rotation
		FillType:  FillTypeNearest,
	}

	newImg, err := img.Resample(interpolator)
	if err != nil {
		t.Fatalf("Resample failed: %v", err)
	}

	expected := [][]float32{
		{3, 6, 9},
		{2, 5, 8},
		{1, 4, 7},
	}

	for y := 0; y < 3; y++ {
		for x := 0; x < 3; x++ {
			pixelValue, err := newImg.GetPixelAsFloat32([]uint32{uint32(x), uint32(y)})
			if err != nil {
				t.Fatalf("Failed to get pixel: %v", err)
			}
			if pixelValue != expected[y][x] {
				t.Errorf("Expected %v, got %v", expected[y][x], pixelValue)
			}
		}
	}
}
