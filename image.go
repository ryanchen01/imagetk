// Package image provides functions for creating, manipulating, and analyzing images.
package imagetk

import (
	"encoding/binary"
	"fmt"
	"math"
	"reflect"
	"runtime"
	"sync"
)

const (
	// PixelTypeUnknown is an unknown pixel type.
	PixelTypeUnknown = iota
	// PixelTypeUInt8 is an unsigned 8-bit integer pixel type.
	PixelTypeUInt8
	// PixelTypeInt8 is a signed 8-bit integer pixel type.
	PixelTypeInt8
	// PixelTypeUInt16 is an unsigned 16-bit integer pixel type.
	PixelTypeUInt16
	// PixelTypeInt16 is a signed 16-bit integer pixel type.
	PixelTypeInt16
	// PixelTypeUInt32 is an unsigned 32-bit integer pixel type.
	PixelTypeUInt32
	// PixelTypeInt32 is a signed 32-bit integer pixel type.
	PixelTypeInt32
	// PixelTypeUInt64 is an unsigned 64-bit integer pixel type.
	PixelTypeUInt64
	// PixelTypeInt64 is a signed 64-bit integer pixel type.
	PixelTypeInt64
	// PixelTypeFloat32 is a 32-bit floating point pixel type.
	PixelTypeFloat32
	// PixelTypeFloat64 is a 64-bit floating point pixel type.
	PixelTypeFloat64
)

// Image represents an image with various properties such as pixels, pixel type,
// dimensions, size, spacing, origin, and direction.
//
// Fields:
//   - pixels: The pixel data of the image in an array of bytes.
//   - pixelType: An integer representing the type of pixels.
//   - dimension: The number of dimensions of the image (2<=N<=3).
//   - size: A slice of uint32 representing the size of the image in each dimension (2<=N<=3).
//   - spacing: A slice of float64 representing the spacing between pixels in each dimension.
//   - origin: A slice of float64 representing the origin of the image.
//   - direction: An array of 9 float64 values representing the direction cosines of the image.
type Image struct {
	pixels        []byte
	pixelType     int
	bytesPerPixel int
	dimension     uint32
	size          []uint32
	spacing       []float64
	origin        []float64
	direction     [9]float64
}

// NewImage creates a new Image with the specified size and pixel type.
// Parameters:
//   - size: A slice of uint32 representing the size of the image in each dimension.
//   - pixelType: An integer representing the type of pixels.
//
// Supported pixel types are:
//   - PixelTypeUInt8
//   - PixelTypeInt8
//   - PixelTypeUInt16
//   - PixelTypeInt16
//   - PixelTypeUInt32
//   - PixelTypeInt32
//   - PixelTypeUInt64
//   - PixelTypeInt64
//   - PixelTypeFloat32
//   - PixelTypeFloat64
//
// Returns:
//   - *Image: A pointer to the created Image.
//   - error: An error if the image creation fails.
func NewImage(size []uint32, pixelType int) (*Image, error) {
	if len(size) < 2 || len(size) > 3 {
		return nil, fmt.Errorf("invalid size length: %d", len(size))
	}

	numPixels := 1
	for _, s := range size {
		if s == 0 {
			return nil, fmt.Errorf("invalid size: %d", s)
		}
		numPixels *= int(s)
	}

	var pixels []byte
	var bytesPerPixel int
	switch pixelType {
	case PixelTypeUInt8, PixelTypeInt8:
		bytesPerPixel = 1
	case PixelTypeUInt16, PixelTypeInt16:
		bytesPerPixel = 2
	case PixelTypeUInt32, PixelTypeInt32, PixelTypeFloat32:
		bytesPerPixel = 4
	case PixelTypeUInt64, PixelTypeInt64, PixelTypeFloat64:
		bytesPerPixel = 8
	default:
		return nil, fmt.Errorf("unsupported pixel type: %d", pixelType)
	}

	pixels = make([]byte, numPixels*bytesPerPixel)

	spacing := make([]float64, len(size))
	origin := make([]float64, len(size))
	for i := 0; i < len(size); i++ {
		origin[i] = 0
		spacing[i] = 1
	}

	direction := [9]float64{1, 0, 0, 0, 1, 0, 0, 0, 1}

	return &Image{
		pixels:        pixels,
		pixelType:     pixelType,
		bytesPerPixel: bytesPerPixel,
		dimension:     uint32(len(size)),
		size:          size,
		spacing:       spacing,
		origin:        origin,
		direction:     direction,
	}, nil
}

// GetImageFromArray creates an Image from a multi-dimensional array of pixel data.
// The function determines the size and pixel type of the image based on the input data.
//
// Parameters:
//   - data: A multi-dimensional array containing the pixel data. The array can be of any shape.
//
// Returns:
//   - *Image: A pointer to the created Image.
//   - error: An error if the image creation or pixel setting fails.
//
// The function supports the following pixel types:
//   - uint8
//   - int8
//   - uint16
//   - int16
//   - uint32
//   - int32
//   - uint64
//   - int64
//   - float32
//   - float64
//
// If the pixel type is not supported, the function returns an error.
func GetImageFromArray(data any) (*Image, error) {
	var _size []uint32
	value := reflect.ValueOf(data)

	if value.Kind() != reflect.Slice && value.Kind() != reflect.Array {
		return nil, fmt.Errorf("data must be a slice or array, got %s", value.Kind().String())
	}
	pixelType := PixelTypeUnknown

	for value.Kind() == reflect.Slice {
		_size = append(_size, uint32(value.Len()))
		if value.Len() == 0 {
			break
		}
		value = value.Index(0)
	}

	size := make([]uint32, len(_size))
	for i := 0; i < len(_size); i++ {
		size[i] = _size[len(_size)-1-i]
	}

	if len(size) < 2 || len(size) > 3 {
		return nil, fmt.Errorf("invalid dimension: %d", len(size))
	}

	// Determine the pixel type based on the kind of the reflect value.
	// This switch case maps the reflect.Kind to the corresponding PixelType constant.
	switch value.Kind() {
	case reflect.Uint8:
		pixelType = PixelTypeUInt8
	case reflect.Int8:
		pixelType = PixelTypeInt8
	case reflect.Uint16:
		pixelType = PixelTypeUInt16
	case reflect.Int16:
		pixelType = PixelTypeInt16
	case reflect.Uint32:
		pixelType = PixelTypeUInt32
	case reflect.Int32:
		pixelType = PixelTypeInt32
	case reflect.Uint64:
		pixelType = PixelTypeUInt64
	case reflect.Int64:
		pixelType = PixelTypeInt64
	case reflect.Float32:
		pixelType = PixelTypeFloat32
	case reflect.Float64:
		pixelType = PixelTypeFloat64
	default:
		return nil, fmt.Errorf("unsupported pixel type: %s", value.Kind().String())
	}

	img, err := NewImage(size, pixelType)
	if err != nil {
		return nil, err
	}

	err = img.SetPixels(data)
	if err != nil {
		return nil, err
	}

	return img, nil
}

func GetArrayFromImage(img *Image) (any, error) {
	numGoroutines := uint64(runtime.NumCPU())
	chunkSize := uint64(img.NumPixels()) / numGoroutines
	if chunkSize*numGoroutines < uint64(img.NumPixels()) {
		chunkSize += 1
	}
	wg := sync.WaitGroup{}
	switch img.pixelType {
	case PixelTypeUInt8:
		data := make([]uint8, img.NumPixels())
		for chunk := uint64(0); chunk < numGoroutines; chunk++ {
			start := chunk * chunkSize
			end := start + chunkSize
			if end > img.NumPixels() {
				end = img.NumPixels()
			}
			wg.Add(1)
			go func(start, end uint64) {
				defer wg.Done()
				for i := start; i < end; i++ {
					index, err := img.GetIndexFromLinearIndex(i)
					if err != nil {
						return
					}
					value, err := img.GetPixelAsUInt8(index)
					if err != nil {
						return
					}
					data[i] = value
				}
			}(start, end)
		}
		wg.Wait()
		shapedData := reshape(data, img.size)
		return shapedData, nil
	case PixelTypeInt8:
		data := make([]int8, img.NumPixels())
		for chunk := uint64(0); chunk < numGoroutines; chunk++ {
			start := chunk * chunkSize
			end := start + chunkSize
			if end > img.NumPixels() {
				end = img.NumPixels()
			}
			wg.Add(1)
			go func(start, end uint64) {
				defer wg.Done()
				for i := start; i < end; i++ {
					index, err := img.GetIndexFromLinearIndex(i)
					if err != nil {
						return
					}
					value, err := img.GetPixelAsInt8(index)
					if err != nil {
						return
					}
					data[i] = value
				}
			}(start, end)
		}
		wg.Wait()
		shapedData := reshape(data, img.size)
		return shapedData, nil
	case PixelTypeUInt16:
		data := make([]uint16, img.NumPixels())
		for chunk := uint64(0); chunk < numGoroutines; chunk++ {
			start := chunk * chunkSize
			end := start + chunkSize
			if end > img.NumPixels() {
				end = img.NumPixels()
			}
			wg.Add(1)
			go func(start, end uint64) {
				defer wg.Done()
				for i := start; i < end; i++ {
					index, err := img.GetIndexFromLinearIndex(i)
					if err != nil {
						return
					}
					value, err := img.GetPixelAsUInt16(index)
					if err != nil {
						return
					}
					data[i] = value
				}
			}(start, end)
		}
		wg.Wait()
		shapedData := reshape(data, img.size)
		return shapedData, nil
	case PixelTypeInt16:
		data := make([]int16, img.NumPixels())
		for chunk := uint64(0); chunk < numGoroutines; chunk++ {
			start := chunk * chunkSize
			end := start + chunkSize
			if end > img.NumPixels() {
				end = img.NumPixels()
			}
			wg.Add(1)
			go func(start, end uint64) {
				defer wg.Done()
				for i := start; i < end; i++ {
					index, err := img.GetIndexFromLinearIndex(i)
					if err != nil {
						return
					}
					value, err := img.GetPixelAsInt16(index)
					if err != nil {
						return
					}
					data[i] = value
				}
			}(start, end)
		}
		wg.Wait()
		shapedData := reshape(data, img.size)
		return shapedData, nil
	case PixelTypeUInt32:
		data := make([]uint32, img.NumPixels())
		for chunk := uint64(0); chunk < numGoroutines; chunk++ {
			start := chunk * chunkSize
			end := start + chunkSize
			if end > img.NumPixels() {
				end = img.NumPixels()
			}
			wg.Add(1)
			go func(start, end uint64) {
				defer wg.Done()
				for i := start; i < end; i++ {
					index, err := img.GetIndexFromLinearIndex(i)
					if err != nil {
						return
					}
					value, err := img.GetPixelAsUInt32(index)
					if err != nil {
						return
					}
					data[i] = value
				}
			}(start, end)
		}
		wg.Wait()
		shapedData := reshape(data, img.size)
		return shapedData, nil
	case PixelTypeInt32:
		data := make([]int32, img.NumPixels())
		for chunk := uint64(0); chunk < numGoroutines; chunk++ {
			start := chunk * chunkSize
			end := start + chunkSize
			if end > img.NumPixels() {
				end = img.NumPixels()
			}
			wg.Add(1)
			go func(start, end uint64) {
				defer wg.Done()
				for i := start; i < end; i++ {
					index, err := img.GetIndexFromLinearIndex(i)
					if err != nil {
						return
					}
					value, err := img.GetPixelAsInt32(index)
					if err != nil {
						return
					}
					data[i] = value
				}
			}(start, end)
		}
		wg.Wait()
		shapedData := reshape(data, img.size)
		return shapedData, nil
	case PixelTypeUInt64:
		data := make([]uint64, img.NumPixels())
		for chunk := uint64(0); chunk < numGoroutines; chunk++ {
			start := chunk * chunkSize
			end := start + chunkSize
			if end > img.NumPixels() {
				end = img.NumPixels()
			}
			wg.Add(1)
			go func(start, end uint64) {
				defer wg.Done()
				for i := start; i < end; i++ {
					index, err := img.GetIndexFromLinearIndex(i)
					if err != nil {
						return
					}
					value, err := img.GetPixelAsUInt64(index)
					if err != nil {
						return
					}
					data[i] = value
				}
			}(start, end)
		}
		wg.Wait()
		shapedData := reshape(data, img.size)
		return shapedData, nil
	case PixelTypeInt64:
		data := make([]int64, img.NumPixels())
		for chunk := uint64(0); chunk < numGoroutines; chunk++ {
			start := chunk * chunkSize
			end := start + chunkSize
			if end > img.NumPixels() {
				end = img.NumPixels()
			}
			wg.Add(1)
			go func(start, end uint64) {
				defer wg.Done()
				for i := start; i < end; i++ {
					index, err := img.GetIndexFromLinearIndex(i)
					if err != nil {
						return
					}
					value, err := img.GetPixelAsInt64(index)
					if err != nil {
						return
					}
					data[i] = value
				}
			}(start, end)
		}
		wg.Wait()
		shapedData := reshape(data, img.size)
		return shapedData, nil
	case PixelTypeFloat32:
		data := make([]float32, img.NumPixels())
		for chunk := uint64(0); chunk < numGoroutines; chunk++ {
			start := chunk * chunkSize
			end := start + chunkSize
			if end > img.NumPixels() {
				end = img.NumPixels()
			}
			wg.Add(1)
			go func(start, end uint64) {
				defer wg.Done()
				for i := start; i < end; i++ {
					index, err := img.GetIndexFromLinearIndex(i)
					if err != nil {
						return
					}
					value, err := img.GetPixelAsFloat32(index)
					if err != nil {
						return
					}
					data[i] = value
				}
			}(start, end)
		}
		wg.Wait()
		shapedData := reshape(data, img.size)
		return shapedData, nil
	case PixelTypeFloat64:
		data := make([]float64, img.NumPixels())
		for chunk := uint64(0); chunk < numGoroutines; chunk++ {
			start := chunk * chunkSize
			end := start + chunkSize
			if end > img.NumPixels() {
				end = img.NumPixels()
			}
			wg.Add(1)
			go func(start, end uint64) {
				defer wg.Done()
				for i := start; i < end; i++ {
					index, err := img.GetIndexFromLinearIndex(i)
					if err != nil {
						return
					}
					value, err := img.GetPixelAsFloat64(index)
					if err != nil {
						return
					}
					data[i] = value
				}
			}(start, end)
		}
		wg.Wait()
		shapedData := reshape(data, img.size)
		return shapedData, nil
	default:
		return nil, fmt.Errorf("unsupported pixel type: %d", img.pixelType)
	}
}

// GetPixelType returns the pixel type of the image.
// Returns:
//   - int: An integer representing the pixel type.
//
// The function returns one of the following pixel types:
//   - PixelTypeUInt8
//   - PixelTypeInt8
//   - PixelTypeUInt16
//   - PixelTypeInt16
//   - PixelTypeUInt32
//   - PixelTypeInt32
//   - PixelTypeUInt64
//   - PixelTypeInt64
//   - PixelTypeFloat32
//   - PixelTypeFloat64
func (img *Image) GetPixelType() int {
	return img.pixelType
}

// IsPixelType checks if the pixel type of the image matches the given pixel type.
func (img *Image) IsPixelType(pixelType int) bool {
	return img.pixelType == pixelType
}

// GetDimension returns the number of dimensions of the image.
func (img *Image) GetDimension() uint32 {
	return img.dimension
}

// GetSize returns the size of the image.
func (img *Image) GetSize() []uint32 {
	return img.size
}

// GetSpacing returns the spacing of the image.
func (img *Image) GetSpacing() []float64 {
	return img.spacing
}

// GetOrigin returns the origin of the image.
func (img *Image) GetOrigin() []float64 {
	return img.origin
}

// GetDirection returns the direction matrix of the image.
func (img *Image) GetDirection() [9]float64 {
	return img.direction
}

// GetPixel returns the pixel value at the given index.
// Parameters:
//   - index: A slice of uint32 representing the index of the pixel.
//
// Returns:
//   - any: The pixel value as the type of the image.
//   - error: An error if the index is out of range.
func (img *Image) GetPixel(index []uint32) (any, error) {
	if len(index) != int(img.dimension) {
		return nil, fmt.Errorf("invalid index length: %d", len(index))
	}
	idx := uint32(0)
	for i := len(index) - 1; i >= 0; i-- {
		if index[i] >= img.size[i] {
			return nil, fmt.Errorf("index out of range: %d", index[i])
		}
		idx = idx*img.size[i] + index[i]
	}
	switch img.pixelType {
	case PixelTypeUInt8:
		value := uint8(img.pixels[idx])
		return value, nil
	case PixelTypeInt8:
		value := int8(img.pixels[idx])
		return value, nil
	case PixelTypeUInt16:
		value := binary.LittleEndian.Uint16(img.pixels[idx*2 : idx*2+2])
		return value, nil
	case PixelTypeInt16:
		value := int16(binary.LittleEndian.Uint16(img.pixels[idx*2 : idx*2+2]))
		return value, nil
	case PixelTypeUInt32:
		value := binary.LittleEndian.Uint32(img.pixels[idx*4 : idx*4+4])
		return value, nil
	case PixelTypeInt32:
		value := int32(binary.LittleEndian.Uint32(img.pixels[idx*4 : idx*4+4]))
		return value, nil
	case PixelTypeUInt64:
		value := binary.LittleEndian.Uint64(img.pixels[idx*8 : idx*8+8])
		return value, nil
	case PixelTypeInt64:
		value := int64(binary.LittleEndian.Uint64(img.pixels[idx*8 : idx*8+8]))
		return value, nil
	case PixelTypeFloat32:
		value := math.Float32frombits(binary.LittleEndian.Uint32(img.pixels[idx*4 : idx*4+4]))
		return value, nil
	case PixelTypeFloat64:
		value := math.Float64frombits(binary.LittleEndian.Uint64(img.pixels[idx*8 : idx*8+8]))
		return value, nil
	default:
		return nil, fmt.Errorf("unsupported pixel type: %d", img.pixelType)
	}
}

// GetPixelAsUInt8 returns the pixel value at the given index as a uint8.
// Parameters:
//   - index: A slice of uint32 representing the index of the pixel.
//
// Returns:
//   - uint8: The pixel value as a uint8.
//   - error: An error if the index is out of range.
func (img *Image) GetPixelAsUInt8(index []uint32) (uint8, error) {
	value, err := img.GetPixel(index)
	if err != nil {
		return 0, err
	}
	if converter, ok := pixelTypeConverters[PixelTypeUInt8]; ok {
		switch value.(type) {
		case uint8, int8, uint16, int16, uint32, int32, uint64, int64, float32, float64, int:
			return converter(value).(uint8), nil
		default:
			return 0, fmt.Errorf("unsupported value type")
		}
	}
	return 0, fmt.Errorf("unknown pixel type")
}

// GetPixelAsInt8 returns the pixel value at the given index as a int8.
// Parameters:
//   - index: A slice of uint32 representing the index of the pixel.
//
// Returns:
//   - int8: The pixel value as a int8.
//   - error: An error if the index is out of range.
func (img *Image) GetPixelAsInt8(index []uint32) (int8, error) {
	value, err := img.GetPixel(index)
	if err != nil {
		return 0, err
	}
	if converter, ok := pixelTypeConverters[PixelTypeInt8]; ok {
		switch value.(type) {
		case uint8, int8, uint16, int16, uint32, int32, uint64, int64, float32, float64, int:
			return converter(value).(int8), nil
		default:
			return 0, fmt.Errorf("unsupported value type")
		}
	}
	return 0, fmt.Errorf("unknown pixel type")
}

// GetPixelAsUInt16 returns the pixel value at the given index as a uint16.
// Parameters:
//   - index: A slice of uint32 representing the index of the pixel.
//
// Returns:
//   - uint16: The pixel value as a uint16.
//   - error: An error if the index is out of range.
func (img *Image) GetPixelAsUInt16(index []uint32) (uint16, error) {
	value, err := img.GetPixel(index)
	if err != nil {
		return 0, err
	}
	if converter, ok := pixelTypeConverters[PixelTypeUInt16]; ok {
		switch value.(type) {
		case uint8, int8, uint16, int16, uint32, int32, uint64, int64, float32, float64, int:
			return converter(value).(uint16), nil
		default:
			return 0, fmt.Errorf("unsupported value type")
		}
	}
	return 0, fmt.Errorf("unknown pixel type")
}

// GetPixelAsInt16 returns the pixel value at the given index as a int16.
// Parameters:
//   - index: A slice of uint32 representing the index of the pixel.
//
// Returns:
//   - int16: The pixel value as a int16.
//   - error: An error if the index is out of range.
func (img *Image) GetPixelAsInt16(index []uint32) (int16, error) {
	value, err := img.GetPixel(index)
	if err != nil {
		return 0, err
	}
	if converter, ok := pixelTypeConverters[PixelTypeInt16]; ok {
		switch value.(type) {
		case uint8, int8, uint16, int16, uint32, int32, uint64, int64, float32, float64, int:
			return converter(value).(int16), nil
		default:
			return 0, fmt.Errorf("unsupported value type")
		}
	}
	return 0, fmt.Errorf("unknown pixel type")
}

// GetPixelAsUInt32 returns the pixel value at the given index as a uint32.
// Parameters:
//   - index: A slice of uint32 representing the index of the pixel.
//
// Returns:
//   - uint32: The pixel value as a uint32.
//   - error: An error if the index is out of range.
func (img *Image) GetPixelAsUInt32(index []uint32) (uint32, error) {
	value, err := img.GetPixel(index)
	if err != nil {
		return 0, err
	}
	if converter, ok := pixelTypeConverters[PixelTypeUInt32]; ok {
		switch value.(type) {
		case uint8, int8, uint16, int16, uint32, int32, uint64, int64, float32, float64, int:
			return converter(value).(uint32), nil
		default:
			return 0, fmt.Errorf("unsupported value type")
		}
	}
	return 0, fmt.Errorf("unknown pixel type")
}

// GetPixelAsInt32 returns the pixel value at the given index as a int32.
// Parameters:
//   - index: A slice of uint32 representing the index of the pixel.
//
// Returns:
//   - int32: The pixel value as a int32.
//   - error: An error if the index is out of range.
func (img *Image) GetPixelAsInt32(index []uint32) (int32, error) {
	value, err := img.GetPixel(index)
	if err != nil {
		return 0, err
	}
	if converter, ok := pixelTypeConverters[PixelTypeInt32]; ok {
		switch value.(type) {
		case uint8, int8, uint16, int16, uint32, int32, uint64, int64, float32, float64, int:
			return converter(value).(int32), nil
		default:
			return 0, fmt.Errorf("unsupported value type")
		}
	}
	return 0, fmt.Errorf("unknown pixel type")
}

// GetPixelAsUInt64 returns the pixel value at the given index as a uint64.
// Parameters:
//   - index: A slice of uint32 representing the index of the pixel.
//
// Returns:
//   - uint64: The pixel value as a uint64.
//   - error: An error if the index is out of range.
func (img *Image) GetPixelAsUInt64(index []uint32) (uint64, error) {
	value, err := img.GetPixel(index)
	if err != nil {
		return 0, err
	}
	if converter, ok := pixelTypeConverters[PixelTypeUInt64]; ok {
		switch value.(type) {
		case uint8, int8, uint16, int16, uint32, int32, uint64, int64, float32, float64, int:
			return converter(value).(uint64), nil
		default:
			return 0, fmt.Errorf("unsupported value type")
		}
	}
	return 0, fmt.Errorf("unknown pixel type")
}

// GetPixelAsInt64 returns the pixel value at the given index as a int64.
// Parameters:
//   - index: A slice of uint32 representing the index of the pixel.
//
// Returns:
//   - int64: The pixel value as a int64.
//   - error: An error if the index is out of range.
func (img *Image) GetPixelAsInt64(index []uint32) (int64, error) {
	value, err := img.GetPixel(index)
	if err != nil {
		return 0, err
	}
	if converter, ok := pixelTypeConverters[PixelTypeInt64]; ok {
		switch value.(type) {
		case uint8, int8, uint16, int16, uint32, int32, uint64, int64, float32, float64, int:
			return converter(value).(int64), nil
		default:
			return 0, fmt.Errorf("unsupported value type")
		}
	}
	return 0, fmt.Errorf("unknown pixel type")
}

// GetPixelAsFloat32 returns the pixel value at the given index as a float32.
// Parameters:
//   - index: A slice of uint32 representing the index of the pixel.
//
// Returns:
//   - float32: The pixel value as a float32.
//   - error: An error if the index is out of range.
func (img *Image) GetPixelAsFloat32(index []uint32) (float32, error) {
	value, err := img.GetPixel(index)
	if err != nil {
		return 0, err
	}
	if converter, ok := pixelTypeConverters[PixelTypeFloat32]; ok {
		switch value.(type) {
		case uint8, int8, uint16, int16, uint32, int32, uint64, int64, float32, float64, int:
			return converter(value).(float32), nil
		default:
			return 0, fmt.Errorf("unsupported value type")
		}
	}
	return 0, fmt.Errorf("unknown pixel type")
}

// GetPixelAsFloat64 returns the pixel value at the given index as a float64.
// Parameters:
//   - index: A slice of uint32 representing the index of the pixel.
//
// Returns:
//   - float64: The pixel value as a float64.
//   - error: An error if the index is out of range.
func (img *Image) GetPixelAsFloat64(index []uint32) (float64, error) {
	value, err := img.GetPixel(index)
	if err != nil {
		return 0, err
	}
	if converter, ok := pixelTypeConverters[PixelTypeFloat64]; ok {
		switch value.(type) {
		case uint8, int8, uint16, int16, uint32, int32, uint64, int64, float32, float64, int:
			return converter(value).(float64), nil
		default:
			return 0, fmt.Errorf("unsupported value type")
		}
	}
	return 0, fmt.Errorf("unknown pixel type")
}

// AsType creates a new Image with the specified pixel type and copies the pixel data from the current image.
// Parameters:
//   - pixelType: An integer representing the type of pixels.
//
// Supported pixel types are:
//   - PixelTypeUInt8
//   - PixelTypeInt8
//   - PixelTypeUInt16
//   - PixelTypeInt16
//   - PixelTypeUInt32
//   - PixelTypeInt32
//   - PixelTypeUInt64
//   - PixelTypeInt64
//   - PixelTypeFloat32
//   - PixelTypeFloat64
//
// Returns:
//   - *Image: A pointer to the created Image.
//   - error: An error if the image creation or pixel conversion fails.
//
// If the current pixel type is the same as the specified pixel type, the function returns the current image without copying the pixel data.
func (img *Image) AsType(pixelType int) (*Image, error) {
	if img.pixelType == pixelType {
		return img, nil
	}

	newImg, err := NewImage(img.size, pixelType)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(img.size); i++ {
		newImg.spacing[i] = img.spacing[i]
		newImg.origin[i] = img.origin[i]
	}

	newImg.direction = img.direction

	numPixels := 1
	for _, s := range img.size {
		numPixels *= int(s)
	}

	numGoroutines := runtime.NumCPU()
	chunkSize := numPixels / numGoroutines
	if chunkSize*numGoroutines < numPixels {
		chunkSize += 1
	}

	switch pixelType {
	case PixelTypeUInt8:
		newPixelData := make([]byte, numPixels*1)
		switch img.pixelType {
		case PixelTypeUInt8:
			newPixelData = img.pixels
		case PixelTypeInt8:
			for i := 0; i < numPixels; i++ {
				newPixelData[i] = byte(uint8(img.pixels[i]))
			}
		case PixelTypeUInt16:
			for i := 0; i < numPixels; i++ {
				newPixelData[i] = byte(uint8(binary.LittleEndian.Uint16(img.pixels[i*2 : i*2+2])))
			}
		case PixelTypeInt16:
			for i := 0; i < numPixels; i++ {
				newPixelData[i] = byte(uint8(binary.LittleEndian.Uint16(img.pixels[i*2 : i*2+2])))
			}
		case PixelTypeUInt32:
			for i := 0; i < numPixels; i++ {
				newPixelData[i] = byte(uint8(binary.LittleEndian.Uint32(img.pixels[i*4 : i*4+4])))
			}
		case PixelTypeInt32:
			for i := 0; i < numPixels; i++ {
				newPixelData[i] = byte(uint8(binary.LittleEndian.Uint32(img.pixels[i*4 : i*4+4])))
			}
		case PixelTypeUInt64:
			for i := 0; i < numPixels; i++ {
				newPixelData[i] = byte(uint8(binary.LittleEndian.Uint64(img.pixels[i*8 : i*8+8])))
			}
		case PixelTypeInt64:
			for i := 0; i < numPixels; i++ {
				newPixelData[i] = byte(uint8(binary.LittleEndian.Uint64(img.pixels[i*8 : i*8+8])))
			}
		case PixelTypeFloat32:
			for i := 0; i < numPixels; i++ {
				newPixelData[i] = byte(uint8(math.Float32frombits(binary.LittleEndian.Uint32(img.pixels[i*4 : i*4+4]))))
			}
		case PixelTypeFloat64:
			for i := 0; i < numPixels; i++ {
				newPixelData[i] = byte(uint8(math.Float64frombits(binary.LittleEndian.Uint64(img.pixels[i*8 : i*8+8]))))
			}
		}
		newImg.pixels = newPixelData
	case PixelTypeInt8:
		newPixelData := make([]byte, numPixels*1)
		switch img.pixelType {
		case PixelTypeUInt8:
			for i := 0; i < numPixels*1; i++ {
				newPixelData[i] = byte(int8(img.pixels[i]))
			}
		case PixelTypeInt8:
			newPixelData = img.pixels
		case PixelTypeUInt16:
			for i := 0; i < numPixels; i++ {
				newPixelData[i] = byte(int8(binary.LittleEndian.Uint16(img.pixels[i*2 : i*2+2])))
			}
		case PixelTypeInt16:
			for i := 0; i < numPixels; i++ {
				newPixelData[i] = byte(int8(binary.LittleEndian.Uint16(img.pixels[i*2 : i*2+2])))
			}
		case PixelTypeUInt32:
			for i := 0; i < numPixels; i++ {
				newPixelData[i] = byte(int8(binary.LittleEndian.Uint32(img.pixels[i*4 : i*4+4])))
			}
		case PixelTypeInt32:
			for i := 0; i < numPixels; i++ {
				newPixelData[i] = byte(int8(binary.LittleEndian.Uint32(img.pixels[i*4 : i*4+4])))
			}
		case PixelTypeUInt64:
			for i := 0; i < numPixels; i++ {
				newPixelData[i] = byte(int8(binary.LittleEndian.Uint64(img.pixels[i*8 : i*8+8])))
			}
		case PixelTypeInt64:
			for i := 0; i < numPixels; i++ {
				newPixelData[i] = byte(int8(binary.LittleEndian.Uint64(img.pixels[i*8 : i*8+8])))
			}
		case PixelTypeFloat32:
			for i := 0; i < numPixels; i++ {
				newPixelData[i] = byte(int8(math.Float32frombits(binary.LittleEndian.Uint32(img.pixels[i*4 : i*4+4]))))
			}
		case PixelTypeFloat64:
			for i := 0; i < numPixels; i++ {
				newPixelData[i] = byte(int8(math.Float64frombits(binary.LittleEndian.Uint64(img.pixels[i*8 : i*8+8]))))
			}
		}
		newImg.pixels = newPixelData
	case PixelTypeUInt16, PixelTypeInt16:
		newPixelData := make([]byte, numPixels*2)
		switch img.pixelType {
		case PixelTypeUInt8:
			for i := 0; i < numPixels; i++ {
				binary.LittleEndian.PutUint16(newPixelData[i*2:i*2+2], uint16(img.pixels[i]))
			}
		case PixelTypeInt8:
			for i := 0; i < numPixels; i++ {
				binary.LittleEndian.PutUint16(newPixelData[i*2:i*2+2], uint16(img.pixels[i]))
			}
		case PixelTypeUInt16:
			newPixelData = img.pixels
		case PixelTypeInt16:
			for i := 0; i < numPixels; i++ {
				binary.LittleEndian.PutUint16(newPixelData[i*2:i*2+2], uint16(binary.LittleEndian.Uint16(img.pixels[i*2:i*2+2])))
			}
		case PixelTypeUInt32:
			for i := 0; i < numPixels; i++ {
				binary.LittleEndian.PutUint16(newPixelData[i*2:i*2+2], uint16(binary.LittleEndian.Uint32(img.pixels[i*4:i*4+4])))
			}
		case PixelTypeInt32:
			for i := 0; i < numPixels; i++ {
				binary.LittleEndian.PutUint16(newPixelData[i*2:i*2+2], uint16(binary.LittleEndian.Uint32(img.pixels[i*4:i*4+4])))
			}
		case PixelTypeUInt64:
			for i := 0; i < numPixels; i++ {
				binary.LittleEndian.PutUint16(newPixelData[i*2:i*2+2], uint16(binary.LittleEndian.Uint64(img.pixels[i*8:i*8+8])))
			}
		case PixelTypeInt64:
			for i := 0; i < numPixels; i++ {
				binary.LittleEndian.PutUint16(newPixelData[i*2:i*2+2], uint16(binary.LittleEndian.Uint64(img.pixels[i*8:i*8+8])))
			}
		case PixelTypeFloat32:
			for i := 0; i < numPixels; i++ {
				binary.LittleEndian.PutUint16(newPixelData[i*2:i*2+2], uint16(math.Float32frombits(binary.LittleEndian.Uint32(img.pixels[i*4:i*4+4]))))
			}
		case PixelTypeFloat64:
			for i := 0; i < numPixels; i++ {
				binary.LittleEndian.PutUint16(newPixelData[i*2:i*2+2], uint16(math.Float64frombits(binary.LittleEndian.Uint64(img.pixels[i*8:i*8+8]))))
			}
		}
		newImg.pixels = newPixelData
	case PixelTypeUInt32, PixelTypeInt32:
		newPixelData := make([]byte, numPixels*4)
		switch img.pixelType {
		case PixelTypeUInt8:
			for i := 0; i < numPixels; i++ {
				binary.LittleEndian.PutUint32(newPixelData[i*4:i*4+4], uint32(img.pixels[i]))
			}
		case PixelTypeInt8:
			for i := 0; i < numPixels; i++ {
				binary.LittleEndian.PutUint32(newPixelData[i*4:i*4+4], uint32(img.pixels[i]))
			}
		case PixelTypeUInt16:
			for i := 0; i < numPixels; i++ {
				binary.LittleEndian.PutUint32(newPixelData[i*4:i*4+4], uint32(binary.LittleEndian.Uint16(img.pixels[i*2:i*2+2])))
			}
		case PixelTypeInt16:
			for i := 0; i < numPixels; i++ {
				binary.LittleEndian.PutUint32(newPixelData[i*4:i*4+4], uint32(binary.LittleEndian.Uint16(img.pixels[i*2:i*2+2])))
			}
		case PixelTypeUInt32:
			newPixelData = img.pixels
		case PixelTypeInt32:
			for i := 0; i < numPixels; i++ {
				binary.LittleEndian.PutUint32(newPixelData[i*4:i*4+4], uint32(binary.LittleEndian.Uint32(img.pixels[i*4:i*4+4])))
			}
		case PixelTypeUInt64:
			for i := 0; i < numPixels; i++ {
				binary.LittleEndian.PutUint32(newPixelData[i*4:i*4+4], uint32(binary.LittleEndian.Uint64(img.pixels[i*8:i*8+8])))
			}
		case PixelTypeInt64:
			for i := 0; i < numPixels; i++ {
				binary.LittleEndian.PutUint32(newPixelData[i*4:i*4+4], uint32(binary.LittleEndian.Uint64(img.pixels[i*8:i*8+8])))
			}
		case PixelTypeFloat32:
			for i := 0; i < numPixels; i++ {
				binary.LittleEndian.PutUint32(newPixelData[i*4:i*4+4], uint32(math.Float32frombits(binary.LittleEndian.Uint32(img.pixels[i*4:i*4+4]))))
			}
		case PixelTypeFloat64:
			for i := 0; i < numPixels; i++ {
				binary.LittleEndian.PutUint32(newPixelData[i*4:i*4+4], uint32(math.Float64frombits(binary.LittleEndian.Uint64(img.pixels[i*8:i*8+8]))))
			}
		}
		newImg.pixels = newPixelData
	case PixelTypeUInt64, PixelTypeInt64:
		newPixelData := make([]byte, numPixels*8)
		switch img.pixelType {
		case PixelTypeUInt8:
			for i := 0; i < numPixels; i++ {
				binary.LittleEndian.PutUint64(newPixelData[i*8:i*8+8], uint64(img.pixels[i]))
			}
		case PixelTypeInt8:
			for i := 0; i < numPixels; i++ {
				binary.LittleEndian.PutUint64(newPixelData[i*8:i*8+8], uint64(img.pixels[i]))
			}
		case PixelTypeUInt16:
			for i := 0; i < numPixels; i++ {
				binary.LittleEndian.PutUint64(newPixelData[i*8:i*8+8], uint64(binary.LittleEndian.Uint16(img.pixels[i*2:i*2+2])))
			}
		case PixelTypeInt16:
			for i := 0; i < numPixels; i++ {
				binary.LittleEndian.PutUint64(newPixelData[i*8:i*8+8], uint64(binary.LittleEndian.Uint16(img.pixels[i*2:i*2+2])))
			}
		case PixelTypeUInt32:
			for i := 0; i < numPixels; i++ {
				binary.LittleEndian.PutUint64(newPixelData[i*8:i*8+8], uint64(binary.LittleEndian.Uint32(img.pixels[i*4:i*4+4])))
			}
		case PixelTypeInt32:
			for i := 0; i < numPixels; i++ {
				binary.LittleEndian.PutUint64(newPixelData[i*8:i*8+8], uint64(binary.LittleEndian.Uint32(img.pixels[i*4:i*4+4])))
			}
		case PixelTypeUInt64:
			newPixelData = img.pixels
		case PixelTypeInt64:
			for i := 0; i < numPixels; i++ {
				binary.LittleEndian.PutUint64(newPixelData[i*8:i*8+8], uint64(binary.LittleEndian.Uint64(img.pixels[i*8:i*8+8])))
			}
		case PixelTypeFloat32:
			for i := 0; i < numPixels; i++ {
				binary.LittleEndian.PutUint64(newPixelData[i*8:i*8+8], uint64(math.Float32frombits(binary.LittleEndian.Uint32(img.pixels[i*4:i*4+4]))))
			}
		case PixelTypeFloat64:
			for i := 0; i < numPixels; i++ {
				binary.LittleEndian.PutUint64(newPixelData[i*8:i*8+8], uint64(math.Float64frombits(binary.LittleEndian.Uint64(img.pixels[i*8:i*8+8]))))
			}
		}
		newImg.pixels = newPixelData
	case PixelTypeFloat32:
		newPixelData := make([]byte, numPixels*4)
		switch img.pixelType {
		case PixelTypeUInt8:
			for i := 0; i < numPixels; i++ {
				binary.LittleEndian.PutUint32(newPixelData[i*4:i*4+4], math.Float32bits(float32(img.pixels[i])))
			}
		case PixelTypeInt8:
			for i := 0; i < numPixels; i++ {
				binary.LittleEndian.PutUint32(newPixelData[i*4:i*4+4], math.Float32bits(float32(img.pixels[i])))
			}
		case PixelTypeUInt16:
			for i := 0; i < numPixels; i++ {
				binary.LittleEndian.PutUint32(newPixelData[i*4:i*4+4], math.Float32bits(float32(binary.LittleEndian.Uint16(img.pixels[i*2:i*2+2]))))
			}
		case PixelTypeInt16:
			for i := 0; i < numPixels; i++ {
				binary.LittleEndian.PutUint32(newPixelData[i*4:i*4+4], math.Float32bits(float32(binary.LittleEndian.Uint16(img.pixels[i*2:i*2+2]))))
			}
		case PixelTypeUInt32:
			for i := 0; i < numPixels; i++ {
				binary.LittleEndian.PutUint32(newPixelData[i*4:i*4+4], math.Float32bits(float32(binary.LittleEndian.Uint32(img.pixels[i*4:i*4+4]))))
			}
		case PixelTypeInt32:
			for i := 0; i < numPixels; i++ {
				binary.LittleEndian.PutUint32(newPixelData[i*4:i*4+4], math.Float32bits(float32(binary.LittleEndian.Uint32(img.pixels[i*4:i*4+4]))))
			}
		case PixelTypeUInt64:
			for i := 0; i < numPixels; i++ {
				binary.LittleEndian.PutUint32(newPixelData[i*4:i*4+4], math.Float32bits(float32(binary.LittleEndian.Uint64(img.pixels[i*8:i*8+8]))))
			}
		case PixelTypeInt64:
			for i := 0; i < numPixels; i++ {
				binary.LittleEndian.PutUint32(newPixelData[i*4:i*4+4], math.Float32bits(float32(binary.LittleEndian.Uint64(img.pixels[i*8:i*8+8]))))
			}
		case PixelTypeFloat32:
			newPixelData = img.pixels
		case PixelTypeFloat64:
			for i := 0; i < numPixels; i++ {
				binary.LittleEndian.PutUint32(newPixelData[i*4:i*4+4], math.Float32bits(float32(binary.LittleEndian.Uint64(img.pixels[i*8:i*8+8]))))
			}
		}
		newImg.pixels = newPixelData
	case PixelTypeFloat64:
		newPixelData := make([]byte, numPixels*8)
		switch img.pixelType {
		case PixelTypeUInt8:
			for i := 0; i < numPixels; i++ {
				binary.LittleEndian.PutUint64(newPixelData[i*8:i*8+8], math.Float64bits(float64(img.pixels[i])))
			}
		case PixelTypeInt8:
			for i := 0; i < numPixels; i++ {
				binary.LittleEndian.PutUint64(newPixelData[i*8:i*8+8], math.Float64bits(float64(img.pixels[i])))
			}
		case PixelTypeUInt16:
			for i := 0; i < numPixels; i++ {
				binary.LittleEndian.PutUint64(newPixelData[i*8:i*8+8], math.Float64bits(float64(binary.LittleEndian.Uint16(img.pixels[i*2:i*2+2]))))
			}
		case PixelTypeInt16:
			for i := 0; i < numPixels; i++ {
				binary.LittleEndian.PutUint64(newPixelData[i*8:i*8+8], math.Float64bits(float64(binary.LittleEndian.Uint16(img.pixels[i*2:i*2+2]))))
			}
		case PixelTypeUInt32:
			for i := 0; i < numPixels; i++ {
				binary.LittleEndian.PutUint64(newPixelData[i*8:i*8+8], math.Float64bits(float64(binary.LittleEndian.Uint32(img.pixels[i*4:i*4+4]))))
			}
		case PixelTypeInt32:
			for i := 0; i < numPixels; i++ {
				binary.LittleEndian.PutUint64(newPixelData[i*8:i*8+8], math.Float64bits(float64(binary.LittleEndian.Uint32(img.pixels[i*4:i*4+4]))))
			}
		case PixelTypeUInt64:
			for i := 0; i < numPixels; i++ {
				binary.LittleEndian.PutUint64(newPixelData[i*8:i*8+8], math.Float64bits(float64(binary.LittleEndian.Uint64(img.pixels[i*8:i*8+8]))))
			}
		case PixelTypeInt64:
			for i := 0; i < numPixels; i++ {
				binary.LittleEndian.PutUint64(newPixelData[i*8:i*8+8], math.Float64bits(float64(binary.LittleEndian.Uint64(img.pixels[i*8:i*8+8]))))
			}
		case PixelTypeFloat32:
			for i := 0; i < numPixels; i++ {
				binary.LittleEndian.PutUint64(newPixelData[i*8:i*8+8], math.Float64bits(float64(math.Float32frombits(binary.LittleEndian.Uint32(img.pixels[i*4:i*4+4])))))
			}
		case PixelTypeFloat64:
			newPixelData = img.pixels
		}
		newImg.pixels = newPixelData
	default:
		return nil, fmt.Errorf("unsupported pixel type: %d", pixelType)
	}

	return newImg, nil
}

func (img *Image) SetSpacing(spacing []float64) error {
	if len(spacing) != int(img.dimension) {
		return fmt.Errorf("invalid spacing length: %d", len(spacing))
	}
	for _, s := range spacing {
		if s <= 0 {
			return fmt.Errorf("invalid spacing: %f", s)
		}
	}
	img.spacing = spacing
	return nil
}

func (img *Image) SetOrigin(origin []float64) error {
	if len(origin) != int(img.dimension) {
		return fmt.Errorf("invalid origin length: %d", len(origin))
	}
	img.origin = origin
	return nil
}

func (img *Image) SetDirection(direction [9]float64) {
	img.direction = direction
}

func (img *Image) SetPixel(index []uint32, value any) error {
	if len(index) != int(img.dimension) {
		return fmt.Errorf("invalid index length: %d", len(index))
	}
	idx := uint32(0)
	for i := len(index) - 1; i >= 0; i-- {
		if index[i] >= img.size[i] {
			return fmt.Errorf("index out of range: %d", index[i])
		}
		idx = idx*img.size[i] + index[i]
	}
	bytesPerPixel := uint32(img.bytesPerPixel)

	valueBytes, err := getValueAsBytes(value)
	if err != nil {
		return err
	}
	copy(img.pixels[idx*bytesPerPixel:idx*bytesPerPixel+bytesPerPixel], valueBytes)

	return nil
}

func (img *Image) SetPixels(pixels any) error {
	flattened, err := flattenToBytes(pixels)
	if err != nil {
		return err
	}
	numPixels := 1
	for _, s := range img.size {
		numPixels *= int(s)
	}

	if len(flattened)/img.bytesPerPixel != numPixels {
		return fmt.Errorf("invalid number of pixels, expected %d, got %d", numPixels, len(flattened))
	}

	img.pixels = flattened

	return nil
}

func (img *Image) SetSize(size []uint32) error {
	totalSize := uint32(1)
	for _, s := range size {
		totalSize *= s
	}

	numPixels := len(img.pixels) / img.bytesPerPixel

	if uint32(numPixels) != totalSize {
		return fmt.Errorf("invalid number of pixels, expected %d, got %d", totalSize, numPixels)
	}
	img.size = size
	img.dimension = uint32(len(size))
	return nil
}

// GetIndexFromLinearIndex converts a linear index to multi-dimensional indices
func (img *Image) GetIndexFromLinearIndex(linearIdx uint64) ([]uint32, error) {
	if linearIdx >= img.NumPixels() {
		return nil, fmt.Errorf("linear index out of range: %d", linearIdx)
	}

	indices := make([]uint32, img.dimension)
	remaining := linearIdx

	// Working backwards through dimensions
	for i := 0; i < int(img.dimension); i++ {
		// Calculate the size of the sub-array for the current dimension
		subArraySize := uint64(1)
		for j := 0; j < i; j++ {
			subArraySize *= uint64(img.size[j])
		}

		// Calculate the index for this dimension
		indices[i] = uint32((remaining / subArraySize) % uint64(img.size[i]))
		remaining = remaining % subArraySize
	}

	// Reverse the indices array since the original indexing was in reverse order
	for i := 0; i < len(indices)/2; i++ {
		indices[i], indices[len(indices)-1-i] = indices[len(indices)-1-i], indices[i]
	}

	return indices, nil
}

// GetTotalPixels returns the total number of pixels in the image
func (img *Image) NumPixels() uint64 {
	total := uint64(1)
	for _, size := range img.size {
		total *= uint64(size)
	}
	return total
}
