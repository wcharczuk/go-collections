package collections

import (
	"reflect"

	"github.com/blendlabs/go-exception"
)

func NewList(contentValue ...interface{}) *List {
	l := &List{}

	for _, value := range contentValue {
		if isSlice(value) {
			contentsValue := reflect.ValueOf(value)
			for i := 0; i < contentsValue.Len(); i++ {
				element := elementAtIndex(value, i)
				l.contents = append(l.contents, element)
			}
		} else {
			l.contents = append(l.contents, value)
		}
	}

	return l
}

type List struct {
	contents []interface{}
}

func (l *List) Add(item interface{}) {
	l.contents = append(l.contents, item)
}

func (l *List) At(index int) interface{} {
	return elementAtIndex(l.contents, index)
}

func (l *List) Last() interface{} {
	if len(l.contents) == 0 {
		return nil
	}

	return l.contents[l.Len()-1]
}

func (l *List) Len() int {
	if l.contents == nil {
		return 0
	}

	return len(l.contents)
}

func (l *List) RemoveAt(index int) error {
	if index >= l.Len() {
		return exception.Newf("Invalid index for RemoveAt(%d)", index)
	}

	if index == 0 {
		l.contents = l.contents[1:]
	} else if index == l.Len()-1 {
		l.contents = l.contents[0:index]
	} else {
		l.contents = append(l.contents[0:index], l.contents[index+1:]...)
	}
	return nil
}

func (l *List) Clear() {
	l.contents = []interface{}{}
}

func (l *List) GetEnumerator() Enumerator {
	return NewSliceEnumerator(l.contents)
}

func (l *List) Swap(i, j int) {
	l.contents[i], l.contents[j] = l.contents[j], l.contents[i]
}
