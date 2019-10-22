# Strand [![Build Status](https://travis-ci.com/joseluisq/strand.svg?branch=master)](https://travis-ci.com/joseluisq/strand) [![codecov](https://codecov.io/gh/joseluisq/strand/branch/master/graph/badge.svg)](https://codecov.io/gh/joseluisq/strand) [![Go Report Card](https://goreportcard.com/badge/github.com/joseluisq/strand)](https://goreportcard.com/report/github.com/joseluisq/strand) [![GoDoc](https://godoc.org/github.com/joseluisq/strand?status.svg)](https://godoc.org/github.com/joseluisq/strand)

> Tiny package to just generate [secure](https://golang.org/pkg/crypto/rand/) random bytes and strings.

## Supported Go versions

- 1.10.3+
- 1.11+

## Usage

```go
package main

import (
	"fmt"

	"github.com/joseluisq/strand"
)

func main() {
	str, err := strand.randomString(32)

	if err != nil {
		panic(err)
	}

	fmt.Println("Result: ", str)

	// Result: 55b3b071fb00e08460639c05f5ba3ef0
}
```

## Contributions

Unless you explicitly state otherwise, any contribution intentionally submitted for inclusion in current work by you, as defined in the Apache-2.0 license, shall be dual licensed as described below, without any additional terms or conditions.

Feel free to send some [Pull request](https://github.com/joseluisq/strand/pulls) or [issue](https://github.com/joseluisq/strand/issues).

## License

This work is primarily distributed under the terms of both the [MIT license](LICENSE-MIT) and the [Apache License (Version 2.0)](LICENSE-APACHE).

© 2019 [Jose Quintana](http://git.io/joseluisq)
