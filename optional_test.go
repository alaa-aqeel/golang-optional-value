package optional

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestOptionalSetValue(t *testing.T) {
	o := SetValue("hello")
	if !reflect.DeepEqual(o, Optional[string]{Value: "hello", IsSet: true}) {
		t.Fatal("Failed to create optional value")
	}

	o = SetValueWithDefault("hello", "world")
	if !reflect.DeepEqual(o, Optional[string]{Value: "hello", IsSet: true, Default: "world"}) {
		t.Fatal("Failed to create optional value with default value")
	}

	o = SetValueWithDefault("hello", "")
	if !reflect.DeepEqual(o, Optional[string]{Value: "hello", IsSet: true, Default: ""}) {
		t.Fatal("Failed to create optional value with default value")
	}
}

func TestOptionalWithJson(t *testing.T) {
	jsonInput := []byte(`{"name": "test", "active": true}`)
	var args struct {
		Name   Optional[string] `json:"name"`
		Active Optional[bool]   `json:"active"`
	}
	json.Unmarshal(jsonInput, &args)

	if !reflect.DeepEqual(args.Name, SetValue("test")) {
		t.Fatal("Failed to unmarshal optional value")
	}
	if !reflect.DeepEqual(args.Active, SetValue(true)) {
		t.Fatal("Failed to unmarshal optional value")
	}
}

func TestOptionalGetDefaultValue(t *testing.T) {
	o := SetValueWithDefault[int64](0, 10)

	if !reflect.DeepEqual(o, Optional[int64]{Value: 0, IsSet: true, Default: 10}) {
		t.Fatal("Failed to create optional value with default value")
	}

	jsonInput := []byte(`{"name": "hello world"}`)
	var args struct {
		Name   Optional[string] `json:"name"`
		Active Optional[bool]   `json:"active"`
		Age    Optional[int64]  `json:"age"`
	}
	json.Unmarshal(jsonInput, &args)

	if v := args.Name.ValueOrDefault("hi world"); v != "hello world" {
		t.Fatal("Failed to unmarshal optional value : " + v)
	}

	if args.Active.ValueOrDefault(true) != true {
		t.Fatal("Failed to unmarshal optional value")
	}

	if args.Age.ValueOrDefault(20) != 20 {
		t.Fatal("Failed to unmarshal optional value")
	}
}

func TestOptional(t *testing.T) {
	if !SetValue(" ").IsEmpty() {
		t.Fatal("Failed to check if optional value is empty")
	}

	if !SetValue(0).IsZero() {
		t.Fatal("Failed to check if optional value is zero")
	}
	if !SetValue(0.0).IsZero() {
		t.Fatal("Failed to check if optional value is zero")
	}
}
