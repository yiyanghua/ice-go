package codec

import (
	"reflect"
	"testing"
)

func TestPkg(t *testing.T) {
	h := &Header{}
	path := reflect.TypeOf(h).PkgPath()
	println(path)
}
