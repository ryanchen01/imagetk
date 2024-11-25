package imagetk

import (
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
		for _, value := range img.pixels.([]uint8) {
			if value < minValue {
				minValue = value
			}
		}
		return minValue
	case PixelTypeInt8:
		minValue := int8(math.MaxInt8)
		for _, value := range img.pixels.([]int8) {
			if value < minValue {
				minValue = value
			}
		}
		return minValue
	case PixelTypeUInt16:
		minValue := uint16(math.MaxUint16)
		for _, value := range img.pixels.([]uint16) {
			if value < minValue {
				minValue = value
			}
		}
		return minValue
	case PixelTypeInt16:
		minValue := int16(math.MaxInt16)
		for _, value := range img.pixels.([]int16) {
			if value < minValue {
				minValue = value
			}
		}
		return minValue
	case PixelTypeUInt32:
		minValue := uint32(math.MaxUint32)
		for _, value := range img.pixels.([]uint32) {
			if value < minValue {
				minValue = value
			}
		}
		return minValue
	case PixelTypeInt32:
		minValue := int32(math.MaxInt32)
		for _, value := range img.pixels.([]int32) {
			if value < minValue {
				minValue = value
			}
		}
		return minValue
	case PixelTypeUInt64:
		minValue := uint64(math.MaxUint64)
		for _, value := range img.pixels.([]uint64) {
			if value < minValue {
				minValue = value
			}
		}
		return minValue
	case PixelTypeInt64:
		minValue := int64(math.MaxInt64)
		for _, value := range img.pixels.([]int64) {
			if value < minValue {
				minValue = value
			}
		}
		return minValue
	case PixelTypeFloat32:
		minValue := float32(math.MaxFloat32)
		for _, value := range img.pixels.([]float32) {
			if value < minValue {
				minValue = value
			}
		}
		return minValue
	case PixelTypeFloat64:
		minValue := float64(math.MaxFloat64)
		for _, value := range img.pixels.([]float64) {
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
		for _, value := range img.pixels.([]uint8) {
			if value > maxValue {
				maxValue = value
			}
		}
		return maxValue
	case PixelTypeInt8:
		maxValue := int8(math.MinInt8)
		for _, value := range img.pixels.([]int8) {
			if value > maxValue {
				maxValue = value
			}
		}
		return maxValue
	case PixelTypeUInt16:
		maxValue := uint16(0)
		for _, value := range img.pixels.([]uint16) {
			if value > maxValue {
				maxValue = value
			}
		}
		return maxValue
	case PixelTypeInt16:
		maxValue := int16(math.MinInt16)
		for _, value := range img.pixels.([]int16) {
			if value > maxValue {
				maxValue = value
			}
		}
		return maxValue
	case PixelTypeUInt32:
		maxValue := uint32(0)
		for _, value := range img.pixels.([]uint32) {
			if value > maxValue {
				maxValue = value
			}
		}
		return maxValue
	case PixelTypeInt32:
		maxValue := int32(math.MinInt32)
		for _, value := range img.pixels.([]int32) {
			if value > maxValue {
				maxValue = value
			}
		}
		return maxValue
	case PixelTypeUInt64:
		maxValue := uint64(0)
		for _, value := range img.pixels.([]uint64) {
			if value > maxValue {
				maxValue = value
			}
		}
		return maxValue
	case PixelTypeInt64:
		maxValue := int64(math.MinInt64)
		for _, value := range img.pixels.([]int64) {
			if value > maxValue {
				maxValue = value
			}
		}
		return maxValue
	case PixelTypeFloat32:
		maxValue := float32(-math.MaxFloat32)
		for _, value := range img.pixels.([]float32) {
			if value > maxValue {
				maxValue = value
			}
		}
		return maxValue
	case PixelTypeFloat64:
		maxValue := float64(-math.MaxFloat64)
		for _, value := range img.pixels.([]float64) {
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
		for _, value := range img.pixels.([]uint8) {
			sumValue += uint64(value)
		}
		return sumValue
	case PixelTypeInt8:
		sumValue := int64(0)
		for _, value := range img.pixels.([]int8) {
			sumValue += int64(value)
		}
		return sumValue
	case PixelTypeUInt16:
		sumValue := uint64(0)
		for _, value := range img.pixels.([]uint16) {
			sumValue += uint64(value)
		}
		return sumValue
	case PixelTypeInt16:
		sumValue := int64(0)
		for _, value := range img.pixels.([]int16) {
			sumValue += int64(value)
		}
		return sumValue
	case PixelTypeUInt32:
		sumValue := uint64(0)
		for _, value := range img.pixels.([]uint32) {
			sumValue += uint64(value)
		}
		return sumValue
	case PixelTypeInt32:
		sumValue := int64(0)
		for _, value := range img.pixels.([]int32) {
			sumValue += int64(value)
		}
		return sumValue
	case PixelTypeUInt64:
		sumValue := uint64(0)
		for _, value := range img.pixels.([]uint64) {
			sumValue += uint64(value)
		}
		return sumValue
	case PixelTypeInt64:
		sumValue := int64(0)
		for _, value := range img.pixels.([]int64) {
			sumValue += int64(value)
		}
		return sumValue
	case PixelTypeFloat32:
		sumValue := float64(0)
		for _, value := range img.pixels.([]float32) {
			sumValue += float64(value)
		}
		return sumValue
	case PixelTypeFloat64:
		sumValue := float64(0)
		for _, value := range img.pixels.([]float64) {
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
		for _, value := range img.pixels.([]uint8) {
			productValue *= uint64(value)
		}
		return productValue
	case PixelTypeInt8:
		productValue := int64(1)
		for _, value := range img.pixels.([]int8) {
			productValue *= int64(value)
		}
		return productValue
	case PixelTypeUInt16:
		productValue := uint64(1)
		for _, value := range img.pixels.([]uint16) {
			productValue *= uint64(value)
		}
		return productValue
	case PixelTypeInt16:
		productValue := int64(1)
		for _, value := range img.pixels.([]int16) {
			productValue *= int64(value)
		}
		return productValue
	case PixelTypeUInt32:
		productValue := uint64(1)
		for _, value := range img.pixels.([]uint32) {
			productValue *= uint64(value)
		}
		return productValue
	case PixelTypeInt32:
		productValue := int64(1)
		for _, value := range img.pixels.([]int32) {
			productValue *= int64(value)
		}
		return productValue
	case PixelTypeUInt64:
		productValue := uint64(1)
		for _, value := range img.pixels.([]uint64) {
			productValue *= uint64(value)
		}
		return productValue
	case PixelTypeInt64:
		productValue := int64(1)
		for _, value := range img.pixels.([]int64) {
			productValue *= int64(value)
		}
		return productValue
	case PixelTypeFloat32:
		productValue := float64(1)
		for _, value := range img.pixels.([]float32) {
			productValue *= float64(value)
		}
		return productValue
	case PixelTypeFloat64:
		productValue := float64(1)
		for _, value := range img.pixels.([]float64) {
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
	case PixelTypeUInt8:
		sumValue := uint64(0)
		for _, value := range img.pixels.([]uint8) {
			sumValue += uint64(value)
		}
		return float64(sumValue) / float64(len(img.pixels.([]uint8)))
	case PixelTypeInt8:
		sumValue := int64(0)
		for _, value := range img.pixels.([]int8) {
			sumValue += int64(value)
		}
		return float64(sumValue) / float64(len(img.pixels.([]int8)))
	case PixelTypeUInt16:
		sumValue := uint64(0)
		for _, value := range img.pixels.([]uint16) {
			sumValue += uint64(value)
		}
		return float64(sumValue) / float64(len(img.pixels.([]uint16)))
	case PixelTypeInt16:
		sumValue := int64(0)
		for _, value := range img.pixels.([]int16) {
			sumValue += int64(value)
		}
		return float64(sumValue) / float64(len(img.pixels.([]int16)))
	case PixelTypeUInt32:
		sumValue := uint64(0)
		for _, value := range img.pixels.([]uint32) {
			sumValue += uint64(value)
		}
		return float64(sumValue) / float64(len(img.pixels.([]uint32)))
	case PixelTypeInt32:
		sumValue := int64(0)
		for _, value := range img.pixels.([]int32) {
			sumValue += int64(value)
		}
		return float64(sumValue) / float64(len(img.pixels.([]int32)))
	case PixelTypeUInt64:
		sumValue := uint64(0)
		for _, value := range img.pixels.([]uint64) {
			sumValue += uint64(value)
		}
		return float64(sumValue) / float64(len(img.pixels.([]uint64)))
	case PixelTypeInt64:
		sumValue := int64(0)
		for _, value := range img.pixels.([]int64) {
			sumValue += int64(value)
		}
		return float64(sumValue) / float64(len(img.pixels.([]int64)))
	case PixelTypeFloat32:
		sumValue := float64(0)
		for _, value := range img.pixels.([]float32) {
			sumValue += float64(value)
		}
		return sumValue / float64(len(img.pixels.([]float32)))
	case PixelTypeFloat64:
		sumValue := float64(0)
		for _, value := range img.pixels.([]float64) {
			sumValue += float64(value)
		}
		return sumValue / float64(len(img.pixels.([]float64)))
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
		sumValue := uint8(0)
		for _, value := range img.pixels.([]uint8) {
			sumValue += value
		}
		return sumValue / uint8(len(img.pixels.([]uint8)))
	case PixelTypeInt8:
		sumValue := int8(0)
		for _, value := range img.pixels.([]int8) {
			sumValue += value
		}
		return sumValue / int8(len(img.pixels.([]int8)))
	case PixelTypeUInt16:
		sumValue := uint16(0)
		for _, value := range img.pixels.([]uint16) {
			sumValue += value
		}
		return sumValue / uint16(len(img.pixels.([]uint16)))
	case PixelTypeInt16:
		sumValue := int16(0)
		for _, value := range img.pixels.([]int16) {
			sumValue += value
		}
		return sumValue / int16(len(img.pixels.([]int16)))
	case PixelTypeUInt32:
		sumValue := uint32(0)
		for _, value := range img.pixels.([]uint32) {
			sumValue += value
		}
		return sumValue / uint32(len(img.pixels.([]uint32)))
	case PixelTypeInt32:
		sumValue := int32(0)
		for _, value := range img.pixels.([]int32) {
			sumValue += value
		}
		return sumValue / int32(len(img.pixels.([]int32)))
	case PixelTypeUInt64:
		sumValue := uint64(0)
		for _, value := range img.pixels.([]uint64) {
			sumValue += value
		}
		return sumValue / uint64(len(img.pixels.([]uint64)))
	case PixelTypeInt64:
		sumValue := int64(0)
		for _, value := range img.pixels.([]int64) {
			sumValue += value
		}
		return sumValue / int64(len(img.pixels.([]int64)))
	case PixelTypeFloat32:
		sumValue := float32(0)
		for _, value := range img.pixels.([]float32) {
			sumValue += value
		}
		return sumValue / float32(len(img.pixels.([]float32)))
	case PixelTypeFloat64:
		sumValue := float64(0)
		for _, value := range img.pixels.([]float64) {
			sumValue += value
		}
		return sumValue / float64(len(img.pixels.([]float64)))
	default:
		return nil
	}
}

// Median returns the median of the image.
//
// Returns:
//   - float64: The median of the image.
func (img *Image) Median() float64 {
	switch img.pixelType {
	case PixelTypeUInt8:
		sort.Slice(img.pixels.([]uint8), func(i, j int) bool { return img.pixels.([]uint8)[i] < img.pixels.([]uint8)[j] })
		if len(img.pixels.([]uint8))%2 == 0 {
			return float64(img.pixels.([]uint8)[len(img.pixels.([]uint8))/2-1]+img.pixels.([]uint8)[len(img.pixels.([]uint8))/2]) / 2
		}
		return float64(img.pixels.([]uint8)[len(img.pixels.([]uint8))/2])
	case PixelTypeInt8:
		sort.Slice(img.pixels.([]int8), func(i, j int) bool { return img.pixels.([]int8)[i] < img.pixels.([]int8)[j] })
		if len(img.pixels.([]int8))%2 == 0 {
			return float64(img.pixels.([]int8)[len(img.pixels.([]int8))/2-1]+img.pixels.([]int8)[len(img.pixels.([]int8))/2]) / 2
		}
		return float64(img.pixels.([]int8)[len(img.pixels.([]int8))/2])
	case PixelTypeUInt16:
		sort.Slice(img.pixels.([]uint16), func(i, j int) bool { return img.pixels.([]uint16)[i] < img.pixels.([]uint16)[j] })
		if len(img.pixels.([]uint16))%2 == 0 {
			return float64(img.pixels.([]uint16)[len(img.pixels.([]uint16))/2-1]+img.pixels.([]uint16)[len(img.pixels.([]uint16))/2]) / 2
		}
		return float64(img.pixels.([]uint16)[len(img.pixels.([]uint16))/2])
	case PixelTypeInt16:
		sort.Slice(img.pixels.([]int16), func(i, j int) bool { return img.pixels.([]int16)[i] < img.pixels.([]int16)[j] })
		if len(img.pixels.([]int16))%2 == 0 {
			return float64(img.pixels.([]int16)[len(img.pixels.([]int16))/2-1]+img.pixels.([]int16)[len(img.pixels.([]int16))/2]) / 2
		}
		return float64(img.pixels.([]int16)[len(img.pixels.([]int16))/2])
	case PixelTypeUInt32:
		sort.Slice(img.pixels.([]uint32), func(i, j int) bool { return img.pixels.([]uint32)[i] < img.pixels.([]uint32)[j] })
		if len(img.pixels.([]uint32))%2 == 0 {
			return float64(img.pixels.([]uint32)[len(img.pixels.([]uint32))/2-1]+img.pixels.([]uint32)[len(img.pixels.([]uint32))/2]) / 2
		}
		return float64(img.pixels.([]uint32)[len(img.pixels.([]uint32))/2])
	case PixelTypeInt32:
		sort.Slice(img.pixels.([]int32), func(i, j int) bool { return img.pixels.([]int32)[i] < img.pixels.([]int32)[j] })
		if len(img.pixels.([]int32))%2 == 0 {
			return float64(img.pixels.([]int32)[len(img.pixels.([]int32))/2-1]+img.pixels.([]int32)[len(img.pixels.([]int32))/2]) / 2
		}
		return float64(img.pixels.([]int32)[len(img.pixels.([]int32))/2])
	case PixelTypeUInt64:
		sort.Slice(img.pixels.([]uint64), func(i, j int) bool { return img.pixels.([]uint64)[i] < img.pixels.([]uint64)[j] })
		if len(img.pixels.([]uint64))%2 == 0 {
			return float64(img.pixels.([]uint64)[len(img.pixels.([]uint64))/2-1]+img.pixels.([]uint64)[len(img.pixels.([]uint64))/2]) / 2
		}
		return float64(img.pixels.([]uint64)[len(img.pixels.([]uint64))/2])
	case PixelTypeInt64:
		sort.Slice(img.pixels.([]int64), func(i, j int) bool { return img.pixels.([]int64)[i] < img.pixels.([]int64)[j] })
		if len(img.pixels.([]int64))%2 == 0 {
			return float64(img.pixels.([]int64)[len(img.pixels.([]int64))/2-1]+img.pixels.([]int64)[len(img.pixels.([]int64))/2]) / 2
		}
		return float64(img.pixels.([]int64)[len(img.pixels.([]int64))/2])
	case PixelTypeFloat32:
		sort.Slice(img.pixels.([]float32), func(i, j int) bool { return img.pixels.([]float32)[i] < img.pixels.([]float32)[j] })
		if len(img.pixels.([]float32))%2 == 0 {
			return float64(img.pixels.([]float32)[len(img.pixels.([]float32))/2-1]+img.pixels.([]float32)[len(img.pixels.([]float32))/2]) / 2
		}
		return float64(img.pixels.([]float32)[len(img.pixels.([]float32))/2])
	case PixelTypeFloat64:
		sort.Slice(img.pixels.([]float64), func(i, j int) bool { return img.pixels.([]float64)[i] < img.pixels.([]float64)[j] })
		if len(img.pixels.([]float64))%2 == 0 {
			return (img.pixels.([]float64)[len(img.pixels.([]float64))/2-1] + img.pixels.([]float64)[len(img.pixels.([]float64))/2]) / 2
		}
		return float64(img.pixels.([]float64)[len(img.pixels.([]float64))/2])
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
		for _, v := range img.pixels.([]uint8) {
			value := float64(v)
			sumValue += (value - meanValue) * (value - meanValue)
		}
		return math.Sqrt(sumValue / float64(len(img.pixels.([]uint8))))
	case PixelTypeInt8:
		meanValue := img.ExactMean()
		sumValue := 0.0
		for _, v := range img.pixels.([]int8) {
			value := float64(v)
			sumValue += (value - meanValue) * (value - meanValue)
		}
		return math.Sqrt(sumValue / float64(len(img.pixels.([]int8))))
	case PixelTypeUInt16:
		meanValue := img.ExactMean()
		sumValue := 0.0
		for _, v := range img.pixels.([]uint16) {
			value := float64(v)
			sumValue += float64(value-meanValue) * float64(value-meanValue)
		}
		return math.Sqrt(float64(sumValue) / float64(len(img.pixels.([]uint16))))
	case PixelTypeInt16:
		meanValue := img.ExactMean()
		sumValue := 0.0
		for _, v := range img.pixels.([]int16) {
			value := float64(v)
			sumValue += float64(value-meanValue) * float64(value-meanValue)
		}
		return math.Sqrt(float64(sumValue) / float64(len(img.pixels.([]int16))))
	case PixelTypeUInt32:
		meanValue := img.ExactMean()
		sumValue := 0.0
		for _, v := range img.pixels.([]uint32) {
			value := float64(v)
			sumValue += (value - meanValue) * (value - meanValue)
		}
		return math.Sqrt(sumValue / float64(len(img.pixels.([]uint32))))
	case PixelTypeInt32:
		meanValue := img.ExactMean()
		sumValue := 0.0
		for _, v := range img.pixels.([]int32) {
			value := float64(v)
			sumValue += (value - meanValue) * (value - meanValue)
		}
		return math.Sqrt(sumValue / float64(len(img.pixels.([]int32))))
	case PixelTypeUInt64:
		meanValue := img.ExactMean()
		sumValue := 0.0
		for _, v := range img.pixels.([]uint64) {
			value := float64(v)
			sumValue += (value - meanValue) * (value - meanValue)
		}
		return math.Sqrt(sumValue / float64(len(img.pixels.([]uint64))))
	case PixelTypeInt64:
		meanValue := img.ExactMean()
		sumValue := 0.0
		for _, v := range img.pixels.([]int64) {
			value := float64(v)
			sumValue += (value - meanValue) * (value - meanValue)
		}
		return math.Sqrt(sumValue / float64(len(img.pixels.([]int64))))
	case PixelTypeFloat32:
		meanValue := img.ExactMean()
		sumValue := 0.0
		for _, v := range img.pixels.([]float32) {
			value := float64(v)
			sumValue += (value - meanValue) * (value - meanValue)
		}
		return math.Sqrt(sumValue / float64(len(img.pixels.([]float32))))
	case PixelTypeFloat64:
		meanValue := img.ExactMean()
		sumValue := 0.0
		for _, v := range img.pixels.([]float64) {
			value := float64(v)
			sumValue += (value - meanValue) * (value - meanValue)
		}
		return math.Sqrt(sumValue / float64(len(img.pixels.([]float64))))
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
	switch img.pixelType {
	case PixelTypeUInt8:
		pixelData := make([]uint8, len(img.pixels.([]uint8)))
		copy(pixelData, img.pixels.([]uint8))
		sort.Slice(pixelData, func(i, j int) bool { return pixelData[i] < pixelData[j] })
		indexFloat := p * float64(len(pixelData)-1)
		index := int(indexFloat)
		if indexFloat == float64(index) {
			return float64(img.pixels.([]uint8)[index])
		}
		value := float64(img.pixels.([]uint8)[index])
		nextValue := float64(img.pixels.([]uint8)[index+1])
		return value + (indexFloat-float64(index))*(nextValue-value)
	case PixelTypeInt8:
		pixelData := make([]int8, len(img.pixels.([]int8)))
		copy(pixelData, img.pixels.([]int8))
		sort.Slice(pixelData, func(i, j int) bool { return pixelData[i] < pixelData[j] })
		indexFloat := p * float64(len(pixelData)-1)
		index := int(indexFloat)
		if indexFloat == float64(index) {
			return float64(img.pixels.([]int8)[index])
		}
		value := float64(img.pixels.([]int8)[index])
		nextValue := float64(img.pixels.([]int8)[index+1])
		return value + (indexFloat-float64(index))*(nextValue-value)
	case PixelTypeUInt16:
		pixelData := make([]uint16, len(img.pixels.([]uint16)))
		copy(pixelData, img.pixels.([]uint16))
		sort.Slice(pixelData, func(i, j int) bool { return pixelData[i] < pixelData[j] })
		indexFloat := p * float64(len(pixelData)-1)
		index := int(indexFloat)
		if indexFloat == float64(index) {
			return float64(img.pixels.([]uint16)[index])
		}
		value := float64(img.pixels.([]uint16)[index])
		nextValue := float64(img.pixels.([]uint16)[index+1])
		return value + (indexFloat-float64(index))*(nextValue-value)
	case PixelTypeInt16:
		pixelData := make([]int16, len(img.pixels.([]int16)))
		copy(pixelData, img.pixels.([]int16))
		sort.Slice(pixelData, func(i, j int) bool { return pixelData[i] < pixelData[j] })
		indexFloat := p * float64(len(pixelData)-1)
		index := int(indexFloat)
		if indexFloat == float64(index) {
			return float64(img.pixels.([]int16)[index])
		}
		value := float64(img.pixels.([]int16)[index])
		nextValue := float64(img.pixels.([]int16)[index+1])
		return value + (indexFloat-float64(index))*(nextValue-value)
	case PixelTypeUInt32:
		pixelData := make([]uint32, len(img.pixels.([]uint32)))
		copy(pixelData, img.pixels.([]uint32))
		sort.Slice(pixelData, func(i, j int) bool { return pixelData[i] < pixelData[j] })
		indexFloat := p * float64(len(pixelData)-1)
		index := int(indexFloat)
		if indexFloat == float64(index) {
			return float64(img.pixels.([]uint32)[index])
		}
		value := float64(img.pixels.([]uint32)[index])
		nextValue := float64(img.pixels.([]uint32)[index+1])
		return value + (indexFloat-float64(index))*(nextValue-value)
	case PixelTypeInt32:
		pixelData := make([]int32, len(img.pixels.([]int32)))
		copy(pixelData, img.pixels.([]int32))
		sort.Slice(pixelData, func(i, j int) bool { return pixelData[i] < pixelData[j] })
		indexFloat := p * float64(len(pixelData)-1)
		index := int(indexFloat)
		if indexFloat == float64(index) {
			return float64(img.pixels.([]int32)[index])
		}
		value := float64(img.pixels.([]int32)[index])
		nextValue := float64(img.pixels.([]int32)[index+1])
		return value + (indexFloat-float64(index))*(nextValue-value)
	case PixelTypeUInt64:
		pixelData := make([]uint64, len(img.pixels.([]uint64)))
		copy(pixelData, img.pixels.([]uint64))
		sort.Slice(pixelData, func(i, j int) bool { return pixelData[i] < pixelData[j] })
		indexFloat := p * float64(len(pixelData)-1)
		index := int(indexFloat)
		if indexFloat == float64(index) {
			return float64(img.pixels.([]uint64)[index])
		}
		value := float64(img.pixels.([]uint64)[index])
		nextValue := float64(img.pixels.([]uint64)[index+1])
		return value + (indexFloat-float64(index))*(nextValue-value)
	case PixelTypeInt64:
		pixelData := make([]int64, len(img.pixels.([]int64)))
		copy(pixelData, img.pixels.([]int64))
		sort.Slice(pixelData, func(i, j int) bool { return pixelData[i] < pixelData[j] })
		indexFloat := p * float64(len(pixelData)-1)
		index := int(indexFloat)
		if indexFloat == float64(index) {
			return float64(img.pixels.([]int64)[index])
		}
		value := float64(img.pixels.([]int64)[index])
		nextValue := float64(img.pixels.([]int64)[index+1])
		return value + (indexFloat-float64(index))*(nextValue-value)
	case PixelTypeFloat32:
		pixelData := make([]float32, len(img.pixels.([]float32)))
		copy(pixelData, img.pixels.([]float32))
		sort.Slice(pixelData, func(i, j int) bool { return pixelData[i] < pixelData[j] })
		indexFloat := p * float64(len(pixelData)-1)
		index := int(indexFloat)
		if indexFloat == float64(index) {
			return float64(img.pixels.([]float32)[index])
		}
		value := float64(img.pixels.([]float32)[index])
		nextValue := float64(img.pixels.([]float32)[index+1])
		return value + (indexFloat-float64(index))*(nextValue-value)
	case PixelTypeFloat64:
		pixelData := make([]float64, len(img.pixels.([]float64)))
		copy(pixelData, img.pixels.([]float64))
		sort.Slice(pixelData, func(i, j int) bool { return pixelData[i] < pixelData[j] })
		indexFloat := p * float64(len(pixelData)-1)
		index := int(indexFloat)
		if indexFloat == float64(index) {
			return float64(img.pixels.([]float64)[index])
		}
		value := float64(img.pixels.([]float64)[index])
		nextValue := float64(img.pixels.([]float64)[index+1])
		return value + (indexFloat-float64(index))*(nextValue-value)
	default:
		return 0
	}
}

// OtsuThreshold returns the threshold value for the Otsu thresholding method.
// Returns:
// - float64: The threshold value.
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
