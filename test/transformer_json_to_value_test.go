package test

import (
	"github.com/lanvard/support/transformer"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJsonToValue(t *testing.T) {
	value := transformer.JsonToValue(`{"name":{"first":"Janet","last":"Prichard"},"age":47}`)

	assert.Equal(t, "Janet", value.Get("name.first").String())
}

func TestDeepJsonToValue(t *testing.T) {
	value := transformer.JsonToValue(`{
  "data": {
    "tracktraces": [
      {
        "history": [
          {
            "code": "A01",
            "status": 2,
            "main": "registered"
          },
          {
            "code": "B01",
            "status": 3,
            "main": "handed_to_carrier"
          },
          {
            "code": "J01",
            "status": 4,
            "main": "sorting"
          },
          {
            "code": "J05",
            "status": 5,
            "main": "distribution"
          }
        ]
      }
    ]
  }
}`)

	assert.Equal(t, "J01", value.Get("data.tracktraces.0.history.2.code").String())
}
