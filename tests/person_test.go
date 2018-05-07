package tests

import (
	"testing"
)

const schema = `{
  "type": "object",
  "properties": {
    "name": {
      "type": "string"
    }
  },
  "required": ["origin"]
}`

func Test_GetPerson(t *testing.T) {
	BalooClient.Get("/people/1").
		Expect(t).
		Status(200).
		Type("json").
		JSONSchema(schema).
		Done()
}
