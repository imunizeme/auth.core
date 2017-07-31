package goConfig

import (
	"errors"
	"os"
	"testing"

	"github.com/crgimenes/goConfig/goFlags"
	"github.com/crgimenes/goConfig/structTag"
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

func TestFindFileFormat(t *testing.T) {
	_, err := findFileFormat(".json")
	if err != ErrFileFormatNotDefined {
		t.Fatal(err)
	}
	Formats = []Fileformat{{Extension: ".json"}}
	_, err = findFileFormat(".json")
	if err != nil {
		t.Fatal(err)
	}
}

// -=-=-=-=-=-=-=-=-=

func mLoad(config interface{}) (err error) {
	return
}

func mSave(config interface{}) (err error) {
	return
}

func mPrepareHelp(config interface{}) (help string, err error) {
	return
}

// -=-=-=-=-=-=-=-=-
func eLoad(config interface{}) (err error) {
	err = errors.New("test")
	return
}

func eSave(config interface{}) (err error) {
	err = errors.New("test")
	return
}

func ePrepareHelp(config interface{}) (help string, err error) {
	err = errors.New("test")
	return
}

// -=-=-=-=-=-=-=-=-

func TestParse(t *testing.T) {

	s := &testStruct{A: 1, S: testSub{A: 1, B: "2"}}
	File = "config.txt"

	Formats = []Fileformat{{Extension: ".json", Save: mSave, Load: mLoad, PrepareHelp: mPrepareHelp}}

	err := Parse(s)
	if err != ErrFileFormatNotDefined {
		t.Fatal("Error ErrFileFormatNotDefined expected")
	}

	File = "config.json"

	Formats = []Fileformat{{Extension: ".json", Save: mSave, Load: eLoad, PrepareHelp: mPrepareHelp}}

	err = Parse(s)
	if err == nil {
		t.Fatal("Error expected")
	}

	Formats = []Fileformat{{Extension: ".json", Save: mSave, Load: mLoad, PrepareHelp: ePrepareHelp}}

	err = Parse(s)
	if err == nil {
		t.Fatal("Error expected")
	}

	Formats = []Fileformat{{Extension: ".json", Save: mSave, Load: mLoad, PrepareHelp: mPrepareHelp}}

	err = os.Setenv("A", "900")
	if err != nil {
		t.Fatal(err)
	}
	err = os.Setenv("B", "TEST")
	if err != nil {
		t.Fatal(err)
	}

	Tag = ""
	err = Parse(s)
	if err != structTag.ErrUndefinedTag {
		t.Fatal("Error structTag.ErrUndefinedTag expected")
	}

	Tag = "cfg"
	err = Parse(s)
	if err != nil {
		t.Fatal(err)
	}

	os.Setenv("A", "900ERROR")

	goFlags.Reset()
	err = Parse(s)
	if err == nil {
		t.Fatal("Error expected")
	}

	err = os.Setenv("A", "")
	if err != nil {
		t.Fatal(err)
	}

	goFlags.Reset()
	err = Parse(s)
	if err != nil {
		t.Fatal(err)
	}

	s1 := "test"
	goFlags.Reset()
	err = Parse(s1)
	if err == nil {
		t.Fatal("Error expected")
	}

	goFlags.Reset()
	err = Parse(&s1)
	if err == nil {
		t.Fatal("Error expected")
	}
}

func ExampleParse() {

	type config struct {
		Name  string `cfg:"Name" cfgDefault:"root"`
		Value int    `cfg:"Value" cfgDefault:"123"`
	}

	cfg := config{}

	err := Parse(&cfg)
	if err != nil {
		println(err)
	}

	println("Name:", cfg.Name, "Value:", cfg.Value)

}
