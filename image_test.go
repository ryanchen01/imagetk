package imagetk

import (
	"testing"
)

func TestImageCreation(t *testing.T) {
	// Create a new image
	img, err := NewImage([]uint32{10, 11, 12}, PixelTypeFloat32)
	if err != nil {
		t.Errorf("Error creating image: %v", err)
	}

	size := img.GetSize()
	// Check the size
	if size[0] != 10 || size[1] != 11 || size[2] != 12 {
		t.Errorf("Expected size to be [10, 11, 12], got %v", size)
	}

	dimension := img.GetDimension()
	// Check the dimension
	if dimension != 3 {
		t.Errorf("Expected dimension to be 3, got %v", dimension)
	}

	// Check the pixel type
	if img.GetPixelType() != PixelTypeFloat32 {
		t.Errorf("Expected pixel type to be PixelTypeFloat32, got %v", img.GetPixelType())
	}

	arr := make([][][]float32, 12)
	for z := 0; z < 12; z++ {
		arr[z] = make([][]float32, 11)
		for y := 0; y < 11; y++ {
			arr[z][y] = make([]float32, 10)
			for x := 0; x < 10; x++ {
				arr[z][y][x] = float32(z*11*10 + y*10 + x)
			}
		}
	}

	img, err = GetImageFromArray(arr)
	if err != nil {
		t.Errorf("Error creating image from array: %v", err)
	}

	size = img.GetSize()
	// Check the size
	if size[0] != 10 || size[1] != 11 || size[2] != 12 {
		t.Errorf("Expected size to be [10, 11, 12], got %v", size)
	}

	dimension = img.GetDimension()
	// Check the dimension
	if dimension != 3 {
		t.Errorf("Expected dimension to be 3, got %v", dimension)
	}

	// Check the pixel type
	if img.GetPixelType() != PixelTypeFloat32 {
		t.Errorf("Expected pixel type to be PixelTypeFloat32, got %v", img.GetPixelType())
	}

	// Check the pixel values
	for z := 0; z < 12; z++ {
		for y := 0; y < 11; y++ {
			for x := 0; x < 10; x++ {
				pixel, err := img.GetPixel([]uint32{uint32(x), uint32(y), uint32(z)})
				if err != nil {
					t.Errorf("Error getting pixel: %v", err)
				}

				if pixel.(float32) != float32(z*11*10+y*10+x) {
					t.Errorf("Expected pixel value to be %v, got %v", float32(z*11*10+y*10+x), pixel)
				}
			}
		}
	}
}

func TestImageCreationWithInvalidSize(t *testing.T) {
	_, err := NewImage([]uint32{10, 0, 12}, PixelTypeFloat32)
	if err == nil {
		t.Errorf("Expected error when creating image with invalid size")
	}

	_, err = NewImage([]uint32{10, 11, 0}, PixelTypeFloat32)
	if err == nil {
		t.Errorf("Expected error when creating image with invalid size")
	}

	_, err = NewImage([]uint32{}, PixelTypeFloat32)
	if err == nil {
		t.Errorf("Expected error when creating image with invalid size")
	}
}

func TestSinglePixelSetting(t *testing.T) {
	// Create a new image
	img, err := NewImage([]uint32{10, 10, 10}, PixelTypeFloat32)
	if err != nil {
		t.Errorf("Error creating image: %v", err)
	}

	// Set a pixel
	err = img.SetPixel([]uint32{0, 0, 0}, float32(1.0))
	if err != nil {
		t.Errorf("Error setting pixel: %v", err)
	}

	// Get a pixel
	pixel, err := img.GetPixel([]uint32{0, 0, 0})
	if err != nil {
		t.Errorf("Error getting pixel: %v", err)
	}

	// Check if the pixel is correct
	if pixel.(float32) != 1.0 {
		t.Errorf("Expected pixel value to be 1.0, got %v", pixel)
	}
}

func TestSinglePixelSettingWithInvalidCoordinates(t *testing.T) {
	// Create a new image
	img, err := NewImage([]uint32{10, 10, 10}, PixelTypeFloat32)
	if err != nil {
		t.Errorf("Error creating image: %v", err)
	}

	// Set a pixel with invalid coordinates
	err = img.SetPixel([]uint32{11, 0, 0}, float32(1.0))
	if err == nil {
		t.Errorf("Expected error when setting pixel with invalid coordinates")
	}
}

func TestSinglePixelSettingWithInvalidType(t *testing.T) {
	// Create a new image
	img, err := NewImage([]uint32{10, 10, 10}, PixelTypeFloat32)
	if err != nil {
		t.Errorf("Error creating image: %v", err)
	}

	// Set a pixel with invalid type
	err = img.SetPixel([]uint32{0, 0, 0}, 1)
	if err == nil {
		t.Errorf("Expected error when setting pixel with invalid type")
	}
}

func TestPixelsSetting(t *testing.T) {
	// Create a new image
	img, err := NewImage([]uint32{10, 10, 10}, PixelTypeFloat32)
	if err != nil {
		t.Errorf("Error creating image: %v", err)
	}

	pixels := make([]float32, 10*10*10)
	for i := 0; i < 10*10*10; i++ {
		pixels[i] = float32(i)
	}

	// Set pixels
	err = img.SetPixels(pixels)
	if err != nil {
		t.Errorf("Error setting pixels: %v", err)
	}

	// Get pixel
	gotPixel, err := img.GetPixel([]uint32{0, 0, 0})
	if err != nil {
		t.Errorf("Error getting pixel: %v", err)
	}

	// Check if the pixel is correct
	if gotPixel.(float32) != float32(0.0) {
		t.Errorf("Expected pixel value to be 0.0, got %v", gotPixel)
	}
}

func TestPixelsSettingWithInvalidSize(t *testing.T) {
	// Create a new image
	img, err := NewImage([]uint32{10, 10, 10}, PixelTypeFloat32)
	if err != nil {
		t.Errorf("Error creating image: %v", err)
	}

	pixels := make([]float32, 10*10*10-1)
	for i := 0; i < 10*10*10-1; i++ {
		pixels[i] = float32(i)
	}

	// Set pixels with invalid size
	err = img.SetPixels(pixels)
	if err == nil {
		t.Errorf("Expected error when setting pixels with invalid size")
	}

	pixels = make([]float32, 10*10*10+1)
	for i := 0; i < 10*10*10+1; i++ {
		pixels[i] = float32(i)
	}

	// Set pixels with invalid size
	err = img.SetPixels(pixels)
	if err == nil {
		t.Errorf("Expected error when setting pixels with invalid size")
	}
}

func TestPixelsSettingWithInvalidType(t *testing.T) {
	// Create a new image
	img, err := NewImage([]uint32{10, 10, 10}, PixelTypeFloat32)
	if err != nil {
		t.Errorf("Error creating image: %v", err)
	}

	pixels := make([]int, 10*10*10)
	for i := 0; i < 10*10*10; i++ {
		pixels[i] = i
	}

	// Set pixels with invalid type
	err = img.SetPixels(pixels)
	if err == nil {
		t.Errorf("Expected error when setting pixels with invalid type")
	}
}

func TestPixelGetting(t *testing.T) {
	// Create a new image
	img, err := NewImage([]uint32{10, 10, 10}, PixelTypeFloat32)
	if err != nil {
		t.Errorf("Error creating image: %v", err)
	}

	err = img.SetPixel([]uint32{0, 0, 0}, float32(0.5))
	if err != nil {
		t.Errorf("Error setting pixel: %v", err)
	}

	valUInt8, err := img.GetPixelAsUInt8([]uint32{0, 0, 0})
	if err != nil {
		t.Errorf("Error getting pixel as uint8: %v", err)
	}

	if valUInt8 != uint8(0) {
		t.Errorf("Expected pixel value to be 0, got %v", valUInt8)
	}

	valInt8, err := img.GetPixelAsInt8([]uint32{0, 0, 0})
	if err != nil {
		t.Errorf("Error getting pixel as int8: %v", err)
	}

	if valInt8 != int8(0) {
		t.Errorf("Expected pixel value to be 0, got %v", valInt8)
	}

	valUInt16, err := img.GetPixelAsUInt16([]uint32{0, 0, 0})
	if err != nil {
		t.Errorf("Error getting pixel as uint16: %v", err)
	}

	if valUInt16 != uint16(0) {
		t.Errorf("Expected pixel value to be 0, got %v", valUInt16)
	}

	valInt16, err := img.GetPixelAsInt16([]uint32{0, 0, 0})
	if err != nil {
		t.Errorf("Error getting pixel as int16: %v", err)
	}

	if valInt16 != int16(0) {
		t.Errorf("Expected pixel value to be 0, got %v", valInt16)
	}

	valUInt32, err := img.GetPixelAsUInt32([]uint32{0, 0, 0})
	if err != nil {
		t.Errorf("Error getting pixel as uint32: %v", err)
	}

	if valUInt32 != uint32(0) {
		t.Errorf("Expected pixel value to be 0, got %v", valUInt32)
	}

	valInt32, err := img.GetPixelAsInt32([]uint32{0, 0, 0})
	if err != nil {
		t.Errorf("Error getting pixel as int32: %v", err)
	}

	if valInt32 != int32(0) {
		t.Errorf("Expected pixel value to be 0, got %v", valInt32)
	}

	valUInt64, err := img.GetPixelAsUInt64([]uint32{0, 0, 0})
	if err != nil {
		t.Errorf("Error getting pixel as uint64: %v", err)
	}

	if valUInt64 != uint64(0) {
		t.Errorf("Expected pixel value to be 0, got %v", valUInt64)
	}

	valInt64, err := img.GetPixelAsInt64([]uint32{0, 0, 0})
	if err != nil {
		t.Errorf("Error getting pixel as int64: %v", err)
	}

	if valInt64 != int64(0) {
		t.Errorf("Expected pixel value to be 0, got %v", valInt64)
	}

	valFloat32, err := img.GetPixelAsFloat32([]uint32{0, 0, 0})
	if err != nil {
		t.Errorf("Error getting pixel as float32: %v", err)
	}

	if valFloat32 != float32(0.5) {
		t.Errorf("Expected pixel value to be 0.5, got %v", valFloat32)
	}

	valFloat64, err := img.GetPixelAsFloat64([]uint32{0, 0, 0})
	if err != nil {
		t.Errorf("Error getting pixel as float64: %v", err)
	}

	if valFloat64 != float64(0.5) {
		t.Errorf("Expected pixel value to be 0.5, got %v", valFloat64)
	}
}

func TestAsType(t *testing.T) {
	pixels := make([][]float32, 10)
	for i := 0; i < 10; i++ {
		pixels[i] = make([]float32, 10)
		for j := 0; j < 10; j++ {
			pixels[i][j] = float32(i*10 + j)
		}
	}

	// Create a new image
	img, err := GetImageFromArray(pixels)
	if err != nil {
		t.Errorf("Error creating image: %v", err)
	}

	uin8Img, err := img.AsType(PixelTypeUInt8)
	if err != nil {
		t.Errorf("Error converting image to uint8: %v", err)
	}

	if uin8Img.GetPixelType() != PixelTypeUInt8 {
		t.Errorf("Expected pixel type to be PixelTypeUInt8, got %v", uin8Img.GetPixelType())
	}

	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			pixel, err := uin8Img.GetPixel([]uint32{uint32(i), uint32(j)})
			if err != nil {
				t.Errorf("Error getting pixel: %v", err)
			}

			expected, err := img.GetPixelAsUInt8([]uint32{uint32(i), uint32(j)})
			if err != nil {
				t.Errorf("Error getting pixel as uint8: %v", err)
			}

			if pixel.(uint8) != expected {
				t.Errorf("Expected pixel value %v to be %v, got %v", expected, pixel, expected)
			}
		}
	}
}
