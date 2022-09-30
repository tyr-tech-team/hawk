package status

import (
	"log"
	"sync"
)

// List -
type List struct {
	mu         sync.RWMutex
	statusList map[string]Status
}

var list = &List{statusList: make(map[string]Status)}

// Set -
func (l *List) Set(code string, s Status) {
	if _, ok := l.Get(code); ok {
		log.Fatalf("Have the same error code [%s]", code)
	}

	l.statusList[code] = s
}

// Get -
func (l *List) Get(code string) (Status, bool) {
	s, ok := l.statusList[code]
	if ok {
		return s, ok
	}
	return UnKnownError, false
}
