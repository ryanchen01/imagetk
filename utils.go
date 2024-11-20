package imagetk

import (
	"fmt"
	"reflect"
)

func getValueAsPixelType(value any, pixelType int) (any, error) {
	switch value := value.(type) {
	case uint8:
		switch pixelType {
		case PixelTypeUInt8:
			return value, nil
		case PixelTypeInt8:
			return int8(value), nil
		case PixelTypeUInt16:
			return uint16(value), nil
		case PixelTypeInt16:
			return int16(value), nil
		case PixelTypeUInt32:
			return uint32(value), nil
		case PixelTypeInt32:
			return int32(value), nil
		case PixelTypeUInt64:
			return uint64(value), nil
		case PixelTypeInt64:
			return int64(value), nil
		case PixelTypeFloat32:
			return float32(value), nil
		case PixelTypeFloat64:
			return float64(value), nil
		default:
			return nil, fmt.Errorf("unknown pixel type")
		}
	case int8:
		switch pixelType {
		case PixelTypeUInt8:
			return uint8(value), nil
		case PixelTypeInt8:
			return value, nil
		case PixelTypeUInt16:
			return uint16(value), nil
		case PixelTypeInt16:
			return int16(value), nil
		case PixelTypeUInt32:
			return uint32(value), nil
		case PixelTypeInt32:
			return int32(value), nil
		case PixelTypeUInt64:
			return uint64(value), nil
		case PixelTypeInt64:
			return int64(value), nil
		case PixelTypeFloat32:
			return float32(value), nil
		case PixelTypeFloat64:
			return float64(value), nil
		default:
			return nil, fmt.Errorf("unknown pixel type")
		}
	case uint16:
		switch pixelType {
		case PixelTypeUInt8:
			return uint8(value), nil
		case PixelTypeInt8:
			return int8(value), nil
		case PixelTypeUInt16:
			return value, nil
		case PixelTypeInt16:
			return int16(value), nil
		case PixelTypeUInt32:
			return uint32(value), nil
		case PixelTypeInt32:
			return int32(value), nil
		case PixelTypeUInt64:
			return uint64(value), nil
		case PixelTypeInt64:
			return int64(value), nil
		case PixelTypeFloat32:
			return float32(value), nil
		case PixelTypeFloat64:
			return float64(value), nil
		default:
			return nil, fmt.Errorf("unknown pixel type")
		}
	case int16:
		switch pixelType {
		case PixelTypeUInt8:
			return uint8(value), nil
		case PixelTypeInt8:
			return int8(value), nil
		case PixelTypeUInt16:
			return uint16(value), nil
		case PixelTypeInt16:
			return value, nil
		case PixelTypeUInt32:
			return uint32(value), nil
		case PixelTypeInt32:
			return int32(value), nil
		case PixelTypeUInt64:
			return uint64(value), nil
		case PixelTypeInt64:
			return int64(value), nil
		case PixelTypeFloat32:
			return float32(value), nil
		case PixelTypeFloat64:
			return float64(value), nil
		default:
			return nil, fmt.Errorf("unknown pixel type")
		}
	case uint32:
		switch pixelType {
		case PixelTypeUInt8:
			return uint8(value), nil
		case PixelTypeInt8:
			return int8(value), nil
		case PixelTypeUInt16:
			return uint16(value), nil
		case PixelTypeInt16:
			return int16(value), nil
		case PixelTypeUInt32:
			return value, nil
		case PixelTypeInt32:
			return int32(value), nil
		case PixelTypeUInt64:
			return uint64(value), nil
		case PixelTypeInt64:
			return int64(value), nil
		case PixelTypeFloat32:
			return float32(value), nil
		case PixelTypeFloat64:
			return float64(value), nil
		default:
			return nil, fmt.Errorf("unknown pixel type")
		}
	case int32:
		switch pixelType {
		case PixelTypeUInt8:
			return uint8(value), nil
		case PixelTypeInt8:
			return int8(value), nil
		case PixelTypeUInt16:
			return uint16(value), nil
		case PixelTypeInt16:
			return int16(value), nil
		case PixelTypeUInt32:
			return uint32(value), nil
		case PixelTypeInt32:
			return value, nil
		case PixelTypeUInt64:
			return uint64(value), nil
		case PixelTypeInt64:
			return int64(value), nil
		case PixelTypeFloat32:
			return float32(value), nil
		case PixelTypeFloat64:
			return float64(value), nil
		default:
			return nil, fmt.Errorf("unknown pixel type")
		}
	case uint64:
		switch pixelType {
		case PixelTypeUInt8:
			return uint8(value), nil
		case PixelTypeInt8:
			return int8(value), nil
		case PixelTypeUInt16:
			return uint16(value), nil
		case PixelTypeInt16:
			return int16(value), nil
		case PixelTypeUInt32:
			return uint32(value), nil
		case PixelTypeInt32:
			return int32(value), nil
		case PixelTypeUInt64:
			return value, nil
		case PixelTypeInt64:
			return int64(value), nil
		case PixelTypeFloat32:
			return float32(value), nil
		case PixelTypeFloat64:
			return float64(value), nil
		default:
			return nil, fmt.Errorf("unknown pixel type")
		}
	case int64:
		switch pixelType {
		case PixelTypeUInt8:
			return uint8(value), nil
		case PixelTypeInt8:
			return int8(value), nil
		case PixelTypeUInt16:
			return uint16(value), nil
		case PixelTypeInt16:
			return int16(value), nil
		case PixelTypeUInt32:
			return uint32(value), nil
		case PixelTypeInt32:
			return int32(value), nil
		case PixelTypeUInt64:
			return uint64(value), nil
		case PixelTypeInt64:
			return value, nil
		case PixelTypeFloat32:
			return float32(value), nil
		case PixelTypeFloat64:
			return float64(value), nil
		default:
			return nil, fmt.Errorf("unknown pixel type")
		}
	case float32:
		switch pixelType {
		case PixelTypeUInt8:
			return uint8(value), nil
		case PixelTypeInt8:
			return int8(value), nil
		case PixelTypeUInt16:
			return uint16(value), nil
		case PixelTypeInt16:
			return int16(value), nil
		case PixelTypeUInt32:
			return uint32(value), nil
		case PixelTypeInt32:
			return int32(value), nil
		case PixelTypeUInt64:
			return uint64(value), nil
		case PixelTypeInt64:
			return int64(value), nil
		case PixelTypeFloat32:
			return value, nil
		case PixelTypeFloat64:
			return float64(value), nil
		default:
			return nil, fmt.Errorf("unknown pixel type")
		}
	case float64:
		switch pixelType {
		case PixelTypeUInt8:
			return uint8(value), nil
		case PixelTypeInt8:
			return int8(value), nil
		case PixelTypeUInt16:
			return uint16(value), nil
		case PixelTypeInt16:
			return int16(value), nil
		case PixelTypeUInt32:
			return uint32(value), nil
		case PixelTypeInt32:
			return int32(value), nil
		case PixelTypeUInt64:
			return uint64(value), nil
		case PixelTypeInt64:
			return int64(value), nil
		case PixelTypeFloat32:
			return float32(value), nil
		case PixelTypeFloat64:
			return value, nil
		default:
			return nil, fmt.Errorf("unknown pixel type")
		}
	case int:
		switch pixelType {
		case PixelTypeUInt8:
			return uint8(value), nil
		case PixelTypeInt8:
			return int8(value), nil
		case PixelTypeUInt16:
			return uint16(value), nil
		case PixelTypeInt16:
			return int16(value), nil
		case PixelTypeUInt32:
			return uint32(value), nil
		case PixelTypeInt32:
			return int32(value), nil
		case PixelTypeUInt64:
			return uint64(value), nil
		case PixelTypeInt64:
			return int64(value), nil
		case PixelTypeFloat32:
			return float32(value), nil
		case PixelTypeFloat64:
			return float64(value), nil
		default:
			return nil, fmt.Errorf("unknown pixel type")
		}
	default:
		return nil, fmt.Errorf("unsupported value type")
	}
}

// buildNestedSlice constructs an n-dimensional nested slice from the flat data according to the shape.
func buildNestedSlice(data []reflect.Value, shape []uint32, elemType reflect.Type) reflect.Value {
	var index int

	var build func(level int) reflect.Value
	build = func(level int) reflect.Value {
		if level == len(shape) {
			if index >= len(data) {
				panic("Insufficient data to fill the shape")
			}
			val := data[index]
			index++
			return val
		}

		// Create a slice for the current dimension.
		size := shape[level]
		sliceType := elemType
		for i := len(shape) - level - 1; i >= 0; i-- {
			sliceType = reflect.SliceOf(sliceType)
		}
		slice := reflect.MakeSlice(sliceType, int(size), int(size))

		// Recursively build nested slices.
		for i := 0; i < int(size); i++ {
			slice.Index(i).Set(build(level + 1))
		}
		return slice
	}

	return build(0)
}

// flatten converts an n-dimensional nested slice into a flat slice of reflect.Values.
// It also determines the element type of the innermost elements.
func flatten(data interface{}, elemType *reflect.Type) []reflect.Value {
	var flat []reflect.Value
	var helper func(interface{})

	helper = func(d interface{}) {
		val := reflect.ValueOf(d)
		switch val.Kind() {
		case reflect.Slice:
			for i := 0; i < val.Len(); i++ {
				helper(val.Index(i).Interface())
			}
		default:
			if *elemType == nil {
				*elemType = val.Type()
			}
			flat = append(flat, val)
		}
	}

	helper(data)
	return flat
}

// Reshape reshapes an n-dimensional nested slice into the specified shape.
func reshape(data interface{}, shape []uint32) any {
	if data == nil {
		return nil
	}
	// Flatten the input data and determine the element type.
	var elemType reflect.Type
	flatData := flatten(data, &elemType)

	// Calculate the total number of elements required by the new shape.
	totalSize := uint64(1)
	for _, dim := range shape {
		totalSize *= uint64(dim)
	}

	// Ensure the total number of elements matches.
	if uint64(len(flatData)) != totalSize {
		panic(fmt.Sprintf("Cannot reshape array of size %d into shape %v", len(flatData), shape))
	}

	// Build the reshaped nested slice.
	reshaped := buildNestedSlice(flatData, shape, elemType)

	return reshaped.Interface()
}
