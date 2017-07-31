//Package goConfig uses a struct as input and populates the
//fields of this struct with parameters fom command
//line, environment variables and configuration file.
package goConfig

import (
	"errors"
	"fmt"
	"path"

	"github.com/crgimenes/goConfig/goEnv"
	"github.com/crgimenes/goConfig/goFlags"
)

// Tag to set main name of field
var Tag = "cfg"

// TagDefault to set default value
var TagDefault = "cfgDefault"

// Path sets default config path
var Path string

// File name of default config file
var File string

// FileRequired config file required
var FileRequired bool

// HelpString temporarily saves help
var HelpString string

// PrefixFlag is a string that would be placed at the beginning of the generated Flag tags.
var PrefixFlag string

// PrefixEnv is a string that would be placed at the beginning of the generated Event tags.
var PrefixEnv string

// ErrFileFormatNotDefined Is the error that is returned when there is no defined configuration file format.
var ErrFileFormatNotDefined = errors.New("file format not defined")

//Usage is a function to show the help, can be replaced by your own version.
var Usage func()

// Fileformat struct holds the functions to Load and Save the file containing the settings
type Fileformat struct {
	Extension   string
	Save        func(config interface{}) (err error)
	Load        func(config interface{}) (err error)
	PrepareHelp func(config interface{}) (help string, err error)
}

// Formats is the list of registered formats.
var Formats []Fileformat

func findFileFormat(extension string) (format Fileformat, err error) {
	format = Fileformat{}
	for _, f := range Formats {
		if f.Extension == extension {
			format = f
			return
		}
	}
	err = ErrFileFormatNotDefined
	return
}

func init() {
	Usage = DefaultUsage
	Path = "./"
	File = ""
	FileRequired = false
}

// Parse configuration
func Parse(config interface{}) (err error) {
	ext := path.Ext(File)
	if ext != "" {
		var format Fileformat
		format, err = findFileFormat(ext)
		if err != nil {
			return
		}
		err = format.Load(config)
		if err != nil {
			return
		}
		HelpString, err = format.PrepareHelp(config)
		if err != nil {
			return
		}
	}

	goEnv.Prefix = PrefixEnv
	goEnv.Setup(Tag, TagDefault)
	err = goEnv.Parse(config)
	if err != nil {
		return
	}

	goFlags.Prefix = PrefixFlag
	goFlags.Setup(Tag, TagDefault)
	goFlags.Usage = Usage
	goFlags.Preserve = true
	err = goFlags.Parse(config)
	if err != nil {
		return
	}

	return
}

// PrintDefaults print the default help
func PrintDefaults() {
	if File != "" {
		fmt.Printf("Config file %q:\n", Path+File)
		fmt.Println(HelpString)
	}
}

// DefaultUsage is assigned for Usage function by default
func DefaultUsage() {
	fmt.Println("Usage")
	goFlags.PrintDefaults()
	goEnv.PrintDefaults()
	PrintDefaults()
}
