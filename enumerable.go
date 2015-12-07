package collections

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
