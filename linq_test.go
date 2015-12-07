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

	a.Equal(2, filtered.GetLength())
	a.Equal(1, filtered.At(0))
	a.Equal(2, filtered.At(1))
}
