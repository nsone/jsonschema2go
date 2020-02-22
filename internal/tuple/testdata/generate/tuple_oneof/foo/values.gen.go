// Code generated by jsonschema2go. DO NOT EDIT.
package foo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"regexp"
)

// generated from https://example.com/testdata/generate/tuple_oneof/foo/bar.json#/items/2
type Baz struct {
	Value interface{}
}

func (m *Baz) Validate() error {
	return nil
}

func (m *Baz) UnmarshalJSON(data []byte) error {
	tok, err := json.NewDecoder(bytes.NewReader(data)).Token()
	if err != nil {
		return err
	}
	switch tok.(type) {
	case float64:
		var i int64
		if err := json.Unmarshal(data, &i); err != nil {
			return err
		}
		m.Value = i
		return nil
	case bool:
		var b bool
		if err := json.Unmarshal(data, &b); err != nil {
			return err
		}
		m.Value = b
		return nil
	}
	return fmt.Errorf("unsupported type: %T", tok)
}

func (m *Baz) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.Value)
}

// Bar gives you some dumb info
// generated from https://example.com/testdata/generate/tuple_oneof/foo/bar.json
type Bar [3]interface{}

var (
	bar0Pattern = regexp.MustCompile(`^abcdef$`)
)

func (t *Bar) Validate() error {
	if v, ok := m[0].(string); !ok {
		return &validationError{
			errType:  "type",
			path:     []interface{}{0},
			jsonPath: []interface{}{0},
			message:  fmt.Sprintf("must be string but got %T", m[0]),
		}
	} else if !bar0Pattern.MatchString(v) {
		return &validationError{
			errType:  "pattern",
			path:     []interface{}{0},
			jsonPath: []interface{}{0},
			message:  fmt.Sprintf(`must match '^abcdef$' but got %q`, v),
		}
	}
	if v, ok := m[1].(float64); !ok {
		return &validationError{
			errType:  "type",
			path:     []interface{}{1},
			jsonPath: []interface{}{1},
			message:  fmt.Sprintf("must be float64 but got %T", m[1]),
		}
	} else if v < 42.3 {
		return &validationError{
			errType:  "minimum",
			path:     []interface{}{1},
			jsonPath: []interface{}{1},
			message:  fmt.Sprintf("must be greater than or equal to 42.3 but was %v", v),
		}
	}
	return nil
}

func (t *Bar) UnmarshalJSON(data []byte) error {
	var msgs []json.RawMessage
	if err := json.Unmarshal(data, &msgs); err != nil {
		return err
	}
	if len(msgs) > 0 {
		var item string
		if err := json.Unmarshal(msgs[0], &item); err != nil {
			return err
		}
		t[0] = item
	}
	if len(msgs) > 1 {
		var item float64
		if err := json.Unmarshal(msgs[1], &item); err != nil {
			return err
		}
		t[1] = item
	}
	if len(msgs) > 2 {
		var item Baz
		if err := json.Unmarshal(msgs[2], &item); err != nil {
			return err
		}
		t[2] = item
	}
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
