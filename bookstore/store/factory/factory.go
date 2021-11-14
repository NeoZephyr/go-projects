package factory

import (
	"bookstore/store"
	"fmt"
	"sync"
)

var (
	providerMutex sync.RWMutex
	providers = make(map[string]store.Store)
)

func Register(name string, s store.Store) {
	providerMutex.Lock()
	defer providerMutex.Unlock()

	if s == nil {
		panic("store: Register provider is nil")
	}

	if _, dup := providers[name]; dup {
		panic("store: Register twice for provider: " + name)
	}

	providers[name] = s
}

func New(name string) (store.Store, error) {
	providerMutex.RLock()
	s, ok := providers[name]
	providerMutex.RUnlock()

	if !ok {
		return nil, fmt.Errorf("store: Unknow provider: %s", name)
	}

	return s, nil
}