package als

import (
	"reflect"
)

// python's getattr in golang
// attr must be exported field
func GetAttr(object interface{}, attr string,
	default_ interface{}) (ret interface{}) {
	ret = default_
	objectValue := reflect.ValueOf(object)
	objectValue = reflect.Indirect(objectValue) // Dereference if it's a pointer.
	if objectValue.Kind().String() != "struct" {
		// `FieldByName` will panic if we're not a struct.
		return
	}

	attrVal := objectValue.FieldByName(attr)
	if !attrVal.IsValid() {
		return
	}
	return attrVal.Interface()
}
