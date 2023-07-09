package thiggle

import "encoding/json"

// Bool represents a nullable boolean type.
type Bool struct {
	Value  bool
	Exists bool
}

// NewBool creates a new Bool value. The resulting value will return true
// for Exists.
func NewBool(value bool) Bool {
	return Bool{Value: value, Exists: true}
}

// MarshalJSON marshals the Bool value or null if the bool doesn't exist.
func (b Bool) MarshalJSON() ([]byte, error) {
	if !b.Exists {
		return []byte("null"), nil
	}
	return json.Marshal(b.Value)
}

// UnmarshalJSON unmarshals an Bool value or null into the provided boolean.
func (b *Bool) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		b.Value, b.Exists = false, false
		return nil
	}

	var val bool
	if err := json.Unmarshal(data, &val); err != nil {
		return err
	}

	b.Value, b.Exists = val, true
	return nil
}
