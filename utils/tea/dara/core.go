package dara

import "reflect"

func IsNil(val interface{}) bool {
	defer func() {
		recover()
	}()
	if val == nil {
		return true
	}

	v := reflect.ValueOf(val)
	if v.Kind() == reflect.Ptr || v.Kind() == reflect.Slice || v.Kind() == reflect.Map {
		return v.IsNil()
	}

	valType := reflect.TypeOf(val)
	valZero := reflect.Zero(valType)
	return valZero == v
}

func isNil(a interface{}) bool {
	defer func() {
		recover()
	}()
	vi := reflect.ValueOf(a)
	return vi.IsNil()
}

func isNilOrZero(value interface{}) bool {
	if value == nil {
		return true
	}

	v := reflect.ValueOf(value)
	switch v.Kind() {
	case reflect.Chan, reflect.Func, reflect.Map, reflect.Ptr, reflect.Interface, reflect.Slice:
		return v.IsNil()
	default:
		// Check for zero value
		return reflect.DeepEqual(value, reflect.Zero(v.Type()).Interface())
	}
}
