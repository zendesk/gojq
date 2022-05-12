package gojq

import (
	"testing"
)

func TestTracing(t *testing.T) {
	query, err := Parse(".foo * 100000")
	if err != nil {
		t.Fatal(err)
	}

	iter := query.Run(map[string]interface{}{
		"foo": "bar",
	})

	var item interface{}
	var ok bool
	for {
		item, ok = iter.Next()
		if !ok {
			break
		}

		t.Logf("item: %#v", item)
	}
}
