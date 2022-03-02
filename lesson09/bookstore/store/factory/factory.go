package factory

import (
	"fmt"
	"sync"

	"github.com/aseara/baigo/lesson09/bookstore/store"
)

var (
	providersMu sync.RWMutex
	providers   = make(map[string]store.Store)
)

// Register register
func Register(name string, p store.Store) {
	providersMu.Lock()
	defer providersMu.Unlock()
	if p == nil {
		panic("store: Register provider is nil")
	}

	if _, dup := providers[name]; dup {
		panic("store: Register called twice for provider " + name)
	}
	providers[name] = p
}

// New new
func New(providerName string) (store.Store, error) {
	providersMu.RLock()
	p, ok := providers[providerName]
	providersMu.RUnlock()
	if !ok {
		return nil, fmt.Errorf("store: unknown provider %s", providerName)
	}

	return p, nil
}
