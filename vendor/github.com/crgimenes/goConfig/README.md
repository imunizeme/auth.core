# goConfig
[![Build Status](https://travis-ci.org/crgimenes/goConfig.svg?branch=master)](https://travis-ci.org/crgimenes/goConfig)
[![Go Report Card](https://goreportcard.com/badge/github.com/crgimenes/goConfig)](https://goreportcard.com/report/github.com/crgimenes/goConfig)
[![codecov](https://codecov.io/gh/crgimenes/goConfig/branch/master/graph/badge.svg)](https://codecov.io/gh/crgimenes/goConfig)
[![GoDoc](https://godoc.org/github.com/crgimenes/goConfig?status.png)](https://godoc.org/github.com/crgimenes/goConfig)
[![Go project version](https://badge.fury.io/go/github.com%2Fcrgimenes%2FgoConfig.svg)](https://badge.fury.io/go/github.com%2Fcrgimenes%2FgoConfig)
[![MIT Licensed](https://img.shields.io/badge/license-MIT-green.svg)](https://tldrlegal.com/license/mit-license)


goConfig uses a struct as input and populates the fields of this struct with parameters fom command line, environment variables and configuration file.

![goConfig simple anda fast config package](http://crg.eti.br/img/goConfig.gif)

## Install

```
go get github.com/crgimenes/goConfig
```

## Example

```go
package main

import "github.com/crgimenes/goConfig"

/*
step 1: Declare your configuration struct,
it may or may not contain substructures.
*/

type mongoDB struct {
	Host string `cfgDefault:"example.com"`
	Port int    `cfgDefault:"999"`
}

type configTest struct {
	Domain    string
	DebugMode bool `json:"db" cfg:"db" cfgDefault:"false"`
	MongoDB   mongoDB
}

func main() {

	// step 2: Instantiate your structure.
	config := configTest{}

	// step 3: Pass the instance pointer to the parser
	err := goConfig.Parse(&config)
	if err != nil {
		println(err)
		return
	}

	/*
	   The parser populated your struct with the data
	   it took from environment variables and command
	   line and now you can use it.
	*/

	println("config.Domain......:", config.Domain)
	println("config.DebugMode...:", config.DebugMode)
	println("config.MongoDB.Host:", config.MongoDB.Host)
	println("config.MongoDB.Port:", config.MongoDB.Port)
}
```

With the example above try environment variables like *$DOMAIN* or *$MONGODB_HOST* and run the example again to see what happens.

You can also try using parameters on the command line, try -h to see the help.

## Contributing

- Fork the repo on GitHub
- Clone the project to your own machine
- Create a *branch* with your modifications `git checkout -b fantastic-feature`.
- Then _commit_ your changes `git commit -m 'Implementation of new fantastic feature'`
- Make a _push_ to your _branch_ `git push origin fantastic-feature`.
- Submit a **Pull Request** so that we can review your changes
