package totallynotaglobal

import "sync"

var (
	initOnce   sync.Once
	lock       sync.RWMutex
	notGlobals map[string]interface{}
)

func initialize() {
	initOnce.Do(
		func() {
			notGlobals = map[string]interface{}{}
		},
	)
}

// SetOnce sets name's value to the resuls of fn if there is no existing value called name.
func SetOnce(name string, fn func() interface{}) {
	initialize()
	lock.Lock()
	defer lock.Unlock()
	if _, ok := notGlobals[name]; ok {
		return
	}
	notGlobals[name] = fn()
}

// Set a value. It's ok, it's totally not a global, at least not in your package.
func Set(name string, value interface{}) {
	initialize()
	lock.Lock()
	notGlobals[name] = value
	lock.Unlock()
}

// Lookup returns a value and whether that value has been Set.
func Lookup(name string) (value interface{}, ok bool) {
	initialize()
	lock.RLock()
	value, ok = notGlobals[name]
	lock.RUnlock()
	return value, ok
}

// Get returns a value
func Get(name string) interface{} {
	initialize()
	lock.RLock()
	value := notGlobals[name]
	lock.RUnlock()
	return value
}

// WithTotallyNotAGlobalValue runs fn with the named value if a value named name has been set
func WithTotallyNotAGlobalValue(name string, fn func(notAGlobalValue interface{})) {
	initialize()
	lock.RLock()
	value, ok := notGlobals[name]
	if ok {
		fn(value)
	}
	lock.RUnlock()
}
