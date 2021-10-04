package totallynotaglobal_test

import (
	"fmt"

	totallynotaglobal "github.com/willabides/totally-not-a-global"
)

func Example() {
	totallynotaglobal.Set("foo", "bar")

	foo, ok := totallynotaglobal.Get("foo")
	if ok {
		fmt.Println(foo)
	}
	// output: bar
}
