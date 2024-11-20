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
