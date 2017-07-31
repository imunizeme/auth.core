package toml

import (
	"io/ioutil"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/crgimenes/goConfig"
)

func init() {
	f := goConfig.Fileformat{
		Extension:   ".toml",
		Save:        SaveTOML,
		Load:        LoadTOML,
		PrepareHelp: PrepareHelp,
	}
	goConfig.Formats = append(goConfig.Formats, f)
}

// LoadTOML config file
func LoadTOML(config interface{}) (err error) {
	configFile := goConfig.Path + goConfig.File
	_, err = os.Stat(configFile)
	if os.IsNotExist(err) && !goConfig.FileRequired {
		err = nil
		return
	} else if err != nil {
		return
	}

	_, err = toml.DecodeFile(configFile, config)
	return
}

// SaveTOML config file
func SaveTOML(config interface{}) (err error) {
	_, err = os.Stat(goConfig.Path)
	if os.IsNotExist(err) {
		err = os.Mkdir(goConfig.Path, os.ModePerm)
		if err != nil {
			return
		}
	} else if err != nil {
		return
	}

	configFile := goConfig.Path + goConfig.File

	_, err = os.Stat(configFile)
	if err != nil {
		return
	}
	file, err := os.Create(configFile)
	if err != nil {
		return
	}
	defer file.Close()
	err = toml.NewEncoder(file).Encode(config)
	return
}

// PrepareHelp return help string for this file format.
func PrepareHelp(config interface{}) (help string, err error) {
	tmpFile, err := ioutil.TempFile(os.TempDir(), goConfig.File)
	if err != nil {
		return
	}
	defer tmpFile.Close()
	if err = toml.NewEncoder(tmpFile).Encode(config); err != nil {
		return
	}
	helpAux, err := ioutil.ReadAll(tmpFile)
	if err != nil {
		return
	}
	help = string(helpAux)
	return
}
