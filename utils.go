package imagetk

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"reflect"
)

type converter func(interface{}) interface{}

var pixelTypeConverters = map[int]converter{
	PixelTypeUInt8: func(v interface{}) interface{} {
		return uint8(reflect.ValueOf(v).Convert(reflect.TypeOf(uint8(0))).Uint())
	},
	PixelTypeInt8: func(v interface{}) interface{} {
		return int8(reflect.ValueOf(v).Convert(reflect.TypeOf(int8(0))).Int())
	},
	PixelTypeUInt16: func(v interface{}) interface{} {
		return uint16(reflect.ValueOf(v).Convert(reflect.TypeOf(uint16(0))).Uint())
	},
	PixelTypeInt16: func(v interface{}) interface{} {
		return int16(reflect.ValueOf(v).Convert(reflect.TypeOf(int16(0))).Int())
	},
	PixelTypeUInt32: func(v interface{}) interface{} {
		return uint32(reflect.ValueOf(v).Convert(reflect.TypeOf(uint32(0))).Uint())
	},
	PixelTypeInt32: func(v interface{}) interface{} {
		return int32(reflect.ValueOf(v).Convert(reflect.TypeOf(int32(0))).Int())
	},
	PixelTypeUInt64: func(v interface{}) interface{} {
		return uint64(reflect.ValueOf(v).Convert(reflect.TypeOf(uint64(0))).Uint())
	},
	PixelTypeInt64: func(v interface{}) interface{} {
		return int64(reflect.ValueOf(v).Convert(reflect.TypeOf(int64(0))).Int())
	},
	PixelTypeFloat32: func(v interface{}) interface{} {
		return float32(reflect.ValueOf(v).Convert(reflect.TypeOf(float32(0))).Float())
	},
	PixelTypeFloat64: func(v interface{}) interface{} {
		return float64(reflect.ValueOf(v).Convert(reflect.TypeOf(float64(0))).Float())
	},
}

func getValueAsPixelType(value any, pixelType int) (any, error) {
	if converter, ok := pixelTypeConverters[pixelType]; ok {
		switch value.(type) {
		case uint8, int8, uint16, int16, uint32, int32, uint64, int64, float32, float64, int:
			return converter(value), nil
		default:
			return nil, fmt.Errorf("unsupported value type")
		}
	}
	return nil, fmt.Errorf("unknown pixel type")
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

func flattenToBytes(data interface{}) ([]byte, error) {
	var bytesSlice []byte
	var helper func(interface{}) error
	helper = func(d interface{}) error {
		val := reflect.ValueOf(d)
		switch val.Kind() {
		case reflect.Slice:
			for i := 0; i < val.Len(); i++ {
				helper(val.Index(i).Interface())
			}
		case reflect.Uint8:
			buf := new(bytes.Buffer)
			v := uint8(val.Uint())
			err := binary.Write(buf, binary.LittleEndian, v)
			if err != nil {
				panic(err)
			}
			bytesSlice = append(bytesSlice, buf.Bytes()...)
		case reflect.Int8:
			buf := new(bytes.Buffer)
			v := int8(val.Int())
			err := binary.Write(buf, binary.LittleEndian, v)
			if err != nil {
				panic(err)
			}
			bytesSlice = append(bytesSlice, buf.Bytes()...)
		case reflect.Uint16:
			buf := new(bytes.Buffer)
			v := uint16(val.Uint())
			err := binary.Write(buf, binary.LittleEndian, v)
			if err != nil {
				panic(err)
			}
			bytesSlice = append(bytesSlice, buf.Bytes()...)
		case reflect.Int16:
			buf := new(bytes.Buffer)
			v := int16(val.Int())
			err := binary.Write(buf, binary.LittleEndian, v)
			if err != nil {
				panic(err)
			}
			bytesSlice = append(bytesSlice, buf.Bytes()...)
		case reflect.Uint32:
			buf := new(bytes.Buffer)
			v := uint32(val.Uint())
			err := binary.Write(buf, binary.LittleEndian, v)
			if err != nil {
				panic(err)
			}
			bytesSlice = append(bytesSlice, buf.Bytes()...)
		case reflect.Int32:
			buf := new(bytes.Buffer)
			v := int32(val.Int())
			err := binary.Write(buf, binary.LittleEndian, v)
			if err != nil {
				panic(err)
			}
			bytesSlice = append(bytesSlice, buf.Bytes()...)
		case reflect.Uint64:
			buf := new(bytes.Buffer)
			v := uint64(val.Uint())
			err := binary.Write(buf, binary.LittleEndian, v)
			if err != nil {
				panic(err)
			}
			bytesSlice = append(bytesSlice, buf.Bytes()...)
		case reflect.Int64:
			buf := new(bytes.Buffer)
			v := int64(val.Int())
			err := binary.Write(buf, binary.LittleEndian, v)
			if err != nil {
				panic(err)
			}
			bytesSlice = append(bytesSlice, buf.Bytes()...)
		case reflect.Float32:
			buf := new(bytes.Buffer)
			v := float32(val.Float()) // float32
			err := binary.Write(buf, binary.LittleEndian, v)
			if err != nil {
				panic(err)
			}
			bytesSlice = append(bytesSlice, buf.Bytes()...)
		case reflect.Float64:
			buf := new(bytes.Buffer)
			v := float64(val.Float()) // float64
			err := binary.Write(buf, binary.LittleEndian, v)
			if err != nil {
				panic(err)
			}
			bytesSlice = append(bytesSlice, buf.Bytes()...)
		default:
			return fmt.Errorf("unsupported value type")
		}
		return nil
	}
	err := helper(data)
	return bytesSlice, err
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

func getValueAsBytes(value any) ([]byte, error) {
	switch value := value.(type) {
	case uint8:
		buf := new(bytes.Buffer)
		err := binary.Write(buf, binary.LittleEndian, value)
		if err != nil {
			panic(err)
		}
		return buf.Bytes(), nil
	case int8:
		buf := new(bytes.Buffer)
		err := binary.Write(buf, binary.LittleEndian, value)
		if err != nil {
			panic(err)
		}
		return buf.Bytes(), nil
	case uint16:
		buf := new(bytes.Buffer)
		err := binary.Write(buf, binary.LittleEndian, value)
		if err != nil {
			panic(err)
		}
		return buf.Bytes(), nil
	case int16:
		buf := new(bytes.Buffer)
		err := binary.Write(buf, binary.LittleEndian, value)
		if err != nil {
			panic(err)
		}
		return buf.Bytes(), nil
	case uint32:
		buf := new(bytes.Buffer)
		err := binary.Write(buf, binary.LittleEndian, value)
		if err != nil {
			panic(err)
		}
		return buf.Bytes(), nil
	case int32:
		buf := new(bytes.Buffer)
		err := binary.Write(buf, binary.LittleEndian, value)
		if err != nil {
			panic(err)
		}
		return buf.Bytes(), nil
	case uint64:
		buf := new(bytes.Buffer)
		err := binary.Write(buf, binary.LittleEndian, value)
		if err != nil {
			panic(err)
		}
		return buf.Bytes(), nil
	case int64:
		buf := new(bytes.Buffer)
		err := binary.Write(buf, binary.LittleEndian, value)
		if err != nil {
			panic(err)
		}
		return buf.Bytes(), nil
	case float32:
		buf := new(bytes.Buffer)
		err := binary.Write(buf, binary.LittleEndian, value)
		if err != nil {
			panic(err)
		}
		return buf.Bytes(), nil
	case float64:
		buf := new(bytes.Buffer)
		err := binary.Write(buf, binary.LittleEndian, value)
		if err != nil {
			panic(err)
		}
		return buf.Bytes(), nil
	default:
		return nil, fmt.Errorf("unsupported value type")
	}

}
