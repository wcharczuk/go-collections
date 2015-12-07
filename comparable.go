package collections

import (
	"reflect"
	"strings"

	"github.com/blendlabs/go-exception"
)

// --------------------------------------------------------------------------------
// exported interfaces
// --------------------------------------------------------------------------------

// An interface that provides the mechanism for a type to be sorted
type Comparable interface {
	// CompareTo should return -1 if it is less than other, 0 if equal to other, and 1 if greater than other
	CompareTo(other interface{}) (int, error)
}

// Comparer should return -1 if it is less than other, 0 if equal to other, and 1 if greater than other
type Comparer func(this, that interface{}) (int, error)

// --------------------------------------------------------------------------------
// we need this for sorting
// --------------------------------------------------------------------------------

func getComparer(forThis interface{}) (Comparer, error) {
	if !isSlice(forThis) {
		return nil, exception.Newf("%v is not a slice", forThis)
	}

	forThisType := getSliceType(forThis)

	switch forThisType.Kind() {
	case reflect.Int16:
		return int16Comparer, nil
	case reflect.Int32, reflect.Int:
		return intComparer, nil
	case reflect.Int64:
		return int64Comparer, nil
	case reflect.Float32:
		return float32Comparer, nil
	case reflect.Float64:
		return float64Comparer, nil
	default:
		if typed, isComparable := forThis.(Comparable); isComparable {
			return wrapComparable(typed), nil
		} else {
			return nil, exception.Newf("%v does not implement Comparable and is not a builtin type.", forThisType.Name())
		}
	}

	return nil, nil
}

// --------------------------------------------------------------------------------
// Comparable helpers
// --------------------------------------------------------------------------------

func wrapComparable(source Comparable) Comparer {
	return func(this, that interface{}) (int, error) {
		return source.CompareTo(that)
	}
}

// --------------------------------------------------------------------------------
// comparers for builtins
// --------------------------------------------------------------------------------

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
		return -1, nil
	} else if thisTyped > thatTyped {
		return 1, nil
	}

	return 0, nil
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
		return -1, nil
	} else if thisTyped > thatTyped {
		return 1, nil
	}

	return 0, nil
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
		return -1, nil
	} else if thisTyped > thatTyped {
		return 1, nil
	}

	return 0, nil
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
		return -1, nil
	} else if thisTyped > thatTyped {
		return 1, nil
	}

	return 0, nil
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
		return -1, nil
	} else if thisTyped > thatTyped {
		return 1, nil
	}

	return 0, nil
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

	return strings.Compare(thisTyped, thatTyped), nil
}

func castAsInt16(value interface{}) (int16, error) {
	if typedValue, isTyped := value.(int16); isTyped {
		return typedValue, nil
	} else {
		destinationType := reflect.TypeOf(int16(0))
		valueType := reflect.TypeOf(value)
		valueReflected := reflect.ValueOf(value)
		if valueType.ConvertibleTo(destinationType) {
			return valueReflected.Convert(destinationType).Interface().(int16), nil
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
			return valueReflected.Convert(destinationType).Interface().(int), nil
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
			return valueReflected.Convert(destinationType).Interface().(int64), nil
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
			return valueReflected.Convert(destinationType).Interface().(float32), nil
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
			return valueReflected.Convert(destinationType).Interface().(float64), nil
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
