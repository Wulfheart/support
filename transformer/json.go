package transformer

import (
	"github.com/lanvard/support"
	"github.com/tidwall/gjson"
)

func JsonToValue(json string) support.Value {
	result := gjson.Parse(json).Value()

	return support.NewValue(result)
}

func ToJson(raw interface{}) (string, error) {
	// switch value := raw.(type) {
	// case support.Value:
	// 	return value.StringE()

	// }
	return "", nil
}
