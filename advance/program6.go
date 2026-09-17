package main

import (
	"fmt"
	"reflect"
)

type Config struct {
	Host string `env:"APP_HOST" default:"localhost"`
	Port int    `env:"APP_PORT" default:"8080"`
}

func InspectStructTags(target interface{}) {
	t := reflect.TypeOf(target)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		envTag := field.Tag.Get("env")
		defaultTag := field.Tag.Get("default")
		fmt.Printf("Field: %-6s | Type: %-6s | env: %-10s | default: %s\n", 
			field.Name, field.Type, envTag, defaultTag)
	}
}

func main() {
	cfg := Config{}
	InspectStructTags(cfg)
}