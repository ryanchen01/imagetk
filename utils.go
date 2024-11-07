package imagetk

import (
	"fmt"
	"reflect"
)

func flattenReflectValues(data any) []reflect.Value {
	var flattenedValues []reflect.Value
	var helper func(interface{})

	helper = func(d interface{}) {
		val := reflect.ValueOf(d)
		switch val.Kind() {
		case reflect.Slice:
			for i := 0; i < val.Len(); i++ {
				helper(val.Index(i).Interface())
			}
		default:
			flattenedValues = append(flattenedValues, val)
		}
	}

	helper(data)

	return flattenedValues
}

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
