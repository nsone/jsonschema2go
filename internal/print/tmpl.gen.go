// Code generated by internal/cmd/embedtmpl/embedtmpl.go DO NOT EDIT.
package print

import (
	"text/template"
)

var tmpl = template.Must(template.New("").Parse(`{{/* gotype: github.com/ns1/jsonschema2go.Plans */}}
// Code generated by jsonschema2go. DO NOT EDIT.
package {{ .Imports.CurPackage }}

{{ with .Imports.List -}}
import (
{{ range . -}}
    {{ if .Alias -}}{{ .Alias }} {{ end -}}"{{ .GoPath }}"
{{ end -}}
)
{{ end -}}

{{ range .Plans -}}
{{ .Execute $.Imports }}
{{ end -}}

type valErr interface {
    ErrType() string
    JSONPath() []interface{}
    Path() []interface{}
    Message() string
}

type validationError struct {
    errType, message string
    jsonPath, path []interface{}
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

var _ valErr = new(validationError)`))
