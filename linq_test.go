package collections

import (
	"testing"

	"github.com/blendlabs/go-assert"
)

func TestLinqMap(t *testing.T) {
	a := assert.New(t)

	l := NewList(1, 2, 3, 4)
	mapped := Map(l, func(value interface{}) interface{} {
		return (value.(int)) * 2
	}).(*List)

	a.Equal(2, mapped.At(0))
	a.Equal(4, mapped.At(1))
	a.Equal(6, mapped.At(2))
	a.Equal(8, mapped.At(3))
}

func TestLinqFilter(t *testing.T) {
	a := assert.New(t)

	l := NewList(1, 2, 3, 4)
	filtered := Filter(l, func(value interface{}) bool {
		return value.(int) < 3
	}).(*List)

	a.Equal(2, filtered.Len())
	a.Equal(1, filtered.At(0))
	a.Equal(2, filtered.At(1))
}

func TestLinqSortInts(t *testing.T) {
	a := assert.New(t)
	ints := NewList(7, 5, 3, 8, 2, 1, 4, 6)

	sorted := SortBy(ints, DefaultKeySelector).(*List)

	a.Equal(1, sorted.At(0))
	a.Equal(2, sorted.At(1))
	a.Equal(3, sorted.At(2))
	a.Equal(4, sorted.At(3))
	a.Equal(5, sorted.At(4))
	a.Equal(6, sorted.At(5))
	a.Equal(7, sorted.At(6))
	a.Equal(8, sorted.At(7))
}

func TestLinqSortDescendingInts(t *testing.T) {
	a := assert.New(t)
	ints := NewList(7, 5, 3, 8, 2, 1, 4, 6)

	sorted := SortByDescending(ints, DefaultKeySelector).(*List)

	a.Equal(8, sorted.At(0))
	a.Equal(7, sorted.At(1))
	a.Equal(6, sorted.At(2))
	a.Equal(5, sorted.At(3))
	a.Equal(4, sorted.At(4))
	a.Equal(3, sorted.At(5))
	a.Equal(2, sorted.At(6))
	a.Equal(1, sorted.At(7))
}

type myTestType struct {
	Id   int
	Name string
}

func TestLinqSortStructs(t *testing.T) {
	a := assert.New(t)

	l := NewList(myTestType{Id: 1, Name: "Foo"}, myTestType{Id: 2, Name: "Bar"}, myTestType{Id: 3, Name: "Baz"})

	sorted := SortBy(l, func(v interface{}) interface{} {
		return (v.(myTestType)).Id
	}).(*List)

	a.Equal(3, sorted.At(0))
	a.Equal(2, sorted.At(1))
	a.Equal(1, sorted.At(2))
}
