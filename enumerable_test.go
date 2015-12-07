package collections

import (
	"testing"

	"github.com/blendlabs/go-assert"
)

func TestIsSlice(t *testing.T) {
	a := assert.New(t)

	aSlice := []int{1, 2, 3, 4}
	notASlice := 3.14

	a.True(isSlice(aSlice))
	a.False(isSlice(notASlice))
}

func TestIsMap(t *testing.T) {
	a := assert.New(t)

	myMap := map[string]string{"foo": "foo_value", "bar": "bar_value", "baz": "baz_value"}
	notMap := []int{1, 2, 3, 4}

	a.True(isMap(myMap))
	a.False(isMap(notMap))
}

func TestElementAtIndex(t *testing.T) {
	a := assert.New(t)

	mySlice := []int{1, 2, 3, 4}
	value := elementAtIndex(mySlice, 0)
	a.Equal(value, 1)

	value = elementAtIndex(mySlice, 1)
	a.Equal(value, 2)

	value = elementAtIndex(mySlice, 2)
	a.Equal(value, 3)

	value = elementAtIndex(mySlice, 3)
	a.Equal(value, 4)
}

func TestElementAtKey(t *testing.T) {
	a := assert.New(t)
	myMap := map[string]string{"foo": "foo_value", "bar": "bar_value", "baz": "baz_value"}

	value := elementAtKey(myMap, "foo")
	a.Equal(value, "foo_value")

	value = elementAtKey(myMap, "bar")
	a.Equal(value, "bar_value")

	value = elementAtKey(myMap, "baz")
	a.Equal(value, "baz_value")
}

func TestGetMapKeys(t *testing.T) {
	a := assert.New(t)

	myMap := map[string]string{"foo": "foo_value", "bar": "bar_value", "baz": "baz_value"}
	keys := getMapKeys(myMap)

	a.Len(keys, 3)
}

func TestSliceEnumeraotr(t *testing.T) {
	a := assert.New(t)

	mySlice := []int{1, 2, 3, 4}
	se := NewSliceEnumerator(mySlice)
	a.NotNil(se)
	a.Equal(se.length, 4)

	firstValue := se.GetCurrent()
	a.Equal(firstValue, 1)

	a.True(se.MoveNext())
	secondValue := se.GetCurrent()
	a.Equal(secondValue, 2)

	a.True(se.MoveNext())
	thirdValue := se.GetCurrent()
	a.Equal(thirdValue, 3)

	a.True(se.MoveNext())
	fourthValue := se.GetCurrent()
	a.Equal(fourthValue, 4)

	a.False(se.MoveNext())
	se.Reset()

	newFirstValue := se.GetCurrent()
	a.Equal(newFirstValue, 1)
}

func TestMapEnumeraotr(t *testing.T) {
	a := assert.New(t)

	myMap := map[string]string{"foo": "foo_value", "bar": "bar_value", "baz": "baz_value"}
	me := NewMapEnumerator(myMap)
	a.NotNil(me)
	a.Equal(me.length, 3)

	firstKey := me.keys[0]

	firstValue := me.GetCurrent()
	a.Equal(myMap[firstKey.(string)], firstValue)

	key, value := me.GetCurrentWithKey()
	a.Equal(firstKey, key)
	a.Equal(myMap[firstKey.(string)], value)

	a.True(me.MoveNext())

	secondKey := me.keys[1]
	key, value = me.GetCurrentWithKey()
	a.Equal(secondKey, key)
	a.Equal(myMap[secondKey.(string)], value)

	a.True(me.MoveNext())

	thirdKey := me.keys[2]
	key, value = me.GetCurrentWithKey()
	a.Equal(thirdKey, key)
	a.Equal(myMap[thirdKey.(string)], value)

	a.False(me.MoveNext())

	me.Reset()

	key, value = me.GetCurrentWithKey()
	a.Equal(firstKey, key)
	a.Equal(myMap[firstKey.(string)], value)
}
