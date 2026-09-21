package main

import (
	"fmt"
	"reflect"
)

type AppConfig struct {
	DatabaseURL string
	MaxRetries  int
}

func SetFieldDynamically(target interface{}, fieldName string, value interface{}) error {
	v := reflect.ValueOf(target)

	// Ensure target is a pointer to a struct
	if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("target must be a pointer to a struct")
	}

	field := v.Elem().FieldByName(fieldName)
	if !field.IsValid() {
		return fmt.Errorf("no such field: %s", fieldName)
	}

	if !field.CanSet() {
		return fmt.Errorf("cannot set field %s", fieldName)
	}

	val := reflect.ValueOf(value)
	if field.Type() != val.Type() {
		return fmt.Errorf("provided value type didn't match struct field type")
	}

	field.Set(val)
	return nil
}

func main() {
	cfg := AppConfig{}
	
	SetFieldDynamically(&cfg, "DatabaseURL", "postgres://user:pass@localhost/db")
	SetFieldDynamically(&cfg, "MaxRetries", 5)

	fmt.Printf("Mutated Config: %+v\n", cfg)
}