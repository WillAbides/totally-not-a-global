package totallynotaglobal_test

import (
	"fmt"

	totallynotaglobal "github.com/willabides/totally-not-a-global"
)

func Example() {
	totallynotaglobal.Set("foo", "bar")

	foo, ok := totallynotaglobal.Get("foo")
	if ok {
		fmt.Print(foo)
	}

	totallynotaglobal.WithTotallyNotAGlobalValue("foo", func(notAGlobalValue interface{}) {
		fmt.Println(notAGlobalValue)
	})
	// output: barbar
}
