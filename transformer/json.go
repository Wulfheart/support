package transformer

import (
	"github.com/lanvard/support"
	"github.com/tidwall/gjson"
)

func JsonToValue(json string) support.Value {
	result := gjson.Parse(json).Value()

	return support.NewValue(result)
}

func ValueToJson(value support.Value) (string, error) {
	return value.String(), nil
}