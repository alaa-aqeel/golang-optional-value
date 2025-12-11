package optional

import (
	"encoding/json"
	"fmt"
)

func (o Optional[T]) MarshalJSON() ([]byte, error) {
	if !o.IsSet {
		// If the value is not set, return 'null' to omit the field's data structure
		return []byte("null"), nil
	}
	// If set, marshal only the underlying Value
	return json.Marshal(o.Value)
}

func (o *Optional[T]) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		// If the incoming JSON value is 'null', the field is not set.
		o.IsSet = false
		return nil
	}

	// Otherwise, the value is set. Try to unmarshal the data into the Value field.
	if err := json.Unmarshal(data, &o.Value); err != nil {
		return fmt.Errorf("failed to unmarshal optional value: %w", err)
	}

	// Successfully unmarshaled a non-null value
	o.IsSet = true
	return nil
}
