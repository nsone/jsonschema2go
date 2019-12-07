// Code generated by jsonschema2go. DO NOT EDIT.
package foo

import (
	"encoding/json"
	"fmt"
	"github.com/jwilner/jsonschema2go/boxed"
)

// Bar gives you some dumb info
type Bar struct {
	Inner Excluded     `json:"inner,omitempty"`
	Name  boxed.String `json:"name"`
}

func (m *Bar) Validate() error {
	if err := m.Inner.Validate(); err != nil {
		if err, ok := err.(valErr); ok {
			return &validationError{
				errType:  err.ErrType(),
				message:  err.Message(),
				path:     append([]interface{}{"Inner"}, err.Path()...),
				jsonPath: append([]interface{}{"inner"}, err.JSONPath()...),
			}
		}
		return err
	}
	return nil
}

func (m *Bar) MarshalJSON() ([]byte, error) {
	inner := struct {
		Inner Excluded `json:"inner,omitempty"`
		Name  *string  `json:"name,omitempty"`
	}{
		Inner: m.Inner,
	}
	if m.Name.Set {
		inner.Name = &m.Name.String
	}
	return json.Marshal(inner)
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