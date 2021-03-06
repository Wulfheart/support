package support

import (
	"fmt"
	"reflect"
)

func Name(element interface{}) string {
	if Kind(element) == reflect.String {
		return element.(string)
	}

	if Kind(element) == reflect.Struct {
		return reflect.TypeOf(element).String()
	}

	if Kind(element) == reflect.Ptr && element == nil {
		panic("Nil value found. To bind an interface, use the following syntax: (*INTERFACE)(nil)")
	}

	return reflect.TypeOf(element).Elem().String()
}

func Package(element interface{}) string {

	if element == nil {
		return reflect.TypeOf(&element).Elem().PkgPath()
	}

	return reflect.TypeOf(element).Elem().PkgPath()
}

func Kind(element interface{}) reflect.Kind {
	if element == nil {
		return reflect.TypeOf(&element).Kind()
	}

	switch element.(type) {
	case Collection:
		return reflect.Slice
	case Map:
		return reflect.Map
	}

	return reflect.TypeOf(element).Kind()
}

func Dump(expression ...interface{}) {
	fmt.Println(fmt.Sprintf("%#v", expression))
}
