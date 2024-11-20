package imagetk

import (
	"testing"
)

func TestBinaryDilate2D(t *testing.T) {
	// Create a 5x5 test image with a single point in the middle
	data := [][]int8{
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
	}
	img, err := GetImageFromArray(data)
	if err != nil {
		t.Fatal(err)
	}

	// Test with kernel size 3
	dilated, err := BinaryDilate(img, 3)
	if err != nil {
		t.Fatal(err)
	}

	// Expected result with 3x3 kernel
	expected := [][]int8{
		{0, 0, 0, 0, 0},
		{0, 1, 1, 1, 0},
		{0, 1, 1, 1, 0},
		{0, 1, 1, 1, 0},
		{0, 0, 0, 0, 0},
	}

	// Verify result
	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			val, err := dilated.GetPixelAsInt8([]uint32{uint32(x), uint32(y)})
			if err != nil {
				t.Fatal(err)
			}
			if val != expected[y][x] {
				t.Errorf("at (%d,%d): expected %d, got %d", x, y, expected[y][x], val)
			}
		}
	}
}

func TestBinaryDilate3D(t *testing.T) {
	// Create a 3x3x3 test image with a single point in the middle
	data := [][][]int8{
		{{0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}},
		{{0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}},
		{{0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 1, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}},
		{{0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}},
		{{0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}},
	}
	img, err := GetImageFromArray(data)
	if err != nil {
		t.Fatal(err)
	}

	// Test with kernel size 3
	dilated, err := BinaryDilate(img, 3)
	if err != nil {
		t.Fatal(err)
	}

	// Verify middle slice
	expected := [][][]int8{
		{{0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}},
		{{0, 0, 0, 0, 0}, {0, 1, 1, 1, 0}, {0, 1, 1, 1, 0}, {0, 1, 1, 1, 0}, {0, 0, 0, 0, 0}},
		{{0, 0, 0, 0, 0}, {0, 1, 1, 1, 0}, {0, 1, 1, 1, 0}, {0, 1, 1, 1, 0}, {0, 0, 0, 0, 0}},
		{{0, 0, 0, 0, 0}, {0, 1, 1, 1, 0}, {0, 1, 1, 1, 0}, {0, 1, 1, 1, 0}, {0, 0, 0, 0, 0}},
		{{0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}},
	}

	for z := 0; z < 5; z++ {
		for y := 0; y < 5; y++ {
			for x := 0; x < 5; x++ {
				val, err := dilated.GetPixelAsInt8([]uint32{uint32(x), uint32(y), uint32(z)})
				if err != nil {
					t.Fatal(err)
				}
				if val != expected[z][y][x] {
					t.Errorf("at (%d,%d,%d): expected %d, got %d", x, y, z, expected[z][y][x], val)
				}
			}
		}
	}
}

func TestBinaryErode2D(t *testing.T) {
	// Create a 5x5 test image with a 3x3 filled square
	data := [][]int8{
		{0, 0, 0, 0, 0},
		{0, 1, 1, 1, 0},
		{0, 1, 1, 1, 0},
		{0, 1, 1, 1, 0},
		{0, 0, 0, 0, 0},
	}
	img, err := GetImageFromArray(data)
	if err != nil {
		t.Fatal(err)
	}

	// Test with kernel size 3
	eroded, err := BinaryErode(img, 3)
	if err != nil {
		t.Fatal(err)
	}

	// Expected result with 3x3 kernel - only center pixel should remain
	expected := [][]int8{
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
	}

	// Verify result
	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			val, err := eroded.GetPixelAsInt8([]uint32{uint32(x), uint32(y)})
			if err != nil {
				t.Fatal(err)
			}
			if val != expected[y][x] {
				t.Errorf("at (%d,%d): expected %d, got %d", x, y, expected[y][x], val)
			}
		}
	}
}

func TestBinaryErode3D(t *testing.T) {
	// Create a 5x5x5 test image with a 3x3x3 filled cube
	size := 5
	data := [][][]int8{
		{{0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}},
		{{0, 0, 0, 0, 0}, {0, 1, 1, 1, 0}, {0, 1, 1, 1, 0}, {0, 1, 1, 1, 0}, {0, 0, 0, 0, 0}},
		{{0, 0, 0, 0, 0}, {0, 1, 1, 1, 0}, {0, 1, 1, 1, 0}, {0, 1, 1, 1, 0}, {0, 0, 0, 0, 0}},
		{{0, 0, 0, 0, 0}, {0, 1, 1, 1, 0}, {0, 1, 1, 1, 0}, {0, 1, 1, 1, 0}, {0, 0, 0, 0, 0}},
		{{0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}},
	}

	img, err := GetImageFromArray(data)
	if err != nil {
		t.Fatal(err)
	}

	// Test with kernel size 3
	eroded, err := BinaryErode(img, 3)
	if err != nil {
		t.Fatal(err)
	}

	// After erosion, only the center voxel should remain
	for z := 0; z < size; z++ {
		for y := 0; y < size; y++ {
			for x := 0; x < size; x++ {
				val, err := eroded.GetPixelAsInt8([]uint32{uint32(x), uint32(y), uint32(z)})
				if err != nil {
					t.Fatal(err)
				}

				// Only the center voxel (2,2,2) should be 1
				expected := int8(0)
				if x == 2 && y == 2 && z == 2 {
					expected = 1
				}

				if val != expected {
					t.Errorf("at (%d,%d,%d): expected %d, got %d", x, y, z, expected, val)
				}
			}
		}
	}
}

func TestMorphologyClose2D(t *testing.T) {
	// Create a 5x5 test image with a 3x3 filled square
	data := [][]int8{
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
	}
	img, err := GetImageFromArray(data)
	if err != nil {
		t.Fatal(err)
	}

	// Perform morphological closing with 3x3 kernel and 1 iteration
	closed, err := Morphology(img, MORPH_CLOSE, 3, 1)
	if err != nil {
		t.Fatal(err)
	}

	// Expected result after closing
	expected := [][]int8{
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
	}

	// Verify result
	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			val, err := closed.GetPixelAsInt8([]uint32{uint32(x), uint32(y)})
			if err != nil {
				t.Fatal(err)
			}
			if val != expected[y][x] {
				t.Errorf("at (%d,%d): expected %d, got %d", x, y, expected[y][x], val)
			}
		}
	}
}

func TestMorphologyOpen2D(t *testing.T) {
	// Create a 5x5 test image with a 3x3 filled square
	data := [][]int8{
		{0, 0, 0, 0, 0},
		{0, 1, 1, 1, 0},
		{0, 1, 1, 1, 0},
		{0, 1, 1, 1, 0},
		{0, 0, 0, 0, 0},
	}
	img, err := GetImageFromArray(data)
	if err != nil {
		t.Fatal(err)
	}

	// Perform morphological opening with 3x3 kernel and 1 iteration
	opened, err := Morphology(img, MORPH_OPEN, 3, 1)
	if err != nil {
		t.Fatal(err)
	}

	// Expected result after opening
	expected := [][]int8{
		{0, 0, 0, 0, 0},
		{0, 1, 1, 1, 0},
		{0, 1, 1, 1, 0},
		{0, 1, 1, 1, 0},
		{0, 0, 0, 0, 0},
	}

	// Verify result
	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			val, err := opened.GetPixelAsInt8([]uint32{uint32(x), uint32(y)})
			if err != nil {
				t.Fatal(err)
			}
			if val != expected[y][x] {
				t.Errorf("at (%d,%d): expected %d, got %d", x, y, expected[y][x], val)
			}
		}
	}
}

func TestMorphologyClose3D(t *testing.T) {
	// Create a 5x5x5 test image with a 3x3x3 filled cube
	size := 5
	data := [][][]int8{
		{{0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}},
		{{0, 0, 0, 0, 0}, {0, 1, 1, 1, 0}, {0, 1, 1, 1, 0}, {0, 1, 1, 1, 0}, {0, 0, 0, 0, 0}},
		{{0, 0, 0, 0, 0}, {0, 1, 1, 1, 0}, {0, 1, 1, 1, 0}, {0, 1, 1, 1, 0}, {0, 0, 0, 0, 0}},
		{{0, 0, 0, 0, 0}, {0, 1, 1, 1, 0}, {0, 1, 1, 1, 0}, {0, 1, 1, 1, 0}, {0, 0, 0, 0, 0}},
		{{0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}},
	}

	img, err := GetImageFromArray(data)
	if err != nil {
		t.Fatal(err)
	}

	// Perform morphological closing with 3x3 kernel and 1 iteration
	closed, err := Morphology(img, MORPH_CLOSE, 3, 1)
	if err != nil {
		t.Fatal(err)
	}

	// Expected result after closing
	expected := [][][]int8{
		{{0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}},
		{{0, 0, 0, 0, 0}, {0, 1, 1, 1, 0}, {0, 1, 1, 1, 0}, {0, 1, 1, 1, 0}, {0, 0, 0, 0, 0}},
		{{0, 0, 0, 0, 0}, {0, 1, 1, 1, 0}, {0, 1, 1, 1, 0}, {0, 1, 1, 1, 0}, {0, 0, 0, 0, 0}},
		{{0, 0, 0, 0, 0}, {0, 1, 1, 1, 0}, {0, 1, 1, 1, 0}, {0, 1, 1, 1, 0}, {0, 0, 0, 0, 0}},
		{{0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}},
	}

	// Verify result
	for z := 0; z < size; z++ {
		for y := 0; y < size; y++ {
			for x := 0; x < size; x++ {
				val, err := closed.GetPixelAsInt8([]uint32{uint32(x), uint32(y), uint32(z)})
				if err != nil {
					t.Fatal(err)
				}
				if val != expected[z][y][x] {
					t.Errorf("at (%d,%d,%d): expected %d, got %d", x, y, z, expected[z][y][x], val)
				}
			}
		}
	}
}

func TestMorphologyOpen3D(t *testing.T) {
	// Create a 5x5x5 test image with a 3x3x3 filled cube
	size := 5
	data := [][][]int8{
		{{0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}},
		{{0, 0, 0, 0, 0}, {0, 1, 1, 1, 0}, {0, 1, 1, 1, 0}, {0, 1, 1, 1, 0}, {0, 0, 0, 0, 0}},
		{{0, 0, 0, 0, 0}, {0, 1, 1, 1, 0}, {0, 1, 1, 1, 0}, {0, 1, 1, 1, 0}, {0, 0, 0, 0, 0}},
		{{0, 0, 0, 0, 0}, {0, 1, 1, 1, 0}, {0, 1, 1, 1, 0}, {0, 1, 1, 1, 0}, {0, 0, 0, 0, 0}},
		{{0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}},
	}

	img, err := GetImageFromArray(data)
	if err != nil {
		t.Fatal(err)
	}

	// Perform morphological opening with 3x3 kernel and 1 iteration
	opened, err := Morphology(img, MORPH_OPEN, 3, 1)
	if err != nil {
		t.Fatal(err)
	}

	// Expected result after opening
	expected := [][][]int8{
		{{0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}},
		{{0, 0, 0, 0, 0}, {0, 1, 1, 1, 0}, {0, 1, 1, 1, 0}, {0, 1, 1, 1, 0}, {0, 0, 0, 0, 0}},
		{{0, 0, 0, 0, 0}, {0, 1, 1, 1, 0}, {0, 1, 1, 1, 0}, {0, 1, 1, 1, 0}, {0, 0, 0, 0, 0}},
		{{0, 0, 0, 0, 0}, {0, 1, 1, 1, 0}, {0, 1, 1, 1, 0}, {0, 1, 1, 1, 0}, {0, 0, 0, 0, 0}},
		{{0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}},
	}

	// Verify result
	for z := 0; z < size; z++ {
		for y := 0; y < size; y++ {
			for x := 0; x < size; x++ {
				val, err := opened.GetPixelAsInt8([]uint32{uint32(x), uint32(y), uint32(z)})
				if err != nil {
					t.Fatal(err)
				}
				if val != expected[z][y][x] {
					t.Errorf("at (%d,%d,%d): expected %d, got %d", x, y, z, expected[z][y][x], val)
				}
			}
		}
	}
}
