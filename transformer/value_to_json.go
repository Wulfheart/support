package transformer

import "github.com/lanvard/support"

type ValueToJson struct{}

func (v ValueToJson) IsValid(object interface{}) bool {
	value, ok := object.(support.Value)

	return ok && InterfaceToJson{}.IsValid(value.Raw())
}

func (v ValueToJson) Transform(object interface{}) string {
	value, ok := object.(support.Value)
	if !ok {
		panic("can not transform to json with an unsupported type")
	}

	return InterfaceToJson{}.Transform(value.Raw())
}
