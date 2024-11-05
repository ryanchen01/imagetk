package imagetk

import "reflect"

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
