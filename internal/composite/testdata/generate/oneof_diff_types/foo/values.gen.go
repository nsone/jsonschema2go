// Code generated by jsonschema2go. DO NOT EDIT.
package foo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"regexp"
)

// Bar gives you some dumb info
// generated from https://example.com/testdata/generate/oneof_diff_types/foo/bar.json
type Bar struct {
	Value interface{}
}

var (
	barValuePattern = regexp.MustCompile(`^[0-9]{22}$`)
)

func (m *Bar) Validate() error {
	if v, ok := m.Value.(float64); ok {
		if v < 3.7 {
			return &validationError{
				errType: "minimum",
				message: fmt.Sprintf("must be greater than or equal to 3.7 but was %v", v),
			}
		}
	}
	if v, ok := m.Value.(string); ok {
		if !barValuePattern.MatchString(v) {
			return &validationError{
				errType: "pattern",
				message: fmt.Sprintf(`must match '^[0-9]{22}$' but got %q`, v),
			}
		}
	}
	if v, ok := m.Value.(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return err
		}
	}
	return nil
}

func (m *Bar) UnmarshalJSON(data []byte) error {
	tok, err := json.NewDecoder(bytes.NewReader(data)).Token()
	if err != nil {
		return err
	}
	switch t := tok.(type) {
	case json.Delim:
		if t == '{' {
			var obj Baz
			if err := json.Unmarshal(data, &obj); err != nil {
				return err
			}
			m.Value = obj
			return nil
		}
		if t == '[' {
			var arr Bazes
			if err := json.Unmarshal(data, &arr); err != nil {
				return err
			}
			m.Value = arr
			return nil
		}
	case string:
		var s string
		if err := json.Unmarshal(data, &s); err != nil {
			return err
		}
		m.Value = s
		return nil
	case float64:
		var f float64
		if err := json.Unmarshal(data, &f); err != nil {
			return err
		}
		m.Value = f
		return nil
	}
	if tok == nil {
		return nil
	}
	return fmt.Errorf("unsupported type: %T", tok)
}

func (m *Bar) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.Value)
}

// generated from https://example.com/testdata/generate/oneof_diff_types/foo/bar.json#/oneOf/0
type Baz struct {
	Baz *string `json:"baz,omitempty"`
}

func (m *Baz) Validate() error {
	return nil
}

// generated from https://example.com/testdata/generate/oneof_diff_types/foo/bar.json#/oneOf/1
type Bazes []string

func (m Bazes) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte(`[]`), nil
	}
	return json.Marshal([]string(m))
}

func (m Bazes) Validate() error {
	return nil
}

type valErr interface {
	ErrType() string
	JSONPath() []interface{}
	Path() []interface{}
	Message() string
}

type validationError struct {
	errType, message string
	jsonPath, path   []interface{}
}

func (e *validationError) ErrType() string {
	return e.errType
}

func (e *validationError) JSONPath() []interface{} {
	return e.jsonPath
}

func (e *validationError) Path() []interface{} {
	return e.path
}

func (e *validationError) Message() string {
	return e.message
}

func (e *validationError) Error() string {
	return fmt.Sprintf("%v: %v", e.path, e.message)
}

var _ valErr = new(validationError)
