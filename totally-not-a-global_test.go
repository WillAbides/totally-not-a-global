package totallynotaglobal_test

import (
	"fmt"

	totallynotaglobal "github.com/willabides/totally-not-a-global"
)

func Example() {
	totallynotaglobal.Set("foo", "bar")
	totallynotaglobal.SetOnce("baz", func() interface{} {
		return "qux"
	})
	totallynotaglobal.SetOnce("baz", func() interface{} {
		return "not qux"
	})

	foo, ok := totallynotaglobal.Get("foo")
	if ok {
		fmt.Print(foo)
	}

	totallynotaglobal.WithTotallyNotAGlobalValue("baz", func(notAGlobalValue interface{}) {
		fmt.Println(notAGlobalValue)
	})
	// output: barqux
}
