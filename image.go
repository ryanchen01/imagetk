package imagetk

import (
	"fmt"
	"reflect"
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
// - dimension: The number of dimensions of the image.
// - size: A slice of uint32 representing the size of the image in each dimension.
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
	if len(size) == 0 {
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
	var size []uint32
	value := reflect.ValueOf(data)

	if value.Kind() != reflect.Slice && value.Kind() != reflect.Array {
		return nil, fmt.Errorf("data must be a slice or array, got %s", value.Kind().String())
	}
	pixelType := PixelTypeUnknown

	for value.Kind() == reflect.Slice {
		size = append(size, uint32(value.Len()))
		if value.Len() == 0 {
			break
		}
		value = value.Index(0)
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

func (img *Image) IsPixelType(pixelType int) bool {
	return img.pixelType == pixelType
}

func (img *Image) GetDimension() uint32 {
	return img.dimension
}

func (img *Image) GetSize() []uint32 {
	return img.size
}

func (img *Image) GetSpacing() []float64 {
	return img.spacing
}

func (img *Image) GetOrigin() []float64 {
	return img.origin
}

func (img *Image) GetDirection() [9]float64 {
	return img.direction
}

func (img *Image) GetPixel(index []uint32) (any, error) {
	if len(index) != int(img.dimension) {
		return nil, fmt.Errorf("invalid index length: %d", len(index))
	}
	idx := uint32(0)
	for i := 0; i < len(index); i++ {
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

func (img *Image) GetPixelAsUInt8(index []uint32) (uint8, error) {
	idx := uint32(0)
	for i := 0; i < len(index); i++ {
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

func (img *Image) GetPixelAsInt8(index []uint32) (int8, error) {
	idx := uint32(0)
	for i := 0; i < len(index); i++ {
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

func (img *Image) GetPixelAsUInt16(index []uint32) (uint16, error) {
	idx := uint32(0)
	for i := 0; i < len(index); i++ {
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

func (img *Image) GetPixelAsInt16(index []uint32) (int16, error) {
	idx := uint32(0)
	for i := 0; i < len(index); i++ {
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

func (img *Image) GetPixelAsUInt32(index []uint32) (uint32, error) {
	idx := uint32(0)
	for i := 0; i < len(index); i++ {
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

func (img *Image) GetPixelAsInt32(index []uint32) (int32, error) {
	idx := uint32(0)
	for i := 0; i < len(index); i++ {
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

func (img *Image) GetPixelAsUInt64(index []uint32) (uint64, error) {
	idx := uint32(0)
	for i := 0; i < len(index); i++ {
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

func (img *Image) GetPixelAsInt64(index []uint32) (int64, error) {
	idx := uint32(0)
	for i := 0; i < len(index); i++ {
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

func (img *Image) GetPixelAsFloat32(index []uint32) (float32, error) {
	idx := uint32(0)
	for i := 0; i < len(index); i++ {
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

func (img *Image) GetPixelAsFloat64(index []uint32) (float64, error) {
	idx := uint32(0)
	for i := 0; i < len(index); i++ {
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
		return 0, fmt.Errorf("unsupported pixel type: %d", img.pixelType)
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

	switch pixelType {
	case PixelTypeUInt8:
		newPixelData := make([]uint8, numPixels)
		for i := 0; i < numPixels; i++ {
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
		newImg.pixels = newPixelData
	case PixelTypeInt8:
		newPixelData := make([]int8, numPixels)
		for i := 0; i < numPixels; i++ {
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
		newImg.pixels = newPixelData
	case PixelTypeUInt16:
		newPixelData := make([]uint16, numPixels)
		for i := 0; i < numPixels; i++ {
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
		newImg.pixels = newPixelData
	case PixelTypeInt16:
		newPixelData := make([]int16, numPixels)
		for i := 0; i < numPixels; i++ {
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
		newImg.pixels = newPixelData
	case PixelTypeUInt32:
		newPixelData := make([]uint32, numPixels)
		for i := 0; i < numPixels; i++ {
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
		newImg.pixels = newPixelData
	case PixelTypeInt32:
		newPixelData := make([]int32, numPixels)
		for i := 0; i < numPixels; i++ {
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
		newImg.pixels = newPixelData
	case PixelTypeUInt64:
		newPixelData := make([]uint64, numPixels)
		for i := 0; i < numPixels; i++ {
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
		newImg.pixels = newPixelData
	case PixelTypeInt64:
		newPixelData := make([]int64, numPixels)
		for i := 0; i < numPixels; i++ {
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
		newImg.pixels = newPixelData
	case PixelTypeFloat32:
		newPixelData := make([]float32, numPixels)
		for i := 0; i < numPixels; i++ {
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
		newImg.pixels = newPixelData
	case PixelTypeFloat64:
		newPixelData := make([]float64, numPixels)
		for i := 0; i < numPixels; i++ {
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
	for i := 0; i < len(index); i++ {
		if index[i] >= img.size[i] {
			return fmt.Errorf("index out of range: %d", index[i])
		}
		idx = idx*img.size[i] + index[i]
	}
	switch img.pixelType {
	case PixelTypeUInt8:
		if _, ok := value.(uint8); !ok {
			return fmt.Errorf("invalid value type, expected uint8, got %T", value)
		}
		img.pixels.([]uint8)[idx] = value.(uint8)
	case PixelTypeInt8:
		if _, ok := value.(int8); !ok {
			return fmt.Errorf("invalid value type, expected int8, got %T", value)
		}
		img.pixels.([]int8)[idx] = value.(int8)
	case PixelTypeUInt16:
		if _, ok := value.(uint16); !ok {
			return fmt.Errorf("invalid value type, expected uint16, got %T", value)
		}
		img.pixels.([]uint16)[idx] = value.(uint16)
	case PixelTypeInt16:
		if _, ok := value.(int16); !ok {
			return fmt.Errorf("invalid value type, expected int16, got %T", value)
		}
		img.pixels.([]int16)[idx] = value.(int16)
	case PixelTypeUInt32:
		if _, ok := value.(uint32); !ok {
			return fmt.Errorf("invalid value type, expected uint32, got %T", value)
		}
		img.pixels.([]uint32)[idx] = value.(uint32)
	case PixelTypeInt32:
		if _, ok := value.(int32); !ok {
			return fmt.Errorf("invalid value type, expected int32, got %T", value)
		}
		img.pixels.([]int32)[idx] = value.(int32)
	case PixelTypeUInt64:
		if _, ok := value.(uint64); !ok {
			return fmt.Errorf("invalid value type, expected uint64, got %T", value)
		}
		img.pixels.([]uint64)[idx] = value.(uint64)
	case PixelTypeInt64:
		if _, ok := value.(int64); !ok {
			return fmt.Errorf("invalid value type, expected int64, got %T", value)
		}
		img.pixels.([]int64)[idx] = value.(int64)
	case PixelTypeFloat32:
		if _, ok := value.(float32); !ok {
			return fmt.Errorf("invalid value type, expected float32, got %T", value)
		}
		img.pixels.([]float32)[idx] = value.(float32)
	case PixelTypeFloat64:
		if _, ok := value.(float64); !ok {
			return fmt.Errorf("invalid value type, expected float64, got %T", value)
		}
		img.pixels.([]float64)[idx] = value.(float64)
	default:
		return fmt.Errorf("unsupported pixel type: %d", img.pixelType)
	}
	return nil
}

func (img *Image) SetPixels(pixels any) error {
	flattened := flattenReflectValues(pixels)

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
