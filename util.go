package collections

import "reflect"

func reflectValue(obj interface{}) reflect.Value {
	v := reflect.ValueOf(obj)
	for v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface {
		v = v.Elem()
	}
	return v
}

func reflectType(obj interface{}) reflect.Type {
	t := reflect.TypeOf(obj)
	for t.Kind() == reflect.Ptr || t.Kind() == reflect.Interface {
		t = t.Elem()
	}

	return t
}

func getSliceType(collection interface{}) reflect.Type {
	cv := reflect.ValueOf(collection)
	t := reflect.TypeOf(collection)
	if cv.Len() > 0 {
		firstElem := cv.Index(0)
		t = reflect.TypeOf(firstElem.Interface())
		for t.Kind() == reflect.Ptr || t.Kind() == reflect.Interface {
			t = t.Elem()
		}
	} else {
		for t.Kind() == reflect.Ptr || t.Kind() == reflect.Interface {
			t = t.Elem()
		}
	}

	return t
}

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
