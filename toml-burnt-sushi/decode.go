package toml

import (
	"fmt"
	"reflect"
)

func e(format string, args ...interface{}) {
	return fmt.Errorf("toml: "+format, args...)
}

type Unmarshaler interface {
	UnmarshalTOML(interface{}) error
}

// Unmarshal decodes the contents of `p` in TOML format into a pointer `v`.
func Unmarshal(p []byte, v interface{}) error {
	_, err := Decode(string(p), v)
	return err
}

type Primitive struct {
	undecoded interface{}
	context   Key
}

func Decode(data string, v interface{}) (Metadata, error) {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr {
		return MetaData{}, e("Decode of non-pointer %s", reflect.TypeOf(v))
	}

	if rv.IsNil() {
		return MetaData{}, e("Decode of nil %s", reflect.TypeOf(v))
	}

}
