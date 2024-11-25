package imagetk

import (
	"fmt"
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
// - pixels: The pixel data of the image, type can vary.
// - pixelType: An integer representing the type of pixels.
// - dimension: The number of dimensions of the image (2<=N<=3).
// - size: A slice of uint32 representing the size of the image in each dimension (2<=N<=3).
// - spacing: A slice of float64 representing the spacing between pixels in each dimension.
// - origin: A slice of float64 representing the origin of the image.
// - direction: An array of 9 float64 values representing the direction cosines of the image.
type Image struct {
	pixels    any
	pixelType int
	dimension uint32
	size      []uint32
	spacing   []float64
	origin    []float64
	direction [9]float64
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
	var pixels any
	switch pixelType {
	case PixelTypeUInt8:
		pixels = make([]uint8, numPixels)
	case PixelTypeInt8:
		pixels = make([]int8, numPixels)
	case PixelTypeUInt16:
		pixels = make([]uint16, numPixels)
	case PixelTypeInt16:
		pixels = make([]int16, numPixels)
	case PixelTypeUInt32:
		pixels = make([]uint32, numPixels)
	case PixelTypeInt32:
		pixels = make([]int32, numPixels)
	case PixelTypeUInt64:
		pixels = make([]uint64, numPixels)
	case PixelTypeInt64:
		pixels = make([]int64, numPixels)
	case PixelTypeFloat32:
		pixels = make([]float32, numPixels)
	case PixelTypeFloat64:
		pixels = make([]float64, numPixels)
	default:
		return nil, fmt.Errorf("unsupported pixel type: %d", pixelType)
	}

	spacing := make([]float64, len(size))
	origin := make([]float64, len(size))
	for i := 0; i < len(size); i++ {
		origin[i] = 0
		spacing[i] = 1
	}

	direction := [9]float64{1, 0, 0, 0, 1, 0, 0, 0, 1}

	return &Image{
		pixels:    pixels,
		pixelType: pixelType,
		dimension: uint32(len(size)),
		size:      size,
		spacing:   spacing,
		origin:    origin,
		direction: direction,
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
		return img.pixels.([]uint8)[idx], nil
	case PixelTypeInt8:
		return img.pixels.([]int8)[idx], nil
	case PixelTypeUInt16:
		return img.pixels.([]uint16)[idx], nil
	case PixelTypeInt16:
		return img.pixels.([]int16)[idx], nil
	case PixelTypeUInt32:
		return img.pixels.([]uint32)[idx], nil
	case PixelTypeInt32:
		return img.pixels.([]int32)[idx], nil
	case PixelTypeUInt64:
		return img.pixels.([]uint64)[idx], nil
	case PixelTypeInt64:
		return img.pixels.([]int64)[idx], nil
	case PixelTypeFloat32:
		return img.pixels.([]float32)[idx], nil
	case PixelTypeFloat64:
		return img.pixels.([]float64)[idx], nil
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
	idx := uint32(0)
	for i := len(index) - 1; i >= 0; i-- {
		if index[i] >= img.size[i] {
			return 0, fmt.Errorf("index out of range: %d", index[i])
		}
		idx = idx*img.size[i] + index[i]
	}
	switch img.pixelType {
	case PixelTypeUInt8:
		return img.pixels.([]uint8)[idx], nil
	case PixelTypeInt8:
		return uint8(img.pixels.([]int8)[idx]), nil
	case PixelTypeUInt16:
		return uint8(img.pixels.([]uint16)[idx]), nil
	case PixelTypeInt16:
		return uint8(img.pixels.([]int16)[idx]), nil
	case PixelTypeUInt32:
		return uint8(img.pixels.([]uint32)[idx]), nil
	case PixelTypeInt32:
		return uint8(img.pixels.([]int32)[idx]), nil
	case PixelTypeUInt64:
		return uint8(img.pixels.([]uint64)[idx]), nil
	case PixelTypeInt64:
		return uint8(img.pixels.([]int64)[idx]), nil
	case PixelTypeFloat32:
		return uint8(img.pixels.([]float32)[idx]), nil
	case PixelTypeFloat64:
		return uint8(img.pixels.([]float64)[idx]), nil
	default:
		return 0, fmt.Errorf("unsupported pixel type: %d", img.pixelType)
	}
}

// GetPixelAsInt8 returns the pixel value at the given index as a int8.
// Parameters:
//   - index: A slice of uint32 representing the index of the pixel.
//
// Returns:
//   - int8: The pixel value as a int8.
//   - error: An error if the index is out of range.
func (img *Image) GetPixelAsInt8(index []uint32) (int8, error) {
	idx := uint32(0)
	for i := len(index) - 1; i >= 0; i-- {
		if index[i] >= img.size[i] {
			return 0, fmt.Errorf("index out of range: %d", index[i])
		}
		idx = idx*img.size[i] + index[i]
	}
	switch img.pixelType {
	case PixelTypeUInt8:
		return int8(img.pixels.([]uint8)[idx]), nil
	case PixelTypeInt8:
		return img.pixels.([]int8)[idx], nil
	case PixelTypeUInt16:
		return int8(img.pixels.([]uint16)[idx]), nil
	case PixelTypeInt16:
		return int8(img.pixels.([]int16)[idx]), nil
	case PixelTypeUInt32:
		return int8(img.pixels.([]uint32)[idx]), nil
	case PixelTypeInt32:
		return int8(img.pixels.([]int32)[idx]), nil
	case PixelTypeUInt64:
		return int8(img.pixels.([]uint64)[idx]), nil
	case PixelTypeInt64:
		return int8(img.pixels.([]int64)[idx]), nil
	case PixelTypeFloat32:
		return int8(img.pixels.([]float32)[idx]), nil
	case PixelTypeFloat64:
		return int8(img.pixels.([]float64)[idx]), nil
	default:
		return 0, fmt.Errorf("unsupported pixel type: %d", img.pixelType)
	}
}

// GetPixelAsUInt16 returns the pixel value at the given index as a uint16.
// Parameters:
//   - index: A slice of uint32 representing the index of the pixel.
//
// Returns:
//   - uint16: The pixel value as a uint16.
//   - error: An error if the index is out of range.
func (img *Image) GetPixelAsUInt16(index []uint32) (uint16, error) {
	idx := uint32(0)
	for i := len(index) - 1; i >= 0; i-- {
		if index[i] >= img.size[i] {
			return 0, fmt.Errorf("index out of range: %d", index[i])
		}
		idx = idx*img.size[i] + index[i]
	}
	switch img.pixelType {
	case PixelTypeUInt8:
		return uint16(img.pixels.([]uint8)[idx]), nil
	case PixelTypeInt8:
		return uint16(img.pixels.([]int8)[idx]), nil
	case PixelTypeUInt16:
		return img.pixels.([]uint16)[idx], nil
	case PixelTypeInt16:
		return uint16(img.pixels.([]int16)[idx]), nil
	case PixelTypeUInt32:
		return uint16(img.pixels.([]uint32)[idx]), nil
	case PixelTypeInt32:
		return uint16(img.pixels.([]int32)[idx]), nil
	case PixelTypeUInt64:
		return uint16(img.pixels.([]uint64)[idx]), nil
	case PixelTypeInt64:
		return uint16(img.pixels.([]int64)[idx]), nil
	case PixelTypeFloat32:
		return uint16(img.pixels.([]float32)[idx]), nil
	case PixelTypeFloat64:
		return uint16(img.pixels.([]float64)[idx]), nil
	default:
		return 0, fmt.Errorf("unsupported pixel type: %d", img.pixelType)
	}
}

// GetPixelAsInt16 returns the pixel value at the given index as a int16.
// Parameters:
//   - index: A slice of uint32 representing the index of the pixel.
//
// Returns:
//   - int16: The pixel value as a int16.
//   - error: An error if the index is out of range.
func (img *Image) GetPixelAsInt16(index []uint32) (int16, error) {
	idx := uint32(0)
	for i := len(index) - 1; i >= 0; i-- {
		if index[i] >= img.size[i] {
			return 0, fmt.Errorf("index out of range: %d", index[i])
		}
		idx = idx*img.size[i] + index[i]
	}
	switch img.pixelType {
	case PixelTypeUInt8:
		return int16(img.pixels.([]uint8)[idx]), nil
	case PixelTypeInt8:
		return int16(img.pixels.([]int8)[idx]), nil
	case PixelTypeUInt16:
		return int16(img.pixels.([]uint16)[idx]), nil
	case PixelTypeInt16:
		return img.pixels.([]int16)[idx], nil
	case PixelTypeUInt32:
		return int16(img.pixels.([]uint32)[idx]), nil
	case PixelTypeInt32:
		return int16(img.pixels.([]int32)[idx]), nil
	case PixelTypeUInt64:
		return int16(img.pixels.([]uint64)[idx]), nil
	case PixelTypeInt64:
		return int16(img.pixels.([]int64)[idx]), nil
	case PixelTypeFloat32:
		return int16(img.pixels.([]float32)[idx]), nil
	case PixelTypeFloat64:
		return int16(img.pixels.([]float64)[idx]), nil
	default:
		return 0, fmt.Errorf("unsupported pixel type: %d", img.pixelType)
	}
}

// GetPixelAsUInt32 returns the pixel value at the given index as a uint32.
// Parameters:
//   - index: A slice of uint32 representing the index of the pixel.
//
// Returns:
//   - uint32: The pixel value as a uint32.
//   - error: An error if the index is out of range.
func (img *Image) GetPixelAsUInt32(index []uint32) (uint32, error) {
	idx := uint32(0)
	for i := len(index) - 1; i >= 0; i-- {
		if index[i] >= img.size[i] {
			return 0, fmt.Errorf("index out of range: %d", index[i])
		}
		idx = idx*img.size[i] + index[i]
	}
	switch img.pixelType {
	case PixelTypeUInt8:
		return uint32(img.pixels.([]uint8)[idx]), nil
	case PixelTypeInt8:
		return uint32(img.pixels.([]int8)[idx]), nil
	case PixelTypeUInt16:
		return uint32(img.pixels.([]uint16)[idx]), nil
	case PixelTypeInt16:
		return uint32(img.pixels.([]int16)[idx]), nil
	case PixelTypeUInt32:
		return img.pixels.([]uint32)[idx], nil
	case PixelTypeInt32:
		return uint32(img.pixels.([]int32)[idx]), nil
	case PixelTypeUInt64:
		return uint32(img.pixels.([]uint64)[idx]), nil
	case PixelTypeInt64:
		return uint32(img.pixels.([]int64)[idx]), nil
	case PixelTypeFloat32:
		return uint32(img.pixels.([]float32)[idx]), nil
	case PixelTypeFloat64:
		return uint32(img.pixels.([]float64)[idx]), nil
	default:
		return 0, fmt.Errorf("unsupported pixel type: %d", img.pixelType)
	}
}

// GetPixelAsInt32 returns the pixel value at the given index as a int32.
// Parameters:
//   - index: A slice of uint32 representing the index of the pixel.
//
// Returns:
//   - int32: The pixel value as a int32.
//   - error: An error if the index is out of range.
func (img *Image) GetPixelAsInt32(index []uint32) (int32, error) {
	idx := uint32(0)
	for i := len(index) - 1; i >= 0; i-- {
		if index[i] >= img.size[i] {
			return 0, fmt.Errorf("index out of range: %d", index[i])
		}
		idx = idx*img.size[i] + index[i]
	}
	switch img.pixelType {
	case PixelTypeUInt8:
		return int32(img.pixels.([]uint8)[idx]), nil
	case PixelTypeInt8:
		return int32(img.pixels.([]int8)[idx]), nil
	case PixelTypeUInt16:
		return int32(img.pixels.([]uint16)[idx]), nil
	case PixelTypeInt16:
		return int32(img.pixels.([]int16)[idx]), nil
	case PixelTypeUInt32:
		return int32(img.pixels.([]uint32)[idx]), nil
	case PixelTypeInt32:
		return img.pixels.([]int32)[idx], nil
	case PixelTypeUInt64:
		return int32(img.pixels.([]uint64)[idx]), nil
	case PixelTypeInt64:
		return int32(img.pixels.([]int64)[idx]), nil
	case PixelTypeFloat32:
		return int32(img.pixels.([]float32)[idx]), nil
	case PixelTypeFloat64:
		return int32(img.pixels.([]float64)[idx]), nil
	default:
		return 0, fmt.Errorf("unsupported pixel type: %d", img.pixelType)
	}
}

// GetPixelAsUInt64 returns the pixel value at the given index as a uint64.
// Parameters:
//   - index: A slice of uint32 representing the index of the pixel.
//
// Returns:
//   - uint64: The pixel value as a uint64.
//   - error: An error if the index is out of range.
func (img *Image) GetPixelAsUInt64(index []uint32) (uint64, error) {
	idx := uint32(0)
	for i := len(index) - 1; i >= 0; i-- {
		if index[i] >= img.size[i] {
			return 0, fmt.Errorf("index out of range: %d", index[i])
		}
		idx = idx*img.size[i] + index[i]
	}
	switch img.pixelType {
	case PixelTypeUInt8:
		return uint64(img.pixels.([]uint8)[idx]), nil
	case PixelTypeInt8:
		return uint64(img.pixels.([]int8)[idx]), nil
	case PixelTypeUInt16:
		return uint64(img.pixels.([]uint16)[idx]), nil
	case PixelTypeInt16:
		return uint64(img.pixels.([]int16)[idx]), nil
	case PixelTypeUInt32:
		return uint64(img.pixels.([]uint32)[idx]), nil
	case PixelTypeInt32:
		return uint64(img.pixels.([]int32)[idx]), nil
	case PixelTypeUInt64:
		return img.pixels.([]uint64)[idx], nil
	case PixelTypeInt64:
		return uint64(img.pixels.([]int64)[idx]), nil
	case PixelTypeFloat32:
		return uint64(img.pixels.([]float32)[idx]), nil
	case PixelTypeFloat64:
		return uint64(img.pixels.([]float64)[idx]), nil
	default:
		return 0, fmt.Errorf("unsupported pixel type: %d", img.pixelType)
	}
}

// GetPixelAsInt64 returns the pixel value at the given index as a int64.
// Parameters:
//   - index: A slice of uint32 representing the index of the pixel.
//
// Returns:
//   - int64: The pixel value as a int64.
//   - error: An error if the index is out of range.
func (img *Image) GetPixelAsInt64(index []uint32) (int64, error) {
	idx := uint32(0)
	for i := len(index) - 1; i >= 0; i-- {
		if index[i] >= img.size[i] {
			return 0, fmt.Errorf("index out of range: %d", index[i])
		}
		idx = idx*img.size[i] + index[i]
	}
	switch img.pixelType {
	case PixelTypeUInt8:
		return int64(img.pixels.([]uint8)[idx]), nil
	case PixelTypeInt8:
		return int64(img.pixels.([]int8)[idx]), nil
	case PixelTypeUInt16:
		return int64(img.pixels.([]uint16)[idx]), nil
	case PixelTypeInt16:
		return int64(img.pixels.([]int16)[idx]), nil
	case PixelTypeUInt32:
		return int64(img.pixels.([]uint32)[idx]), nil
	case PixelTypeInt32:
		return int64(img.pixels.([]int32)[idx]), nil
	case PixelTypeUInt64:
		return int64(img.pixels.([]uint64)[idx]), nil
	case PixelTypeInt64:
		return img.pixels.([]int64)[idx], nil
	case PixelTypeFloat32:
		return int64(img.pixels.([]float32)[idx]), nil
	case PixelTypeFloat64:
		return int64(img.pixels.([]float64)[idx]), nil
	default:
		return 0, fmt.Errorf("unsupported pixel type: %d", img.pixelType)
	}
}

// GetPixelAsFloat32 returns the pixel value at the given index as a float32.
// Parameters:
//   - index: A slice of uint32 representing the index of the pixel.
//
// Returns:
//   - float32: The pixel value as a float32.
//   - error: An error if the index is out of range.
func (img *Image) GetPixelAsFloat32(index []uint32) (float32, error) {
	idx := uint32(0)
	for i := len(index) - 1; i >= 0; i-- {
		if index[i] >= img.size[i] {
			return 0, fmt.Errorf("index out of range: %d", index[i])
		}
		idx = idx*img.size[i] + index[i]
	}
	switch img.pixelType {
	case PixelTypeUInt8:
		return float32(img.pixels.([]uint8)[idx]), nil
	case PixelTypeInt8:
		return float32(img.pixels.([]int8)[idx]), nil
	case PixelTypeUInt16:
		return float32(img.pixels.([]uint16)[idx]), nil
	case PixelTypeInt16:
		return float32(img.pixels.([]int16)[idx]), nil
	case PixelTypeUInt32:
		return float32(img.pixels.([]uint32)[idx]), nil
	case PixelTypeInt32:
		return float32(img.pixels.([]int32)[idx]), nil
	case PixelTypeUInt64:
		return float32(img.pixels.([]uint64)[idx]), nil
	case PixelTypeInt64:
		return float32(img.pixels.([]int64)[idx]), nil
	case PixelTypeFloat32:
		return img.pixels.([]float32)[idx], nil
	case PixelTypeFloat64:
		return float32(img.pixels.([]float64)[idx]), nil
	default:
		return 0, fmt.Errorf("unsupported pixel type: %d", img.pixelType)
	}
}

// GetPixelAsFloat64 returns the pixel value at the given index as a float64.
// Parameters:
//   - index: A slice of uint32 representing the index of the pixel.
//
// Returns:
//   - float64: The pixel value as a float64.
//   - error: An error if the index is out of range.
func (img *Image) GetPixelAsFloat64(index []uint32) (float64, error) {
	idx := uint32(0)
	for i := len(index) - 1; i >= 0; i-- {
		if index[i] >= img.size[i] {
			return 0, fmt.Errorf("index out of range: %d", index[i])
		}
		idx = idx*img.size[i] + index[i]
	}
	switch img.pixelType {
	case PixelTypeUInt8:
		return float64(img.pixels.([]uint8)[idx]), nil
	case PixelTypeInt8:
		return float64(img.pixels.([]int8)[idx]), nil
	case PixelTypeUInt16:
		return float64(img.pixels.([]uint16)[idx]), nil
	case PixelTypeInt16:
		return float64(img.pixels.([]int16)[idx]), nil
	case PixelTypeUInt32:
		return float64(img.pixels.([]uint32)[idx]), nil
	case PixelTypeInt32:
		return float64(img.pixels.([]int32)[idx]), nil
	case PixelTypeUInt64:
		return float64(img.pixels.([]uint64)[idx]), nil
	case PixelTypeInt64:
		return float64(img.pixels.([]int64)[idx]), nil
	case PixelTypeFloat32:
		return float64(img.pixels.([]float32)[idx]), nil
	case PixelTypeFloat64:
		return img.pixels.([]float64)[idx], nil
	default:
		return -1, fmt.Errorf("unsupported pixel type: %d", img.pixelType)
	}
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
	wg := sync.WaitGroup{}
	switch pixelType {
	case PixelTypeUInt8:
		newPixelData := make([]uint8, numPixels)
		for chunk := 0; chunk < numGoroutines; chunk++ {
			start := chunk * chunkSize
			end := start + chunkSize
			if end > numPixels {
				end = numPixels
			}
			wg.Add(1)
			go func(start, end int) {
				defer wg.Done()
				for i := start; i < end; i++ {
					switch img.pixelType {
					case PixelTypeInt8:
						newPixelData[i] = uint8(img.pixels.([]int8)[i])
					case PixelTypeUInt16:
						newPixelData[i] = uint8(img.pixels.([]uint16)[i])
					case PixelTypeInt16:
						newPixelData[i] = uint8(img.pixels.([]int16)[i])
					case PixelTypeUInt32:
						newPixelData[i] = uint8(img.pixels.([]uint32)[i])
					case PixelTypeInt32:
						newPixelData[i] = uint8(img.pixels.([]int32)[i])
					case PixelTypeUInt64:
						newPixelData[i] = uint8(img.pixels.([]uint64)[i])
					case PixelTypeInt64:
						newPixelData[i] = uint8(img.pixels.([]int64)[i])
					case PixelTypeFloat32:
						newPixelData[i] = uint8(img.pixels.([]float32)[i])
					case PixelTypeFloat64:
						newPixelData[i] = uint8(img.pixels.([]float64)[i])
					}
				}
			}(start, end)
		}
		wg.Wait()
		newImg.pixels = newPixelData
	case PixelTypeInt8:
		newPixelData := make([]int8, numPixels)
		for chunk := 0; chunk < numGoroutines; chunk++ {
			start := chunk * chunkSize
			end := start + chunkSize
			if end > numPixels {
				end = numPixels
			}
			wg.Add(1)
			go func(start, end int) {
				defer wg.Done()
				for i := start; i < end; i++ {
					switch img.pixelType {
					case PixelTypeUInt8:
						newPixelData[i] = int8(img.pixels.([]uint8)[i])
					case PixelTypeUInt16:
						newPixelData[i] = int8(img.pixels.([]uint16)[i])
					case PixelTypeInt16:
						newPixelData[i] = int8(img.pixels.([]int16)[i])
					case PixelTypeUInt32:
						newPixelData[i] = int8(img.pixels.([]uint32)[i])
					case PixelTypeInt32:
						newPixelData[i] = int8(img.pixels.([]int32)[i])
					case PixelTypeUInt64:
						newPixelData[i] = int8(img.pixels.([]uint64)[i])
					case PixelTypeInt64:
						newPixelData[i] = int8(img.pixels.([]int64)[i])
					case PixelTypeFloat32:
						newPixelData[i] = int8(img.pixels.([]float32)[i])
					case PixelTypeFloat64:
						newPixelData[i] = int8(img.pixels.([]float64)[i])
					}
				}
			}(start, end)
		}
		wg.Wait()
		newImg.pixels = newPixelData
	case PixelTypeUInt16:
		newPixelData := make([]uint16, numPixels)
		for chunk := 0; chunk < numGoroutines; chunk++ {
			start := chunk * chunkSize
			end := start + chunkSize
			if end > numPixels {
				end = numPixels
			}
			wg.Add(1)
			go func(start, end int) {
				defer wg.Done()
				for i := start; i < end; i++ {
					switch img.pixelType {
					case PixelTypeUInt8:
						newPixelData[i] = uint16(img.pixels.([]uint8)[i])
					case PixelTypeInt8:
						newPixelData[i] = uint16(img.pixels.([]int8)[i])
					case PixelTypeInt16:
						newPixelData[i] = uint16(img.pixels.([]int16)[i])
					case PixelTypeUInt32:
						newPixelData[i] = uint16(img.pixels.([]uint32)[i])
					case PixelTypeInt32:
						newPixelData[i] = uint16(img.pixels.([]int32)[i])
					case PixelTypeUInt64:
						newPixelData[i] = uint16(img.pixels.([]uint64)[i])
					case PixelTypeInt64:
						newPixelData[i] = uint16(img.pixels.([]int64)[i])
					case PixelTypeFloat32:
						newPixelData[i] = uint16(img.pixels.([]float32)[i])
					case PixelTypeFloat64:
						newPixelData[i] = uint16(img.pixels.([]float64)[i])
					}
				}
			}(start, end)
		}
		wg.Wait()
		newImg.pixels = newPixelData
	case PixelTypeInt16:
		newPixelData := make([]int16, numPixels)
		for chunk := 0; chunk < numGoroutines; chunk++ {
			start := chunk * chunkSize
			end := start + chunkSize
			if end > numPixels {
				end = numPixels
			}
			wg.Add(1)
			go func(start, end int) {
				defer wg.Done()
				for i := start; i < end; i++ {
					switch img.pixelType {
					case PixelTypeUInt8:
						newPixelData[i] = int16(img.pixels.([]uint8)[i])
					case PixelTypeInt8:
						newPixelData[i] = int16(img.pixels.([]int8)[i])
					case PixelTypeUInt16:
						newPixelData[i] = int16(img.pixels.([]uint16)[i])
					case PixelTypeUInt32:
						newPixelData[i] = int16(img.pixels.([]uint32)[i])
					case PixelTypeInt32:
						newPixelData[i] = int16(img.pixels.([]int32)[i])
					case PixelTypeUInt64:
						newPixelData[i] = int16(img.pixels.([]uint64)[i])
					case PixelTypeInt64:
						newPixelData[i] = int16(img.pixels.([]int64)[i])
					case PixelTypeFloat32:
						newPixelData[i] = int16(img.pixels.([]float32)[i])
					case PixelTypeFloat64:
						newPixelData[i] = int16(img.pixels.([]float64)[i])
					}
				}
			}(start, end)
		}
		wg.Wait()
		newImg.pixels = newPixelData
	case PixelTypeUInt32:
		newPixelData := make([]uint32, numPixels)
		for chunk := 0; chunk < numGoroutines; chunk++ {
			start := chunk * chunkSize
			end := start + chunkSize
			if end > numPixels {
				end = numPixels
			}
			wg.Add(1)
			go func(start, end int) {
				defer wg.Done()
				for i := start; i < end; i++ {
					switch img.pixelType {
					case PixelTypeUInt8:
						newPixelData[i] = uint32(img.pixels.([]uint8)[i])
					case PixelTypeInt8:
						newPixelData[i] = uint32(img.pixels.([]int8)[i])
					case PixelTypeUInt16:
						newPixelData[i] = uint32(img.pixels.([]uint16)[i])
					case PixelTypeInt16:
						newPixelData[i] = uint32(img.pixels.([]int16)[i])
					case PixelTypeInt32:
						newPixelData[i] = uint32(img.pixels.([]int32)[i])
					case PixelTypeUInt64:
						newPixelData[i] = uint32(img.pixels.([]uint64)[i])
					case PixelTypeInt64:
						newPixelData[i] = uint32(img.pixels.([]int64)[i])
					case PixelTypeFloat32:
						newPixelData[i] = uint32(img.pixels.([]float32)[i])
					case PixelTypeFloat64:
						newPixelData[i] = uint32(img.pixels.([]float64)[i])
					}
				}
			}(start, end)
		}
		wg.Wait()
		newImg.pixels = newPixelData
	case PixelTypeInt32:
		newPixelData := make([]int32, numPixels)
		for chunk := 0; chunk < numGoroutines; chunk++ {
			start := chunk * chunkSize
			end := start + chunkSize
			if end > numPixels {
				end = numPixels
			}
			wg.Add(1)
			go func(start, end int) {
				defer wg.Done()
				for i := start; i < end; i++ {
					switch img.pixelType {
					case PixelTypeUInt8:
						newPixelData[i] = int32(img.pixels.([]uint8)[i])
					case PixelTypeInt8:
						newPixelData[i] = int32(img.pixels.([]int8)[i])
					case PixelTypeUInt16:
						newPixelData[i] = int32(img.pixels.([]uint16)[i])
					case PixelTypeInt16:
						newPixelData[i] = int32(img.pixels.([]int16)[i])
					case PixelTypeUInt32:
						newPixelData[i] = int32(img.pixels.([]uint32)[i])
					case PixelTypeUInt64:
						newPixelData[i] = int32(img.pixels.([]uint64)[i])
					case PixelTypeInt64:
						newPixelData[i] = int32(img.pixels.([]int64)[i])
					case PixelTypeFloat32:
						newPixelData[i] = int32(img.pixels.([]float32)[i])
					case PixelTypeFloat64:
						newPixelData[i] = int32(img.pixels.([]float64)[i])
					}
				}
			}(start, end)
		}
		wg.Wait()
		newImg.pixels = newPixelData
	case PixelTypeUInt64:
		newPixelData := make([]uint64, numPixels)
		for chunk := 0; chunk < numGoroutines; chunk++ {
			start := chunk * chunkSize
			end := start + chunkSize
			if end > numPixels {
				end = numPixels
			}
			wg.Add(1)
			go func(start, end int) {
				defer wg.Done()
				for i := start; i < end; i++ {
					switch img.pixelType {
					case PixelTypeUInt8:
						newPixelData[i] = uint64(img.pixels.([]uint8)[i])
					case PixelTypeInt8:
						newPixelData[i] = uint64(img.pixels.([]int8)[i])
					case PixelTypeUInt16:
						newPixelData[i] = uint64(img.pixels.([]uint16)[i])
					case PixelTypeInt16:
						newPixelData[i] = uint64(img.pixels.([]int16)[i])
					case PixelTypeUInt32:
						newPixelData[i] = uint64(img.pixels.([]uint32)[i])
					case PixelTypeInt32:
						newPixelData[i] = uint64(img.pixels.([]int32)[i])
					case PixelTypeInt64:
						newPixelData[i] = uint64(img.pixels.([]int64)[i])
					case PixelTypeFloat32:
						newPixelData[i] = uint64(img.pixels.([]float32)[i])
					case PixelTypeFloat64:
						newPixelData[i] = uint64(img.pixels.([]float64)[i])
					}
				}
			}(start, end)
		}
		wg.Wait()
		newImg.pixels = newPixelData
	case PixelTypeInt64:
		newPixelData := make([]int64, numPixels)
		for chunk := 0; chunk < numGoroutines; chunk++ {
			start := chunk * chunkSize
			end := start + chunkSize
			if end > numPixels {
				end = numPixels
			}
			wg.Add(1)
			go func(start, end int) {
				defer wg.Done()
				for i := start; i < end; i++ {
					switch img.pixelType {
					case PixelTypeUInt8:
						newPixelData[i] = int64(img.pixels.([]uint8)[i])
					case PixelTypeInt8:
						newPixelData[i] = int64(img.pixels.([]int8)[i])
					case PixelTypeUInt16:
						newPixelData[i] = int64(img.pixels.([]uint16)[i])
					case PixelTypeInt16:
						newPixelData[i] = int64(img.pixels.([]int16)[i])
					case PixelTypeUInt32:
						newPixelData[i] = int64(img.pixels.([]uint32)[i])
					case PixelTypeInt32:
						newPixelData[i] = int64(img.pixels.([]int32)[i])
					case PixelTypeUInt64:
						newPixelData[i] = int64(img.pixels.([]uint64)[i])
					case PixelTypeFloat32:
						newPixelData[i] = int64(img.pixels.([]float32)[i])
					case PixelTypeFloat64:
						newPixelData[i] = int64(img.pixels.([]float64)[i])
					}
				}
			}(start, end)
		}
		wg.Wait()
		newImg.pixels = newPixelData
	case PixelTypeFloat32:
		newPixelData := make([]float32, numPixels)
		for chunk := 0; chunk < numGoroutines; chunk++ {
			start := chunk * chunkSize
			end := start + chunkSize
			if end > numPixels {
				end = numPixels
			}
			wg.Add(1)
			go func(start, end int) {
				defer wg.Done()
				for i := start; i < end; i++ {
					switch img.pixelType {
					case PixelTypeUInt8:
						newPixelData[i] = float32(img.pixels.([]uint8)[i])
					case PixelTypeInt8:
						newPixelData[i] = float32(img.pixels.([]int8)[i])
					case PixelTypeUInt16:
						newPixelData[i] = float32(img.pixels.([]uint16)[i])
					case PixelTypeInt16:
						newPixelData[i] = float32(img.pixels.([]int16)[i])
					case PixelTypeUInt32:
						newPixelData[i] = float32(img.pixels.([]uint32)[i])
					case PixelTypeInt32:
						newPixelData[i] = float32(img.pixels.([]int32)[i])
					case PixelTypeUInt64:
						newPixelData[i] = float32(img.pixels.([]uint64)[i])
					case PixelTypeInt64:
						newPixelData[i] = float32(img.pixels.([]int64)[i])
					case PixelTypeFloat64:
						newPixelData[i] = float32(img.pixels.([]float64)[i])
					}
				}
			}(start, end)
		}
		wg.Wait()
		newImg.pixels = newPixelData
	case PixelTypeFloat64:
		newPixelData := make([]float64, numPixels)
		for chunk := 0; chunk < numGoroutines; chunk++ {
			start := chunk * chunkSize
			end := start + chunkSize
			if end > numPixels {
				end = numPixels
			}
			wg.Add(1)
			go func(start, end int) {
				defer wg.Done()
				for i := start; i < end; i++ {
					switch img.pixelType {
					case PixelTypeUInt8:
						newPixelData[i] = float64(img.pixels.([]uint8)[i])
					case PixelTypeInt8:
						newPixelData[i] = float64(img.pixels.([]int8)[i])
					case PixelTypeUInt16:
						newPixelData[i] = float64(img.pixels.([]uint16)[i])
					case PixelTypeInt16:
						newPixelData[i] = float64(img.pixels.([]int16)[i])
					case PixelTypeUInt32:
						newPixelData[i] = float64(img.pixels.([]uint32)[i])
					case PixelTypeInt32:
						newPixelData[i] = float64(img.pixels.([]int32)[i])
					case PixelTypeUInt64:
						newPixelData[i] = float64(img.pixels.([]uint64)[i])
					case PixelTypeInt64:
						newPixelData[i] = float64(img.pixels.([]int64)[i])
					case PixelTypeFloat32:
						newPixelData[i] = float64(img.pixels.([]float32)[i])
					}
				}
			}(start, end)
		}
		wg.Wait()
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
	switch img.pixelType {
	case PixelTypeUInt8:
		pixelValue, err := getValueAsPixelType(value, PixelTypeUInt8)
		if err != nil {
			return err
		}
		img.pixels.([]uint8)[idx] = pixelValue.(uint8)
	case PixelTypeInt8:
		pixelValue, err := getValueAsPixelType(value, PixelTypeInt8)
		if err != nil {
			return err
		}
		img.pixels.([]int8)[idx] = pixelValue.(int8)
	case PixelTypeUInt16:
		pixelValue, err := getValueAsPixelType(value, PixelTypeUInt16)
		if err != nil {
			return err
		}
		img.pixels.([]uint16)[idx] = pixelValue.(uint16)
	case PixelTypeInt16:
		pixelValue, err := getValueAsPixelType(value, PixelTypeInt16)
		if err != nil {
			return err
		}
		img.pixels.([]int16)[idx] = pixelValue.(int16)
	case PixelTypeUInt32:
		pixelValue, err := getValueAsPixelType(value, PixelTypeUInt32)
		if err != nil {
			return err
		}
		img.pixels.([]uint32)[idx] = pixelValue.(uint32)
	case PixelTypeInt32:
		pixelValue, err := getValueAsPixelType(value, PixelTypeInt32)
		if err != nil {
			return err
		}
		img.pixels.([]int32)[idx] = pixelValue.(int32)
	case PixelTypeUInt64:
		pixelValue, err := getValueAsPixelType(value, PixelTypeUInt64)
		if err != nil {
			return err
		}
		img.pixels.([]uint64)[idx] = pixelValue.(uint64)
	case PixelTypeInt64:
		pixelValue, err := getValueAsPixelType(value, PixelTypeInt64)
		if err != nil {
			return err
		}
		img.pixels.([]int64)[idx] = pixelValue.(int64)
	case PixelTypeFloat32:
		pixelValue, err := getValueAsPixelType(value, PixelTypeFloat32)
		if err != nil {
			return err
		}
		img.pixels.([]float32)[idx] = pixelValue.(float32)
	case PixelTypeFloat64:
		pixelValue, err := getValueAsPixelType(value, PixelTypeFloat64)
		if err != nil {
			return err
		}
		img.pixels.([]float64)[idx] = pixelValue.(float64)
	default:
		return fmt.Errorf("unsupported pixel type: %d", img.pixelType)
	}
	return nil
}

func (img *Image) SetPixels(pixels any) error {
	var elemType reflect.Type
	flattened := flatten(pixels, &elemType)

	numPixels := 1
	for _, s := range img.size {
		numPixels *= int(s)
	}

	if len(flattened) != numPixels {
		return fmt.Errorf("invalid number of pixels, expected %d, got %d", numPixels, len(flattened))
	}

	switch img.pixelType {
	case PixelTypeUInt8:
		if _, ok := flattened[0].Interface().(uint8); !ok {
			return fmt.Errorf("invalid pixel type, expected uint8, got %T", flattened[0].Interface())
		}
		pixelsSlice := make([]uint8, numPixels)
		for i := 0; i < len(flattened); i++ {
			pixelsSlice[i] = flattened[i].Interface().(uint8)
		}
		img.pixels = pixelsSlice
	case PixelTypeInt8:
		if _, ok := flattened[0].Interface().(int8); !ok {
			return fmt.Errorf("invalid pixel type, expected int8, got %T", flattened[0].Interface())
		}
		pixelsSlice := make([]int8, numPixels)
		for i := 0; i < len(flattened); i++ {
			pixelsSlice[i] = flattened[i].Interface().(int8)
		}
		img.pixels = pixelsSlice
	case PixelTypeUInt16:
		if _, ok := flattened[0].Interface().(uint16); !ok {
			return fmt.Errorf("invalid pixel type, expected uint16, got %T", flattened[0].Interface())
		}
		pixelsSlice := make([]uint16, numPixels)
		for i := 0; i < len(flattened); i++ {
			pixelsSlice[i] = flattened[i].Interface().(uint16)
		}
		img.pixels = pixelsSlice
	case PixelTypeInt16:
		if _, ok := flattened[0].Interface().(int16); !ok {
			return fmt.Errorf("invalid pixel type, expected int16, got %T", flattened[0].Interface())
		}
		pixelsSlice := make([]int16, numPixels)
		for i := 0; i < len(flattened); i++ {
			pixelsSlice[i] = flattened[i].Interface().(int16)
		}
		img.pixels = pixelsSlice
	case PixelTypeUInt32:
		if _, ok := flattened[0].Interface().(uint32); !ok {
			return fmt.Errorf("invalid pixel type, expected uint32, got %T", flattened[0].Interface())
		}
		pixelsSlice := make([]uint32, numPixels)
		for i := 0; i < len(flattened); i++ {
			pixelsSlice[i] = flattened[i].Interface().(uint32)
		}
		img.pixels = pixelsSlice
	case PixelTypeInt32:
		if _, ok := flattened[0].Interface().(int32); !ok {
			return fmt.Errorf("invalid pixel type, expected int32, got %T", flattened[0].Interface())
		}
		pixelsSlice := make([]int32, numPixels)
		for i := 0; i < len(flattened); i++ {
			pixelsSlice[i] = flattened[i].Interface().(int32)
		}
		img.pixels = pixelsSlice
	case PixelTypeUInt64:
		if _, ok := flattened[0].Interface().(uint64); !ok {
			return fmt.Errorf("invalid pixel type, expected uint64, got %T", flattened[0].Interface())
		}
		pixelsSlice := make([]uint64, numPixels)
		for i := 0; i < len(flattened); i++ {
			pixelsSlice[i] = flattened[i].Interface().(uint64)
		}
		img.pixels = pixelsSlice
	case PixelTypeInt64:
		if _, ok := flattened[0].Interface().(int64); !ok {
			return fmt.Errorf("invalid pixel type, expected int64, got %T", flattened[0].Interface())
		}
		pixelsSlice := make([]int64, numPixels)
		for i := 0; i < len(flattened); i++ {
			pixelsSlice[i] = flattened[i].Interface().(int64)
		}
		img.pixels = pixelsSlice
	case PixelTypeFloat32:
		if _, ok := flattened[0].Interface().(float32); !ok {
			return fmt.Errorf("invalid pixel type, expected float32, got %T", flattened[0].Interface())
		}
		pixelsSlice := make([]float32, numPixels)
		for i := 0; i < len(flattened); i++ {
			pixelsSlice[i] = flattened[i].Interface().(float32)
		}
		img.pixels = pixelsSlice
	case PixelTypeFloat64:
		if _, ok := flattened[0].Interface().(float64); !ok {
			return fmt.Errorf("invalid pixel type, expected float64, got %T", flattened[0].Interface())
		}
		pixelsSlice := make([]float64, numPixels)
		for i := 0; i < len(flattened); i++ {
			pixelsSlice[i] = flattened[i].Interface().(float64)
		}
		img.pixels = pixelsSlice
	default:
		return fmt.Errorf("unsupported pixel type: %d", img.pixelType)
	}

	return nil
}

func (img *Image) SetSize(size []uint32) error {
	totalSize := uint32(1)
	for _, s := range size {
		totalSize *= s
	}
	var numPixels int
	switch img.pixelType {
	case PixelTypeUInt8:
		numPixels = int(len(img.pixels.([]uint8)))
	case PixelTypeInt8:
		numPixels = int(len(img.pixels.([]int8)))
	case PixelTypeUInt16:
		numPixels = int(len(img.pixels.([]uint16)))
	case PixelTypeInt16:
		numPixels = int(len(img.pixels.([]int16)))
	case PixelTypeUInt32:
		numPixels = int(len(img.pixels.([]uint32)))
	case PixelTypeInt32:
		numPixels = int(len(img.pixels.([]int32)))
	case PixelTypeUInt64:
		numPixels = int(len(img.pixels.([]uint64)))
	case PixelTypeInt64:
		numPixels = int(len(img.pixels.([]int64)))
	case PixelTypeFloat32:
		numPixels = int(len(img.pixels.([]float32)))
	case PixelTypeFloat64:
		numPixels = int(len(img.pixels.([]float64)))
	}
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
