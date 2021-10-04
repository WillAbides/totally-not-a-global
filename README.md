# totally-not-a-global

[![godoc](https://pkg.go.dev/badge/github.com/willabides/totally-not-a-global.svg)](https://pkg.go.dev/github.com/willabides/totally-not-a-global)
[![ci](https://github.com/WillAbides/totally-not-a-global/workflows/ci/badge.svg?branch=main&event=push)](https://github.com/WillAbides/totally-not-a-global/actions?query=workflow%3Aci+branch%3Amain+event%3Apush)

When you don't want to sully your package with globals, don't worry because `totally-not-a-global` will sully itself 
instead.

## Usage

```golang
package main

import (
	"fmt"

	totallynotaglobal "github.com/willabides/totally-not-a-global"
)

func main() {
	totallynotaglobal.Set("foo", "bar")

	foo, ok := totallynotaglobal.Get("foo")
	if ok {
		fmt.Print(foo)
	}

	totallynotaglobal.WithTotallyNotAGlobalValue("foo", func(notAGlobalValue interface{}) {
		fmt.Println(notAGlobalValue)
	})
}
```
