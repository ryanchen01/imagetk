package imagetk

import (
	"reflect"
	"testing"
)

func extractValues(values []reflect.Value) []interface{} {
	result := make([]interface{}, len(values))
	for i, v := range values {
		result[i] = v.Interface()
	}
	return result
}

func TestFlatten(t *testing.T) {
	tests := []struct {
		name   string
		input  any
		expect []reflect.Value
	}{
		{name: "2D array", input: [][]int{{1, 2, 3}, {4, 5, 6}}, expect: []reflect.Value{reflect.ValueOf(1), reflect.ValueOf(2), reflect.ValueOf(3), reflect.ValueOf(4), reflect.ValueOf(5), reflect.ValueOf(6)}},
		{name: "3D array", input: [][][]int{{{1, 2, 3}, {4, 5, 6}}, {{7, 8, 9}, {10, 11, 12}}}, expect: []reflect.Value{reflect.ValueOf(1), reflect.ValueOf(2), reflect.ValueOf(3), reflect.ValueOf(4), reflect.ValueOf(5), reflect.ValueOf(6), reflect.ValueOf(7), reflect.ValueOf(8), reflect.ValueOf(9), reflect.ValueOf(10), reflect.ValueOf(11), reflect.ValueOf(12)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var elemType reflect.Type
			result := flatten(tt.input, &elemType)
			if !reflect.DeepEqual(extractValues(result), extractValues(tt.expect)) {
				t.Errorf("Expected %v, got %v", extractValues(tt.expect), extractValues(result))
			}
		})
	}
}
func TestGetValueAsPixelType(t *testing.T) {
	tests := []struct {
		name      string
		value     any
		pixelType int
		expect    any
		expectErr bool
	}{
		{name: "uint8 to uint8", value: uint8(255), pixelType: PixelTypeUInt8, expect: uint8(255), expectErr: false},
		{name: "int8 to int8", value: int8(-128), pixelType: PixelTypeInt8, expect: int8(-128), expectErr: false},
		{name: "uint16 to uint16", value: uint16(65535), pixelType: PixelTypeUInt16, expect: uint16(65535), expectErr: false},
		{name: "int16 to int16", value: int16(-32768), pixelType: PixelTypeInt16, expect: int16(-32768), expectErr: false},
		{name: "uint32 to uint32", value: uint32(4294967295), pixelType: PixelTypeUInt32, expect: uint32(4294967295), expectErr: false},
		{name: "int32 to int32", value: int32(-2147483648), pixelType: PixelTypeInt32, expect: int32(-2147483648), expectErr: false},
		{name: "uint64 to uint64", value: uint64(18446744073709551615), pixelType: PixelTypeUInt64, expect: uint64(18446744073709551615), expectErr: false},
		{name: "int64 to int64", value: int64(-9223372036854775808), pixelType: PixelTypeInt64, expect: int64(-9223372036854775808), expectErr: false},
		{name: "float32 to float32", value: float32(3.14), pixelType: PixelTypeFloat32, expect: float32(3.14), expectErr: false},
		{name: "float64 to float64", value: float64(3.141592653589793), pixelType: PixelTypeFloat64, expect: float64(3.141592653589793), expectErr: false},
		{name: "unsupported type", value: "string", pixelType: PixelTypeUInt8, expect: nil, expectErr: true},
		{name: "unknown pixel type", value: uint8(255), pixelType: -1, expect: nil, expectErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := getValueAsPixelType(tt.value, tt.pixelType)
			if (err != nil) != tt.expectErr {
				t.Errorf("Expected error: %v, got: %v", tt.expectErr, err)
			}
			if !reflect.DeepEqual(result, tt.expect) {
				t.Errorf("Expected result: %v, got: %v", tt.expect, result)
			}
		})
	}
}
