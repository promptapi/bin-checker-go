![Version](https://img.shields.io/badge/version-0.1.0-orange.svg)
![Go](https://img.shields.io/badge/go-1.15.1-black.svg)
[![Documentation](https://godoc.org/github.com/promptapi/bin-checker-go?status.svg)](https://pkg.go.dev/github.com/promptapi/bin-checker-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/promptapi/bin-checker-go)](https://goreportcard.com/report/github.com/promptapi/bin-checker-go)
[![Build Status](https://travis-ci.org/promptapi/bin-checker-go.svg?branch=main)](https://travis-ci.org/promptapi/bin-checker-go)

# Prompt API - BIN Checker - Golang Package

`binchecker` is a simple golang wrapper for [bincheck-api][bincheck-api].

## Requirements

1. You need to signup for [Prompt API][promptapi-signup]
1. You need to subscribe [bincheck-api][bincheck-api], test drive is **free!!!**
1. You need to set `PROMPTAPI_TOKEN` environment variable after subscription.

---

## Installation

```go
$ go get -u github.com/promptapi/bin-checker-go
```

## Example Usage

```go
// main.go

package main

import (
	"fmt"
	"os"

	binchecker "github.com/promptapi/bin-checker-go"
)

func main() {
	result := new(binchecker.Result)

	if err := binchecker.BinChecker("302596", result); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Printf("%+v\n", result)

	fmt.Printf("BankName: %s\n", result.BankName)
	fmt.Printf("Country: %s\n", result.Country)
	fmt.Printf("URL: %s\n", result.URL)
	fmt.Printf("Type: %s\n", result.Type)
	fmt.Printf("Scheme: %s\n", result.Scheme)
	fmt.Printf("Bin: %s\n", result.Bin)
}
```

Run from cli:

```bash
$ PROMPTAPI_TOKEN="YOUR-PROMPTAPI-TOKEN" go run main.go

&{BankName:Diners Club International Country:United States Of America URL:www.dinersclub.com Type:Credit Scheme:Discover Bin:302596}
BankName: Diners Club International
Country: United States Of America
URL: www.dinersclub.com
Type: Credit
Scheme: Discover
Bin: 302596
```

---

## Development

Create a fake web-server under `/tmp/`;

```bash
$ mkdir /tmp/server/
$ cp test.json /tmp/server/
$ cd /tmp/server/
$ python -m http.server 8000
```

And use:

```bash
$ PROMPTAPI_TOKEN="fake" \
  PROMPTAPI_TEST_ENDPOINT="http://localhost:8000/test.json" \
  go run main.go
```

User rake tasks:

```bash
$ rake

rake default                    # Default task, show avaliable tasks
rake release:check              # Do release check
rake release:publish[revision]  # Publish project with revision: major,minor,patch, default: patch
rake serve_doc[port]            # Run doc server
rake test[verbose]              # Run tests
```

- Run tests: `rake test` or `rake test[-v]`
- Run doc server: `rake serve_doc` or `rake serve_doc[9000]`

Release package (*if you have write access*):

1. Commit your changes
1. Run `rake release:check`
1. If all goes ok, run `rake release:publish`

---

## License

This project is licensed under MIT

---

## Contributer(s)

* [Prompt API](https://github.com/promptapi) - Creator, maintainer

---

## Contribute

All PR’s are welcome!

1. `fork` (https://github.com/promptapi/bin-checker-go/fork)
1. Create your `branch` (`git checkout -b my-feature`)
1. `commit` yours (`git commit -am 'Add awesome features...'`)
1. `push` your `branch` (`git push origin my-feature`)
1. Than create a new **Pull Request**!

This project is intended to be a safe,
welcoming space for collaboration, and contributors are expected to adhere to
the [code of conduct][coc].


---

[bincheck-api]:     https://promptapi.com/marketplace/description/bincheck-api
[promptapi-signup]: https://promptapi.com/#signup-form
[coc]:              https://github.com/promptapi/bin-checker-go/blob/main/CODE_OF_CONDUCT.md
