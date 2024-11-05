package imagetk

import "math"

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

func (img *Image) ExactMean() any {
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
		return nil
	}
}

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
