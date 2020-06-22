# Go plugins manager
[![Build Status](https://img.shields.io/travis/clevergo/plugins?style=for-the-badge)](https://travis-ci.org/clevergo/plugins)
[![Coverage Status](https://img.shields.io/coveralls/github/clevergo/plugins?style=for-the-badge)](https://coveralls.io/github/clevergo/plugins)
[![Go.Dev reference](https://img.shields.io/badge/go.dev-reference-blue?logo=go&logoColor=white&style=for-the-badge)](https://pkg.go.dev/clevergo.tech/plugins?tab=doc)
[![Go Report Card](https://goreportcard.com/badge/github.com/clevergo/plugins?style=for-the-badge)](https://goreportcard.com/report/github.com/clevergo/plugins)
[![Release](https://img.shields.io/github/release/clevergo/plugins.svg?style=for-the-badge)](https://github.com/clevergo/plugins/releases)

The `plugins` is a simple Go plugins manager.

## Installation

```shell
$ go get -u clevergo.tech/plugins
```

## Usage

```go
// Plugins location.
path := "/path/to/plugins"
// Creates a plugins manager.
m := plugins.New(path)

// Opens a Go plugins which located at {path}/foo.so
p, err := m.Open("foo.so")

// Lookup a symbol in a plugin.
sym, err := m.Lookup("foo.so", "Bar")
```
