package imagetk

import (
	"encoding/binary"
	"math"
	"sort"
)

// Min returns the minimum pixel value in the image.
//
// Returns:
//   - any: The minimum pixel value in the image as the type of the image.
func (img *Image) Min() any {
	switch img.pixelType {
	case PixelTypeUInt8:
		minValue := uint8(math.MaxUint8)
		for i := 0; i < len(img.pixels); i++ {
			if uint8(img.pixels[i]) < minValue {
				minValue = uint8(img.pixels[i])
			}
		}
		return minValue
	case PixelTypeInt8:
		minValue := int8(math.MaxInt8)
		for i := 0; i < len(img.pixels); i++ {
			if int8(img.pixels[i]) < minValue {
				minValue = int8(img.pixels[i])
			}
		}
		return minValue
	case PixelTypeUInt16:
		minValue := uint16(math.MaxUint16)
		for i := 0; i < len(img.pixels)/2; i++ {
			value := binary.LittleEndian.Uint16(img.pixels[i*2 : i*2+2])
			if value < minValue {
				minValue = value
			}
		}
		return minValue
	case PixelTypeInt16:
		minValue := int16(math.MaxInt16)
		for i := 0; i < len(img.pixels)/2; i++ {
			value := int16(binary.LittleEndian.Uint16(img.pixels[i*2 : i*2+2]))
			if value < minValue {
				minValue = value
			}
		}
		return minValue
	case PixelTypeUInt32:
		minValue := uint32(math.MaxUint32)
		for i := 0; i < len(img.pixels)/4; i++ {
			value := binary.LittleEndian.Uint32(img.pixels[i*4 : i*4+4])
			if value < minValue {
				minValue = value
			}
		}
		return minValue
	case PixelTypeInt32:
		minValue := int32(math.MaxInt32)
		for i := 0; i < len(img.pixels)/4; i++ {
			value := int32(binary.LittleEndian.Uint32(img.pixels[i*4 : i*4+4]))
			if value < minValue {
				minValue = value
			}
		}
		return minValue
	case PixelTypeUInt64:
		minValue := uint64(math.MaxUint64)
		for i := 0; i < len(img.pixels)/8; i++ {
			value := binary.LittleEndian.Uint64(img.pixels[i*8 : i*8+8])
			if value < minValue {
				minValue = value
			}
		}
		return minValue
	case PixelTypeInt64:
		minValue := int64(math.MaxInt64)
		for i := 0; i < len(img.pixels)/8; i++ {
			value := int64(binary.LittleEndian.Uint64(img.pixels[i*8 : i*8+8]))
			if value < minValue {
				minValue = value
			}
		}
		return minValue
	case PixelTypeFloat32:
		minValue := float32(math.MaxFloat32)
		for i := 0; i < len(img.pixels)/4; i++ {
			value := math.Float32frombits(binary.LittleEndian.Uint32(img.pixels[i*4 : i*4+4]))
			if value < minValue {
				minValue = value
			}
		}
		return minValue
	case PixelTypeFloat64:
		minValue := float64(math.MaxFloat64)
		for i := 0; i < len(img.pixels)/8; i++ {
			value := math.Float64frombits(binary.LittleEndian.Uint64(img.pixels[i*8 : i*8+8]))
			if value < minValue {
				minValue = value
			}
		}
		return minValue
	default:
		return nil
	}
}

// Max returns the maximum pixel value in the image.
//
// Returns:
//   - any: The maximum pixel value in the image as the type of the image.
func (img *Image) Max() any {
	switch img.pixelType {
	case PixelTypeUInt8:
		maxValue := uint8(0)
		for i := 0; i < len(img.pixels); i++ {
			if uint8(img.pixels[i]) > maxValue {
				maxValue = uint8(img.pixels[i])
			}
		}
		return maxValue
	case PixelTypeInt8:
		maxValue := int8(math.MinInt8)
		for i := 0; i < len(img.pixels); i++ {
			if int8(img.pixels[i]) > maxValue {
				maxValue = int8(img.pixels[i])
			}
		}
		return maxValue
	case PixelTypeUInt16:
		maxValue := uint16(0)
		for i := 0; i < len(img.pixels)/2; i++ {
			value := binary.LittleEndian.Uint16(img.pixels[i*2 : i*2+2])
			if value > maxValue {
				maxValue = value
			}
		}
		return maxValue
	case PixelTypeInt16:
		maxValue := int16(math.MinInt16)
		for i := 0; i < len(img.pixels)/2; i++ {
			value := int16(binary.LittleEndian.Uint16(img.pixels[i*2 : i*2+2]))
			if value > maxValue {
				maxValue = value
			}
		}
		return maxValue
	case PixelTypeUInt32:
		maxValue := uint32(0)
		for i := 0; i < len(img.pixels)/4; i++ {
			value := binary.LittleEndian.Uint32(img.pixels[i*4 : i*4+4])
			if value > maxValue {
				maxValue = value
			}
		}
		return maxValue
	case PixelTypeInt32:
		maxValue := int32(math.MinInt32)
		for i := 0; i < len(img.pixels)/4; i++ {
			value := int32(binary.LittleEndian.Uint32(img.pixels[i*4 : i*4+4]))
			if value > maxValue {
				maxValue = value
			}
		}
		return maxValue
	case PixelTypeUInt64:
		maxValue := uint64(0)
		for i := 0; i < len(img.pixels)/8; i++ {
			value := binary.LittleEndian.Uint64(img.pixels[i*8 : i*8+8])
			if value > maxValue {
				maxValue = value
			}
		}
		return maxValue
	case PixelTypeInt64:
		maxValue := int64(math.MinInt64)
		for i := 0; i < len(img.pixels)/8; i++ {
			value := int64(binary.LittleEndian.Uint64(img.pixels[i*8 : i*8+8]))
			if value > maxValue {
				maxValue = value
			}
		}
		return maxValue
	case PixelTypeFloat32:
		maxValue := float32(-math.MaxFloat32)
		for i := 0; i < len(img.pixels)/4; i++ {
			value := math.Float32frombits(binary.LittleEndian.Uint32(img.pixels[i*4 : i*4+4]))
			if value > maxValue {
				maxValue = value
			}
		}
		return maxValue
	case PixelTypeFloat64:
		maxValue := float64(-math.MaxFloat64)
		for i := 0; i < len(img.pixels)/8; i++ {
			value := math.Float64frombits(binary.LittleEndian.Uint64(img.pixels[i*8 : i*8+8]))
			if value > maxValue {
				maxValue = value
			}
		}
		return maxValue
	default:
		return nil
	}
}

// Sum returns the sum of all pixel values in the image.
//
// Returns:
//   - any: The sum of all pixel values in the image. Unsigned pixel types are summed as uint64, signed pixel types are summed as int64, and floating-point pixel types are summed as float64.
func (img *Image) Sum() any {
	switch img.pixelType {
	case PixelTypeUInt8:
		sumValue := uint64(0)
		for i := 0; i < len(img.pixels); i++ {
			sumValue += uint64(uint8(img.pixels[i]))
		}
		return sumValue
	case PixelTypeInt8:
		sumValue := int64(0)
		for i := 0; i < len(img.pixels); i++ {
			sumValue += int64(int8(img.pixels[i]))
		}
		return sumValue
	case PixelTypeUInt16:
		sumValue := uint64(0)
		for i := 0; i < len(img.pixels)/2; i++ {
			value := binary.LittleEndian.Uint16(img.pixels[i*2 : i*2+2])
			sumValue += uint64(value)
		}
		return sumValue
	case PixelTypeInt16:
		sumValue := int64(0)
		for i := 0; i < len(img.pixels)/2; i++ {
			value := int16(binary.LittleEndian.Uint16(img.pixels[i*2 : i*2+2]))
			sumValue += int64(value)
		}
		return sumValue
	case PixelTypeUInt32:
		sumValue := uint64(0)
		for i := 0; i < len(img.pixels)/4; i++ {
			value := binary.LittleEndian.Uint32(img.pixels[i*4 : i*4+4])
			sumValue += uint64(value)
		}
		return sumValue
	case PixelTypeInt32:
		sumValue := int64(0)
		for i := 0; i < len(img.pixels)/4; i++ {
			value := int32(binary.LittleEndian.Uint32(img.pixels[i*4 : i*4+4]))
			sumValue += int64(value)
		}
		return sumValue
	case PixelTypeUInt64:
		sumValue := uint64(0)
		for i := 0; i < len(img.pixels)/8; i++ {
			value := binary.LittleEndian.Uint64(img.pixels[i*8 : i*8+8])
			sumValue += uint64(value)
		}
		return sumValue
	case PixelTypeInt64:
		sumValue := int64(0)
		for i := 0; i < len(img.pixels)/8; i++ {
			value := int64(binary.LittleEndian.Uint64(img.pixels[i*8 : i*8+8]))
			sumValue += int64(value)
		}
		return sumValue
	case PixelTypeFloat32:
		sumValue := float64(0)
		for i := 0; i < len(img.pixels)/4; i++ {
			value := math.Float32frombits(binary.LittleEndian.Uint32(img.pixels[i*4 : i*4+4]))
			sumValue += float64(value)
		}
		return sumValue
	case PixelTypeFloat64:
		sumValue := float64(0)
		for i := 0; i < len(img.pixels)/8; i++ {
			value := math.Float64frombits(binary.LittleEndian.Uint64(img.pixels[i*8 : i*8+8]))
			sumValue += float64(value)
		}
		return sumValue
	default:
		return nil
	}
}

// Product returns the product of all pixel values in the image.
//
// Returns:
//   - any: The product of all pixel values in the image. Unsigned pixel types are multiplied as uint64, signed pixel types are multiplied as int64, and floating-point pixel types are multiplied as float64.
func (img *Image) Product() any {
	switch img.pixelType {
	case PixelTypeUInt8:
		productValue := uint64(1)
		for i := 0; i < len(img.pixels); i++ {
			productValue *= uint64(uint8(img.pixels[i]))
		}
		return productValue
	case PixelTypeInt8:
		productValue := int64(1)
		for i := 0; i < len(img.pixels); i++ {
			productValue *= int64(int8(img.pixels[i]))
		}
		return productValue
	case PixelTypeUInt16:
		productValue := uint64(1)
		for i := 0; i < len(img.pixels)/2; i++ {
			value := binary.LittleEndian.Uint16(img.pixels[i*2 : i*2+2])
			productValue *= uint64(value)
		}
		return productValue
	case PixelTypeInt16:
		productValue := int64(1)
		for i := 0; i < len(img.pixels)/2; i++ {
			value := int16(binary.LittleEndian.Uint16(img.pixels[i*2 : i*2+2]))
			productValue *= int64(value)
		}
		return productValue
	case PixelTypeUInt32:
		productValue := uint64(1)
		for i := 0; i < len(img.pixels)/4; i++ {
			value := binary.LittleEndian.Uint32(img.pixels[i*4 : i*4+4])
			productValue *= uint64(value)
		}
		return productValue
	case PixelTypeInt32:
		productValue := int64(1)
		for i := 0; i < len(img.pixels)/4; i++ {
			value := int32(binary.LittleEndian.Uint32(img.pixels[i*4 : i*4+4]))
			productValue *= int64(value)
		}
		return productValue
	case PixelTypeUInt64:
		productValue := uint64(1)
		for i := 0; i < len(img.pixels)/8; i++ {
			value := binary.LittleEndian.Uint64(img.pixels[i*8 : i*8+8])
			productValue *= uint64(value)
		}
		return productValue
	case PixelTypeInt64:
		productValue := int64(1)
		for i := 0; i < len(img.pixels)/8; i++ {
			value := int64(binary.LittleEndian.Uint64(img.pixels[i*8 : i*8+8]))
			productValue *= int64(value)
		}
		return productValue
	case PixelTypeFloat32:
		productValue := float64(1)
		for i := 0; i < len(img.pixels)/4; i++ {
			value := math.Float32frombits(binary.LittleEndian.Uint32(img.pixels[i*4 : i*4+4]))
			productValue *= float64(value)
		}
		return productValue
	case PixelTypeFloat64:
		productValue := float64(1)
		for i := 0; i < len(img.pixels)/8; i++ {
			value := math.Float64frombits(binary.LittleEndian.Uint64(img.pixels[i*8 : i*8+8]))
			productValue *= float64(value)
		}
		return productValue
	default:
		return nil
	}
}

// ExactMean returns the exact mean of all pixel values in the image.
//
// Returns:
//   - float64: The exact mean of the image.
func (img *Image) ExactMean() float64 {
	switch img.pixelType {
	case PixelTypeUInt8, PixelTypeUInt16, PixelTypeUInt32, PixelTypeUInt64:
		sumValue := img.Sum().(uint64)
		return float64(sumValue) / float64(len(img.pixels)/img.bytesPerPixel)
	case PixelTypeInt8, PixelTypeInt16, PixelTypeInt32, PixelTypeInt64:
		sumValue := img.Sum().(int64)
		return float64(sumValue) / float64(len(img.pixels)/img.bytesPerPixel)
	case PixelTypeFloat32, PixelTypeFloat64:
		sumValue := img.Sum().(float64)
		return sumValue / float64(len(img.pixels)/img.bytesPerPixel)
	default:
		return 0
	}
}

// Mean returns the mean of the image.
//
// Returns:
//   - any: The mean of the image as the type of the image.
func (img *Image) Mean() any {
	switch img.pixelType {
	case PixelTypeUInt8:
		return uint8(img.ExactMean())
	case PixelTypeInt8:
		return int8(img.ExactMean())
	case PixelTypeUInt16:
		return uint16(img.ExactMean())
	case PixelTypeInt16:
		return int16(img.ExactMean())
	case PixelTypeUInt32:
		return uint32(img.ExactMean())
	case PixelTypeInt32:
		return int32(img.ExactMean())
	case PixelTypeUInt64:
		return uint64(img.ExactMean())
	case PixelTypeInt64:
		return int64(img.ExactMean())
	case PixelTypeFloat32:
		return float32(img.ExactMean())
	case PixelTypeFloat64:
		return img.ExactMean()
	default:
		return nil
	}
}

// Median returns the median of the image.
//
// Returns:
//   - float64: The median of the image.
func (img *Image) Median() float64 {
	pixelData := make([]byte, len(img.pixels))
	copy(pixelData, img.pixels)
	switch img.pixelType {
	case PixelTypeUInt8:
		sort.Slice(pixelData, func(i, j int) bool { return uint8(pixelData[i]) < uint8(pixelData[j]) })
		if len(pixelData)%2 == 0 {
			return float64(pixelData[len(pixelData)/2-1]+pixelData[len(pixelData)/2]) / 2
		}
		return float64(uint8(pixelData[len(pixelData)/2]))
	case PixelTypeInt8:
		sort.Slice(pixelData, func(i, j int) bool { return int8(pixelData[i]) < int8(pixelData[j]) })
		if len(pixelData)%2 == 0 {
			return float64(pixelData[len(pixelData)/2-1]+pixelData[len(pixelData)/2]) / 2
		}
		return float64(int8(pixelData[len(pixelData)/2]))
	case PixelTypeUInt16:
		sort.Slice(pixelData, func(i, j int) bool {
			return binary.LittleEndian.Uint16(pixelData[i*2:i*2+2]) < binary.LittleEndian.Uint16(pixelData[j*2:j*2+2])
		})
		if len(pixelData)%2 == 0 {
			return float64(binary.LittleEndian.Uint16(pixelData[len(pixelData)/2*2-2:len(pixelData)/2*2])+binary.LittleEndian.Uint16(pixelData[len(pixelData)/2*2:len(pixelData)/2*2+2])) / 2
		}
		return float64(binary.LittleEndian.Uint16(pixelData[len(pixelData)/2*2 : len(pixelData)/2*2+2]))
	case PixelTypeInt16:
		sort.Slice(pixelData, func(i, j int) bool {
			return int16(binary.LittleEndian.Uint16(pixelData[i*2:i*2+2])) < int16(binary.LittleEndian.Uint16(pixelData[j*2:j*2+2]))
		})
		if len(pixelData)%2 == 0 {
			return float64(int16(binary.LittleEndian.Uint16(pixelData[len(pixelData)/2*2-2:len(pixelData)/2*2]))+int16(binary.LittleEndian.Uint16(pixelData[len(pixelData)/2*2:len(pixelData)/2*2+2]))) / 2
		}
		return float64(int16(binary.LittleEndian.Uint16(pixelData[len(pixelData)/2*2 : len(pixelData)/2*2+2])))
	case PixelTypeUInt32:
		sort.Slice(pixelData, func(i, j int) bool {
			return binary.LittleEndian.Uint32(pixelData[i*4:i*4+4]) < binary.LittleEndian.Uint32(pixelData[j*4:j*4+4])
		})
		if len(pixelData)%2 == 0 {
			return float64(binary.LittleEndian.Uint32(pixelData[len(pixelData)/2*4-4:len(pixelData)/2*4])+binary.LittleEndian.Uint32(pixelData[len(pixelData)/2*4:len(pixelData)/2*4+4])) / 2
		}
		return float64(binary.LittleEndian.Uint32(pixelData[len(pixelData)/2*4 : len(pixelData)/2*4+4]))
	case PixelTypeInt32:
		sort.Slice(pixelData, func(i, j int) bool {
			return int32(binary.LittleEndian.Uint32(pixelData[i*4:i*4+4])) < int32(binary.LittleEndian.Uint32(pixelData[j*4:j*4+4]))
		})
		if len(pixelData)%2 == 0 {
			return float64(int32(binary.LittleEndian.Uint32(pixelData[len(pixelData)/2*4-4:len(pixelData)/2*4]))+int32(binary.LittleEndian.Uint32(pixelData[len(pixelData)/2*4:len(pixelData)/2*4+4]))) / 2
		}
		return float64(int32(binary.LittleEndian.Uint32(pixelData[len(pixelData)/2*4 : len(pixelData)/2*4+4])))
	case PixelTypeUInt64:
		sort.Slice(pixelData, func(i, j int) bool {
			return binary.LittleEndian.Uint64(pixelData[i*8:i*8+8]) < binary.LittleEndian.Uint64(pixelData[j*8:j*8+8])
		})
		if len(pixelData)%2 == 0 {
			return float64(binary.LittleEndian.Uint64(pixelData[len(pixelData)/2*8-8:len(pixelData)/2*8])+binary.LittleEndian.Uint64(pixelData[len(pixelData)/2*8:len(pixelData)/2*8+8])) / 2
		}
		return float64(binary.LittleEndian.Uint64(pixelData[len(pixelData)/2*8 : len(pixelData)/2*8+8]))
	case PixelTypeInt64:
		sort.Slice(pixelData, func(i, j int) bool {
			return int64(binary.LittleEndian.Uint64(pixelData[i*8:i*8+8])) < int64(binary.LittleEndian.Uint64(pixelData[j*8:j*8+8]))
		})
		if len(pixelData)%2 == 0 {
			return float64(int64(binary.LittleEndian.Uint64(pixelData[len(pixelData)/2*8-8:len(pixelData)/2*8]))+int64(binary.LittleEndian.Uint64(pixelData[len(pixelData)/2*8:len(pixelData)/2*8+8]))) / 2
		}
		return float64(int64(binary.LittleEndian.Uint64(pixelData[len(pixelData)/2*8 : len(pixelData)/2*8+8])))
	case PixelTypeFloat32:
		sort.Slice(pixelData, func(i, j int) bool {
			return math.Float32frombits(binary.LittleEndian.Uint32(pixelData[i*4:i*4+4])) < math.Float32frombits(binary.LittleEndian.Uint32(pixelData[j*4:j*4+4]))
		})
		if len(pixelData)%2 == 0 {
			return float64(math.Float32frombits(binary.LittleEndian.Uint32(pixelData[len(pixelData)/2*4-4:len(pixelData)/2*4]))+math.Float32frombits(binary.LittleEndian.Uint32(pixelData[len(pixelData)/2*4:len(pixelData)/2*4+4]))) / 2
		}
		return float64(math.Float32frombits(binary.LittleEndian.Uint32(pixelData[len(pixelData)/2*4 : len(pixelData)/2*4+4])))
	case PixelTypeFloat64:
		sort.Slice(pixelData, func(i, j int) bool {
			return math.Float64frombits(binary.LittleEndian.Uint64(pixelData[i*8:i*8+8])) < math.Float64frombits(binary.LittleEndian.Uint64(pixelData[j*8:j*8+8]))
		})
		if len(pixelData)%2 == 0 {
			return float64(math.Float64frombits(binary.LittleEndian.Uint64(pixelData[len(pixelData)/2*8-8:len(pixelData)/2*8]))+math.Float64frombits(binary.LittleEndian.Uint64(pixelData[len(pixelData)/2*8:len(pixelData)/2*8+8]))) / 2
		}
		return float64(math.Float64frombits(binary.LittleEndian.Uint64(pixelData[len(pixelData)/2*8 : len(pixelData)/2*8+8])))
	default:
		return 0
	}
}

// Std returns the standard deviation of the image.
//
// Returns:
//   - any: The standard deviation of the image.
func (img *Image) Std() any {
	switch img.pixelType {
	case PixelTypeUInt8:
		meanValue := img.ExactMean()
		sumValue := 0.0
		for i := 0; i < len(img.pixels); i++ {
			value := float64(uint8(img.pixels[i]))
			sumValue += (value - meanValue) * (value - meanValue)
		}
		return math.Sqrt(sumValue / float64(len(img.pixels)))
	case PixelTypeInt8:
		meanValue := img.ExactMean()
		sumValue := 0.0
		for i := 0; i < len(img.pixels); i++ {
			value := float64(int8(img.pixels[i]))
			sumValue += (value - meanValue) * (value - meanValue)
		}
		return math.Sqrt(sumValue / float64(len(img.pixels)))
	case PixelTypeUInt16:
		meanValue := img.ExactMean()
		sumValue := 0.0
		for i := 0; i < len(img.pixels)/2; i++ {
			value := float64(binary.LittleEndian.Uint16(img.pixels[i*2 : i*2+2]))
			sumValue += (value - meanValue) * (value - meanValue)
		}
		return math.Sqrt(float64(sumValue) / float64(len(img.pixels)/2))
	case PixelTypeInt16:
		meanValue := img.ExactMean()
		sumValue := 0.0
		for i := 0; i < len(img.pixels)/2; i++ {
			value := float64(int16(binary.LittleEndian.Uint16(img.pixels[i*2 : i*2+2])))
			sumValue += (value - meanValue) * (value - meanValue)
		}
		return math.Sqrt(float64(sumValue) / float64(len(img.pixels)/2))
	case PixelTypeUInt32:
		meanValue := img.ExactMean()
		sumValue := 0.0
		for i := 0; i < len(img.pixels)/4; i++ {
			value := float64(binary.LittleEndian.Uint32(img.pixels[i*4 : i*4+4]))
			sumValue += (value - meanValue) * (value - meanValue)
		}
		return math.Sqrt(sumValue / float64(len(img.pixels)/4))
	case PixelTypeInt32:
		meanValue := img.ExactMean()
		sumValue := 0.0
		for i := 0; i < len(img.pixels)/4; i++ {
			value := float64(int32(binary.LittleEndian.Uint32(img.pixels[i*4 : i*4+4])))
			sumValue += (value - meanValue) * (value - meanValue)
		}
		return math.Sqrt(sumValue / float64(len(img.pixels)/4))
	case PixelTypeUInt64:
		meanValue := img.ExactMean()
		sumValue := 0.0
		for i := 0; i < len(img.pixels)/8; i++ {
			value := float64(binary.LittleEndian.Uint64(img.pixels[i*8 : i*8+8]))
			sumValue += (value - meanValue) * (value - meanValue)
		}
		return math.Sqrt(sumValue / float64(len(img.pixels)/8))
	case PixelTypeInt64:
		meanValue := img.ExactMean()
		sumValue := 0.0
		for i := 0; i < len(img.pixels)/8; i++ {
			value := float64(int64(binary.LittleEndian.Uint64(img.pixels[i*8 : i*8+8])))
			sumValue += (value - meanValue) * (value - meanValue)
		}
		return math.Sqrt(sumValue / float64(len(img.pixels)/8))
	case PixelTypeFloat32:
		meanValue := img.ExactMean()
		sumValue := 0.0
		for i := 0; i < len(img.pixels)/4; i++ {
			value := math.Float32frombits(binary.LittleEndian.Uint32(img.pixels[i*4 : i*4+4]))
			sumValue += (float64(value) - meanValue) * (float64(value) - meanValue)
		}
		return math.Sqrt(sumValue / float64(len(img.pixels)/4))
	case PixelTypeFloat64:
		meanValue := img.ExactMean()
		sumValue := 0.0
		for i := 0; i < len(img.pixels)/8; i++ {
			value := math.Float64frombits(binary.LittleEndian.Uint64(img.pixels[i*8 : i*8+8]))
			sumValue += (value - meanValue) * (value - meanValue)
		}
		return math.Sqrt(sumValue / float64(len(img.pixels)/8))
	default:
		return nil
	}
}

// Percentile returns the percentile value of the image.
// Parameters:
//   - p: The percentile to compute (between 0 and 1).
//
// Returns:
//   - float64: The percentile value.
func (img *Image) Percentile(p float64) float64 {
	pixelData := make([]byte, len(img.pixels))
	copy(pixelData, img.pixels)
	switch img.pixelType {
	case PixelTypeUInt8:
		sort.Slice(pixelData, func(i, j int) bool { return uint8(pixelData[i]) < uint8(pixelData[j]) })
		indexFloat := p * float64(len(pixelData)-1)
		index := int(indexFloat)
		if indexFloat == float64(index) {
			return float64(uint8(pixelData[index]))
		}
		value := float64(uint8(img.pixels[index]))
		nextValue := float64(uint8(img.pixels[index+1]))
		return value + (indexFloat-float64(index))*(nextValue-value)
	case PixelTypeInt8:
		sort.Slice(pixelData, func(i, j int) bool { return int8(pixelData[i]) < int8(pixelData[j]) })
		indexFloat := p * float64(len(pixelData)-1)
		index := int(indexFloat)
		if indexFloat == float64(index) {
			return float64(int8(pixelData[index]))
		}
		value := float64(int8(pixelData[index]))
		nextValue := float64(int8(pixelData[index+1]))
		return value + (indexFloat-float64(index))*(nextValue-value)
	case PixelTypeUInt16:
		sort.Slice(pixelData, func(i, j int) bool {
			return binary.LittleEndian.Uint16(pixelData[i*2:i*2+2]) < binary.LittleEndian.Uint16(pixelData[j*2:j*2+2])
		})
		indexFloat := p * float64(len(pixelData)-1)
		index := int(indexFloat)
		if indexFloat == float64(index) {
			return float64(binary.LittleEndian.Uint16(pixelData[index*2 : index*2+2]))
		}
		value := float64(binary.LittleEndian.Uint16(pixelData[index*2 : index*2+2]))
		nextValue := float64(binary.LittleEndian.Uint16(pixelData[(index+1)*2 : (index+1)*2+2]))
		return value + (indexFloat-float64(index))*(nextValue-value)
	case PixelTypeInt16:
		sort.Slice(pixelData, func(i, j int) bool {
			return int16(binary.LittleEndian.Uint16(pixelData[i*2:i*2+2])) < int16(binary.LittleEndian.Uint16(pixelData[j*2:j*2+2]))
		})
		indexFloat := p * float64(len(pixelData)-1)
		index := int(indexFloat)
		if indexFloat == float64(index) {
			return float64(int16(binary.LittleEndian.Uint16(pixelData[index*2 : index*2+2])))
		}
		value := float64(int16(binary.LittleEndian.Uint16(pixelData[index*2 : index*2+2])))
		nextValue := float64(int16(binary.LittleEndian.Uint16(pixelData[(index+1)*2 : (index+1)*2+2])))
		return value + (indexFloat-float64(index))*(nextValue-value)
	case PixelTypeUInt32:
		sort.Slice(pixelData, func(i, j int) bool {
			return binary.LittleEndian.Uint32(pixelData[i*4:i*4+4]) < binary.LittleEndian.Uint32(pixelData[j*4:j*4+4])
		})
		indexFloat := p * float64(len(pixelData)-1)
		index := int(indexFloat)
		if indexFloat == float64(index) {
			return float64(binary.LittleEndian.Uint32(pixelData[index*4 : index*4+4]))
		}
		value := float64(binary.LittleEndian.Uint32(pixelData[index*4 : index*4+4]))
		nextValue := float64(binary.LittleEndian.Uint32(pixelData[(index+1)*4 : (index+1)*4+4]))
		return value + (indexFloat-float64(index))*(nextValue-value)
	case PixelTypeInt32:
		sort.Slice(pixelData, func(i, j int) bool {
			return int32(binary.LittleEndian.Uint32(pixelData[i*4:i*4+4])) < int32(binary.LittleEndian.Uint32(pixelData[j*4:j*4+4]))
		})
		indexFloat := p * float64(len(pixelData)-1)
		index := int(indexFloat)
		if indexFloat == float64(index) {
			return float64(int32(binary.LittleEndian.Uint32(pixelData[index*4 : index*4+4])))
		}
		value := float64(int32(binary.LittleEndian.Uint32(pixelData[index*4 : index*4+4])))
		nextValue := float64(int32(binary.LittleEndian.Uint32(pixelData[(index+1)*4 : (index+1)*4+4])))
		return value + (indexFloat-float64(index))*(nextValue-value)
	case PixelTypeUInt64:
		sort.Slice(pixelData, func(i, j int) bool {
			return binary.LittleEndian.Uint64(pixelData[i*8:i*8+8]) < binary.LittleEndian.Uint64(pixelData[j*8:j*8+8])
		})
		indexFloat := p * float64(len(pixelData)-1)
		index := int(indexFloat)
		if indexFloat == float64(index) {
			return float64(binary.LittleEndian.Uint64(pixelData[index*8 : index*8+8]))
		}
		value := float64(binary.LittleEndian.Uint64(pixelData[index*8 : index*8+8]))
		nextValue := float64(binary.LittleEndian.Uint64(pixelData[(index+1)*8 : (index+1)*8+8]))
		return value + (indexFloat-float64(index))*(nextValue-value)
	case PixelTypeInt64:
		sort.Slice(pixelData, func(i, j int) bool {
			return int64(binary.LittleEndian.Uint64(pixelData[i*8:i*8+8])) < int64(binary.LittleEndian.Uint64(pixelData[j*8:j*8+8]))
		})
		indexFloat := p * float64(len(pixelData)-1)
		index := int(indexFloat)
		if indexFloat == float64(index) {
			return float64(int64(binary.LittleEndian.Uint64(pixelData[index*8 : index*8+8])))
		}
		value := float64(int64(binary.LittleEndian.Uint64(pixelData[index*8 : index*8+8])))
		nextValue := float64(int64(binary.LittleEndian.Uint64(pixelData[(index+1)*8 : (index+1)*8+8])))
		return value + (indexFloat-float64(index))*(nextValue-value)
	case PixelTypeFloat32:
		sort.Slice(pixelData, func(i, j int) bool {
			return math.Float32frombits(binary.LittleEndian.Uint32(pixelData[i*4:i*4+4])) < math.Float32frombits(binary.LittleEndian.Uint32(pixelData[j*4:j*4+4]))
		})
		indexFloat := p * float64(len(pixelData)-1)
		index := int(indexFloat)
		if indexFloat == float64(index) {
			return float64(math.Float32frombits(binary.LittleEndian.Uint32(pixelData[index*4 : index*4+4])))
		}
		value := float64(math.Float32frombits(binary.LittleEndian.Uint32(pixelData[index*4 : index*4+4])))
		nextValue := float64(math.Float32frombits(binary.LittleEndian.Uint32(pixelData[(index+1)*4 : (index+1)*4+4])))
		return value + (indexFloat-float64(index))*(nextValue-value)
	case PixelTypeFloat64:
		sort.Slice(pixelData, func(i, j int) bool {
			return math.Float64frombits(binary.LittleEndian.Uint64(pixelData[i*8:i*8+8])) < math.Float64frombits(binary.LittleEndian.Uint64(pixelData[j*8:j*8+8]))
		})
		indexFloat := p * float64(len(pixelData)-1)
		index := int(indexFloat)
		if indexFloat == float64(index) {
			return float64(math.Float64frombits(binary.LittleEndian.Uint64(pixelData[index*8 : index*8+8])))
		}
		value := float64(math.Float64frombits(binary.LittleEndian.Uint64(pixelData[index*8 : index*8+8])))
		nextValue := float64(math.Float64frombits(binary.LittleEndian.Uint64(pixelData[(index+2)*8 : (index+1)*8+8])))
		return value + (indexFloat-float64(index))*(nextValue-value)
	default:
		return 0
	}
}

// OtsuThreshold returns the threshold value for the Otsu thresholding method.
// Returns:
//   - float64: The threshold value.
func (img *Image) OtsuThreshold() float64 {
	var maxVal float64
	switch img.pixelType {
	case PixelTypeUInt8:
		maxVal = float64(img.Max().(uint8))
	case PixelTypeInt8:
		maxVal = float64(img.Max().(int8))
	case PixelTypeUInt16:
		maxVal = float64(img.Max().(uint16))
	case PixelTypeInt16:
		maxVal = float64(img.Max().(int16))
	case PixelTypeUInt32:
		maxVal = float64(img.Max().(uint32))
	case PixelTypeInt32:
		maxVal = float64(img.Max().(int32))
	case PixelTypeUInt64:
		maxVal = float64(img.Max().(uint64))
	case PixelTypeInt64:
		maxVal = float64(img.Max().(int64))
	case PixelTypeFloat32:
		maxVal = float64(img.Max().(float32))
	case PixelTypeFloat64:
		maxVal = float64(img.Max().(float64))
	default:
		return 0
	}

	size := img.GetSize()
	hist := make([]int, 256)
	switch img.dimension {
	case 2:
		for y := 0; y < int(size[1]); y++ {
			for x := 0; x < int(size[0]); x++ {
				value, err := img.GetPixelAsFloat64([]uint32{uint32(x), uint32(y)})
				if err != nil {
					return 0
				}
				hist[int(value/maxVal*255)]++
			}
		}
	case 3:
		for z := 0; z < int(size[2]); z++ {
			for y := 0; y < int(size[1]); y++ {
				for x := 0; x < int(size[0]); x++ {
					value, err := img.GetPixelAsFloat64([]uint32{uint32(x), uint32(y), uint32(z)})
					if err != nil {
						return 0
					}
					hist[int(value/maxVal*255)]++
				}
			}
		}
	}

	total := 1
	for _, v := range size {
		total *= int(v)
	}

	sum := 0
	for i := 0; i < 256; i++ {
		sum += hist[i] * i
	}

	sumB, wB, wF, varMax, threshold := 0, 0, 0, 0.0, 0
	for t := 0; t < 256; t++ {
		wB += hist[t]
		if wB == 0 {
			continue
		}
		wF = total - wB
		if wF == 0 {
			break
		}
		sumB += t * hist[t]
		mB := float64(sumB) / float64(wB)
		mF := float64(sum-sumB) / float64(wF)
		varBetween := float64(wB) * float64(wF) * (mB - mF) * (mB - mF)
		if varBetween > varMax {
			varMax = varBetween
			threshold = t
		}
	}
	return float64(threshold) / 255.0 * maxVal
}
