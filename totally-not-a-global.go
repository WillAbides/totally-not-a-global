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

// Set a value. It's ok, it's totally not a global, at least not in your package.
func Set(name string, value interface{}) {
	initialize()
	lock.Lock()
	defer lock.Unlock()
	notGlobals[name] = value
}

// Get returns a value and whether that value has been Set.
func Get(name string) (value interface{}, ok bool) {
	initialize()
	lock.RLock()
	defer lock.RUnlock()
	value, ok = notGlobals[name]
	return value, ok
}
