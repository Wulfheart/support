package support

import (
	"fmt"
	"reflect"
)

func Name(element interface{}) string {
	if Type(element) == reflect.String {
		return element.(string)
	}

	if Type(element) == reflect.Struct {
		return reflect.TypeOf(element).String()
	}

	if Type(element) == reflect.Ptr && element == nil {
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

func Type(element interface{}) reflect.Kind {
	if element == nil {
		return reflect.TypeOf(&element).Kind()
	}

	switch element.(type) {
	case Collection:
		return reflect.Slice
	case Map:
		return reflect.Map
	case string:
		return reflect.String
	}

	return reflect.TypeOf(element).Kind()
}

func Dump(expression ...interface{}) {
	fmt.Println(fmt.Sprintf("%#v", expression))
}
