package structtag

import (
	"errors"
	"reflect"
	"strconv"
	"testing"
)

type testStruct struct {
	A int    `cfg:"A" cfgDefault:"100"`
	B string `cfg:"B" cfgDefault:"200"`
	C string
	N string `cfg:"-"`
	p string
	S testSub `cfg:"S"`
}

type testSub struct {
	A int        `cfg:"A" cfgDefault:"300"`
	B string     `cfg:"C" cfgDefault:"400"`
	S testSubSub `cfg:"S"`
}
type testSubSub struct {
	A int    `cfg:"A" cfgDefault:"500"`
	B string `cfg:"S" cfgDefault:"600"`
}

func reflectIntTestFunc(field *reflect.StructField, value *reflect.Value, tag string) (err error) {
	newValue := field.Tag.Get(TagDefault)

	if newValue == "" {
		return
	}

	var intNewValue int64
	intNewValue, err = strconv.ParseInt(newValue, 10, 64)
	if err != nil {
		return
	}

	value.SetInt(intNewValue)

	return
}

func reflectStringTestFunc(field *reflect.StructField, value *reflect.Value, tag string) (err error) {
	newValue := field.Tag.Get(TagDefault)
	value.SetString(newValue)

	return
}

func reflectReturnError(field *reflect.StructField, value *reflect.Value, tag string) (err error) {
	err = errors.New("error test")
	return
}

func TestParse(t *testing.T) {

	Setup()

	s := &testStruct{A: 1, S: testSub{A: 1, B: "2"}}

	Prefix = "TEST"

	err := Parse(s, "")
	if err != ErrUndefinedTag {
		t.Fatal("ErrUndefinedTag error expected")
	}

	Tag = "cfg"
	TagDefault = "cfgDefault"

	err = Parse(s, "")
	if err != ErrTypeNotSupported {
		t.Fatal("ErrTypeNotSupported error expected")
	}

	ParseMap[reflect.Int] = reflectIntTestFunc
	ParseMap[reflect.String] = reflectStringTestFunc
	err = Parse(s, "")
	if err != nil {
		t.Fatal(err)
	}

	if s.A != 100 ||
		s.B != "200" ||
		s.S.A != 300 ||
		s.S.B != "400" ||
		s.S.S.A != 500 ||
		s.S.S.B != "600" {
		t.Fatal("Default value error")
	}

	//fmt.Printf("\n\nParse: %#v\n\n", s)

	ParseMap[reflect.Int] = reflectReturnError
	ParseMap[reflect.String] = reflectReturnError
	err = Parse(s, "")
	if err == nil {
		t.Fatal("error expected")
	}

	s1 := "test"
	err = Parse(s1, "")
	if err != ErrNotAPointer {
		t.Fatal("ErrNotAPointer error expected")
	}

	err = Parse(&s1, "")
	if err != ErrNotAStruct {
		t.Fatal("ErrNotAStruct error expected")
	}

	Reset()
	err = Parse(&s1, "")
	if err != ErrNotAStruct {
		t.Fatal("ErrNotAStruct error expected")
	}

}
