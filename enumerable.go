package collections

import "reflect"

// --------------------------------------------------------------------------------
// exported interfaces
// --------------------------------------------------------------------------------

type Enumerable interface {
	GetEnumerator() Enumerator
}

type Enumerator interface {
	MoveNext() bool
	GetCurrent() interface{}
	Reset()
}

// --------------------------------------------------------------------------------
// mapEnumerator
// --------------------------------------------------------------------------------

type mapEnumerator struct {
	length   int
	index    int
	keys     []interface{}
	contents interface{}
}

func NewMapEnumerator(contents interface{}) *mapEnumerator {
	if contents == nil {
		return nil
	}

	if !isMap(contents) {
		return nil
	}

	se := mapEnumerator{}
	se.length = getMapLength(contents)
	se.index = 0
	se.keys = getMapKeys(contents)
	se.contents = contents
	return &se
}

func (se *mapEnumerator) MoveNext() bool {
	if se.index >= se.length {
		return false
	}

	se.index = se.index + 1
	return se.index < se.length
}

func (se mapEnumerator) GetCurrent() interface{} {
	if se.index < se.length {
		key := se.keys[se.index]
		return elementAtKey(se.contents, key)
	}
	return nil
}

func (se mapEnumerator) GetCurrentWithKey() (interface{}, interface{}) {
	if se.index < se.length {
		key := se.keys[se.index]
		return key, elementAtKey(se.contents, key)
	}
	return nil, nil
}

func (se *mapEnumerator) Reset() {
	se.index = 0
	se.length = getMapLength(se.contents)
}

// --------------------------------------------------------------------------------
// sliceEnumerator
// --------------------------------------------------------------------------------

type sliceEnumerator struct {
	length   int
	index    int
	contents interface{}
}

func NewSliceEnumerator(contents interface{}) *sliceEnumerator {
	if contents == nil {
		return nil
	}

	if !isSlice(contents) {
		return nil
	}

	se := sliceEnumerator{}
	se.length = getSliceLength(contents)
	se.index = 0
	se.contents = contents
	return &se
}

func (se *sliceEnumerator) MoveNext() bool {
	if se.index >= se.length {
		return false
	}
	se.index = se.index + 1
	return se.index < se.length
}

func (se sliceEnumerator) GetCurrent() interface{} {
	if se.index < se.length {
		return elementAtIndex(se.contents, se.index)
	}
	return nil
}

func (se *sliceEnumerator) Reset() {
	se.index = 0
	se.length = getSliceLength(se.contents)
}

// --------------------------------------------------------------------------------
// internal utility functions
// --------------------------------------------------------------------------------

func getSliceLength(slice interface{}) int {
	if slice == nil {
		return 0
	}

	sliceValue := reflect.ValueOf(slice)
	return sliceValue.Len()
}

func getMapLength(contents interface{}) int {
	if contents == nil {
		return 0
	}

	contentsValue := reflect.ValueOf(contents)
	return contentsValue.Len()
}

func isSlice(thing interface{}) bool {
	if thing == nil {
		return false
	}

	thingType := reflect.TypeOf(thing)
	return thingType.Kind() == reflect.Slice
}

func isMap(thing interface{}) bool {
	if thing == nil {
		return false
	}

	thingType := reflect.TypeOf(thing)
	return thingType.Kind() == reflect.Map
}

func elementAtIndex(slice interface{}, index int) interface{} {
	if slice == nil {
		return nil
	}

	ofSliceValue := reflect.ValueOf(slice)
	return ofSliceValue.Index(index).Interface()
}

func elementAtKey(contents interface{}, key interface{}) interface{} {
	if contents == nil {
		return nil
	}

	contentsValue := reflect.ValueOf(contents)
	return contentsValue.MapIndex(reflect.ValueOf(key)).Interface()
}

func getMapKeys(contents interface{}) []interface{} {
	if contents == nil {
		return nil
	}

	contentsValue := reflect.ValueOf(contents)
	mapKeys := contentsValue.MapKeys()

	keys := make([]interface{}, len(mapKeys))
	for i, keyValue := range mapKeys {
		keys[i] = keyValue.Interface()
	}
	return keys
}
