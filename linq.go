package collections

// KeySelector returns a field from a given value
type KeySelector func(value interface{}) interface{}

// MapAction is the projection from the value to a different value
type MapAction func(value interface{}) interface{}

// Predicate is a function that takes a value and returns a true or false
type Predicate func(value interface{}) bool

func Map(collection Enumerable, mapFn MapAction) Enumerable {
	newList := &List{}
	e := collection.GetEnumerator()
	hasNext := true
	for hasNext {
		newList.Add(mapFn(e.GetCurrent()))
		hasNext = e.MoveNext()
	}
	return newList
}

func Filter(collection Enumerable, predicate Predicate) Enumerable {
	newList := &List{}
	e := collection.GetEnumerator()
	hasNext := true
	for hasNext {
		current := e.GetCurrent()
		if predicate(current) {
			newList.Add(current)
		}
		hasNext = e.MoveNext()
	}
	return newList
}

func ToList(collection Enumerable) Enumerable {
	if typedCollection, isList := collection.(List); isList {
		return typedCollection
	}

	newList := &List{}
	e := collection.GetEnumerator()
	hasNext := true
	for hasNext {
		newList.Add(e.GetCurrent())
	}
	return newList
}

func Sort(collection Enumerable, sortPredicate SortSelector) Enumerable {
	asList := ToList(collection).(*List)

}
