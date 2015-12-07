package collections

import (
	"reflect"
	"strings"

	"github.com/blendlabs/go-exception"
)

// An interface that provides the mechanism for a type to be sorted
type Comparable interface {
	// CompareTo should return -1 if it is less than other, 0 if equal to other, and 1 if greater than other
	CompareTo(other interface{}) (int, error)
}

// Comparer should return -1 if it is less than other, 0 if equal to other, and 1 if greater than other
type Comparer func(this, that interface{}) (int, error)

func int16Comparer(this, that interface{}) (int, error) {
	thisTyped, thisTypedErr := castAsInt16(this)
	if thisTypedErr != nil {
		return 0, thisTypedErr
	}

	thatTyped, thatTypedErr := castAsInt16(that)
	if thatTypedErr != nil {
		return 0, thatTypedErr
	}

	if thisTyped < thatTyped {
		return -1
	} else if thisTyped > thatTyped {
		return 1
	}

	return 0
}

func intComparer(this, that interface{}) (int, error) {
	thisTyped, thisTypedErr := castAsInt(this)
	if thisTypedErr != nil {
		return 0, thisTypedErr
	}

	thatTyped, thatTypedErr := castAsInt(that)
	if thatTypedErr != nil {
		return 0, thatTypedErr
	}

	if thisTyped < thatTyped {
		return -1
	} else if thisTyped > thatTyped {
		return 1
	}

	return 0
}

func int64Comparer(this, that interface{}) (int, error) {
	thisTyped, thisTypedErr := castAsInt64(this)
	if thisTypedErr != nil {
		return 0, thisTypedErr
	}

	thatTyped, thatTypedErr := castAsInt64(that)
	if thatTypedErr != nil {
		return 0, thatTypedErr
	}

	if thisTyped < thatTyped {
		return -1
	} else if thisTyped > thatTyped {
		return 1
	}

	return 0
}

func float32Comparer(this, that interface{}) (int, error) {
	thisTyped, thisTypedErr := castAsFloat32(this)
	if thisTypedErr != nil {
		return 0, thisTypedErr
	}

	thatTyped, thatTypedErr := castAsFloat32(that)
	if thatTypedErr != nil {
		return 0, thatTypedErr
	}

	if thisTyped < thatTyped {
		return -1
	} else if thisTyped > thatTyped {
		return 1
	}

	return 0
}

func float64Comparer(this, that interface{}) (int, error) {
	thisTyped, thisTypedErr := castAsFloat64(this)
	if thisTypedErr != nil {
		return 0, thisTypedErr
	}

	thatTyped, thatTypedErr := castAsFloat64(that)
	if thatTypedErr != nil {
		return 0, thatTypedErr
	}

	if thisTyped < thatTyped {
		return -1
	} else if thisTyped > thatTyped {
		return 1
	}

	return 0
}

func stringComparer(this, that interface{}) (int, error) {
	thisTyped, thisTypedErr := castAsString(this)
	if thisTypedErr != nil {
		return 0, thisTypedErr
	}

	thatTyped, thatTypedErr := castAsString(that)
	if thatTypedErr != nil {
		return 0, thatTypedErr
	}

	return strings.Compare(thisTyped, thatTyped)
}

func castAsInt16(value interface{}) (int16, error) {
	if typedValue, isTyped := value.(int16); isTyped {
		return typedValue, nil
	} else {
		destinationType := reflect.TypeOf(int16(0))
		valueType := reflect.TypeOf(value)
		valueReflected := reflect.ValueOf(value)
		if valueType.ConvertibleTo(destinationType) {
			return valueReflected.Convert(intType).Interface().(int16), nil
		}
	}

	valueType := reflect.TypeOf(value)
	return int16(0), exception.New("Cannot cast %v as int16", valueType.Name())
}

func castAsInt(value interface{}) (int, error) {
	if typedValue, isTyped := value.(int); isTyped {
		return typedValue, nil
	} else {
		destinationType := reflect.TypeOf(int(0))
		valueType := reflect.TypeOf(value)
		valueReflected := reflect.ValueOf(value)
		if valueType.ConvertibleTo(destinationType) {
			return valueReflected.Convert(intType).Interface().(int), nil
		}
	}

	valueType := reflect.TypeOf(value)
	return int(0), exception.New("Cannot cast %v as int", valueType.Name())
}

func castAsInt64(value interface{}) (int64, error) {
	if typedValue, isTyped := value.(int64); isTyped {
		return typedValue, nil
	} else {
		destinationType := reflect.TypeOf(int64(0))
		valueType := reflect.TypeOf(value)
		valueReflected := reflect.ValueOf(value)
		if valueType.ConvertibleTo(destinationType) {
			return valueReflected.Convert(intType).Interface().(int64), nil
		}
	}

	valueType := reflect.TypeOf(value)
	return int64(0), exception.New("Cannot cast %v as int64", valueType.Name())
}

func castAsFloat32(value interface{}) (float32, error) {
	if typedValue, isTyped := value.(float32); isTyped {
		return typedValue, nil
	} else {
		destinationType := reflect.TypeOf(float32(0))
		valueType := reflect.TypeOf(value)
		valueReflected := reflect.ValueOf(value)
		if valueType.ConvertibleTo(destinationType) {
			return valueReflected.Convert(intType).Interface().(float32), nil
		}
	}

	valueType := reflect.TypeOf(value)
	return float32(0), exception.New("Cannot cast %v as float32", valueType.Name())
}

func castAsFloat64(value interface{}) (float64, error) {
	if typedValue, isTyped := value.(float64); isTyped {
		return typedValue, nil
	} else {
		destinationType := reflect.TypeOf(float64(0))
		valueType := reflect.TypeOf(value)
		valueReflected := reflect.ValueOf(value)
		if valueType.ConvertibleTo(destinationType) {
			return valueReflected.Convert(intType).Interface().(float64), nil
		}
	}

	valueType := reflect.TypeOf(value)
	return float64(0), exception.New("Cannot cast %v as float64", valueType.Name())
}

func castAsString(value interface{}) (string, error) {
	if valueAsString, isString := value.(string); isString {
		return valueAsString, nil
	} else {
		valueType := reflect.TypeOf(value)
		return "", exception.New("Cannot cast %v as string", valueType.Name())
	}
}
