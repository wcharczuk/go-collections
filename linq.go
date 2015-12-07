package collections

import "sort"

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

func Sort(collection Enumerable, sortKey KeySelector) Enumerable {
	collectionAsList := ToList(collection).(*List)
	comparer, comparerError := getComparer(collectionAsList.contents)
	if comparerError != nil {
		println(comparerError.Error())
	}

	sort.Sort(newSortableList(collectionAsList, comparer, false))

	return collectionAsList
}

func SortDescending(collection Enumerable, sortKey KeySelector) Enumerable {
	collectionAsList := ToList(collection).(*List)
	comparer, comparerError := getComparer(collectionAsList.contents)
	if comparerError != nil {
		println(comparerError.Error())
	}

	sort.Sort(newSortableList(collectionAsList, comparer, true))

	return collectionAsList
}

func ToList(collection Enumerable) Enumerable {
	if typedCollection, isList := collection.(*List); isList {
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

// --------------------------------------------------------------------------------
// internal Types
// --------------------------------------------------------------------------------

func newSortableList(contents *List, comparer Comparer, descending bool) *sortableList {
	return &sortableList{
		contents:   contents,
		comparer:   comparer,
		descending: descending,
	}
}

type sortableList struct {
	contents   *List
	comparer   Comparer
	descending bool
}

func (s *sortableList) Len() int {
	return s.contents.Len()
}

func (s *sortableList) Swap(i, j int) {
	s.contents.contents[i], s.contents.contents[j] = s.contents.contents[j], s.contents.contents[i]
}

func (s *sortableList) Less(i, j int) bool {
	iValue := s.contents.At(i)
	jValue := s.contents.At(j)

	compareResult, compareErr := s.comparer(iValue, jValue)
	if compareErr != nil {
		println(compareErr.Error())
	}

	if s.descending {
		return compareResult == 1
	} else {
		return compareResult == -1
	}
}
