// +build go1.11,!go1.12

package stdlib

// Code generated by 'goexports image/jpeg'. DO NOT EDIT.

import (
	"image/jpeg"
	"reflect"
)

func init() {
	Value["image/jpeg"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"Decode":         reflect.ValueOf(jpeg.Decode),
		"DecodeConfig":   reflect.ValueOf(jpeg.DecodeConfig),
		"DefaultQuality": reflect.ValueOf(jpeg.DefaultQuality),
		"Encode":         reflect.ValueOf(jpeg.Encode),

		// type definitions
		"FormatError":      reflect.ValueOf((*jpeg.FormatError)(nil)),
		"Options":          reflect.ValueOf((*jpeg.Options)(nil)),
		"Reader":           reflect.ValueOf((*jpeg.Reader)(nil)),
		"UnsupportedError": reflect.ValueOf((*jpeg.UnsupportedError)(nil)),
	}
}