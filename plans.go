package jsonschema2go

import (
	"fmt"
	"strconv"
)

type TypeInfo struct {
	GoPath  string
	Name    string
	Pointer bool
}

func (t TypeInfo) BuiltIn() bool {
	return t.GoPath == ""
}

func (t TypeInfo) Unknown() bool {
	return t == TypeInfo{}
}

var primitives = map[SimpleType]string{
	Boolean: "bool",
	Integer: "int",
	Number:  "float",
	Null:    "interface{}",
	String:  "string",
}

type StructField struct {
	Comment string
	Names   []string
	Type    TypeInfo
	Tag     string
}

type Plan interface {
	Type() TypeInfo
	Deps() []TypeInfo
}

type StructPlan struct {
	typeInfo TypeInfo

	Comment string
	Fields  []StructField
}

func (s *StructPlan) Type() TypeInfo {
	return s.typeInfo
}

func (s *StructPlan) Deps() (deps []TypeInfo) {
	for _, s := range s.Fields {
		deps = append(deps, s.Type)
	}
	return
}

type ArrayPlan struct {
	typeInfo TypeInfo

	Comment  string
	ItemType TypeInfo
}

func (a *ArrayPlan) Type() TypeInfo {
	return a.typeInfo
}

func (a *ArrayPlan) Deps() []TypeInfo {
	return []TypeInfo{a.ItemType, {Name: "Marshal", GoPath: "encoding/json"}}
}

type EnumPlan struct {
	typeInfo TypeInfo

	Comment  string
	BaseType TypeInfo
	Members  []EnumMember
}

type EnumMember struct {
	Name  string
	Field interface{}
}

func (e *EnumPlan) Literal(val interface{}) string {
	switch t := val.(type) {
	case bool:
		return strconv.FormatBool(t)
	case string:
		return fmt.Sprintf("%q", t)
	default:
		return fmt.Sprintf("%d", t)
	}
}

func (e *EnumPlan) Type() TypeInfo {
	return e.typeInfo
}

func (e *EnumPlan) Deps() []TypeInfo {
	return []TypeInfo{e.BaseType}
}
