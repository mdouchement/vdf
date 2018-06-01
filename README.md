# Valve Data Format

[![CircleCI](https://circleci.com/gh/mdouchement/vdf.svg?style=shield)](https://circleci.com/gh/mdouchement/vdf)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg)](https://godoc.org/github.com/mdouchement/vdf)
[![Go Report Card](https://goreportcard.com/badge/github.com/mdouchement/vdf)](https://goreportcard.com/report/github.com/mdouchement/vdf)
[![License](https://img.shields.io/github/license/mdouchement/vdf.svg)](http://opensource.org/licenses/MIT)

A parser and a generator for [Valve Data Format](https://developer.valvesoftware.com/wiki/KeyValues) written in Go. It does not support all the features provided in Valve Data Format spec.

Supported files (at least):
- remotecache.vdf

## Usage

```go
package main

import (
	"io/ioutil"

	"github.com/mdouchement/vdf"
	"github.com/sanity-io/litter"
)

func main() {
	data, err := ioutil.ReadFile("remotecache.vdf")
	check(err)

	root, err := vdf.Parse(string(data))
	check(err)

	litter.Dump(root)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
```

## License

**MIT**


## Contributing

All PRs are welcome.

1. Fork it
2. Create your feature branch (git checkout -b my-new-feature)
3. Commit your changes (git commit -am 'Add some feature')
5. Push to the branch (git push origin my-new-feature)
6. Create new Pull Request

As possible, run the following commands to format and lint the code:

```sh
# Format
find . -name '*.go' -not -path './vendor*' -exec gofmt -s -w {} \;

# Lint
gometalinter --config=gometalinter.json ./...
```
