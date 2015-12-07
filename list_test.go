package collections

import (
	"testing"

	"github.com/blendlabs/go-assert"
)

func TestListAdd(t *testing.T) {
	a := assert.New(t)

	fromSlice := NewList([]int{1, 2, 3, 4})
	a.NotNil(fromSlice)
	a.Equal(4, fromSlice.Len())

	fromSlice.Add(5)
	a.Equal(5, fromSlice.Len())

	value := fromSlice.At(4)
	a.Equal(5, value)
}

func TestListRemoveAt(t *testing.T) {
	a := assert.New(t)

	fromSlice := NewList([]int{1, 2, 3, 4})

	removeErr := fromSlice.RemoveAt(3)
	a.Nil(removeErr)
	a.Equal(3, fromSlice.Len())
}

func TestListGetEnumerator(t *testing.T) {
	a := assert.New(t)

	fromSlice := NewList([]int{1, 2, 3, 4})
	se := fromSlice.GetEnumerator()
	a.NotNil(se)
	value := se.GetCurrent()
	a.Equal(1, value)
}
