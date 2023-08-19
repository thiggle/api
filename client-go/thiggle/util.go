package thiggle

import (
	"encoding/json"
	"fmt"
)

type StringOrSlice []string

func (sos *StringOrSlice) MarshalJSON() ([]byte, error) {
	if len(*sos) == 1 {
		return json.Marshal((*sos)[0])
	}
	return json.Marshal([]string(*sos))
}

func MustNewStringOrSlice(v interface{}) *StringOrSlice {
	sos, err := NewStringOrSlice(v)
	if err != nil {
		panic(err)
	}
	return sos
}

func NewStringOrSlice(v interface{}) (*StringOrSlice, error) {
	switch v := v.(type) {
	case string:
		sos := StringOrSlice([]string{v})
		return &sos, nil
	case []string:
		sos := StringOrSlice(v)
		return &sos, nil
	case *StringOrSlice:
		return v, nil
	default:
		return nil, fmt.Errorf("invalid type for StringOrSlice: %T", v)
	}
}

func (sos *StringOrSlice) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err == nil {
		*sos = []string{s}
		return nil
	}
	var slice []string
	if err := json.Unmarshal(b, &slice); err != nil {
		return err
	}
	*sos = slice
	return nil
}

func (sos *StringOrSlice) Slice() []string {
	return *sos
}

func (sos *StringOrSlice) String() string {
	if len(*sos) == 0 {
		return ""
	}
	return (*sos)[0]
}

func (sos *StringOrSlice) Len() int {
	return len(*sos)
}

func (sos *StringOrSlice) Index(i int) string {
	if len(*sos) <= i {
		return ""
	}
	return (*sos)[i]
}

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
