package store

import (
	"context"
	"fmt"
	"sync"
)

type MemoryStore struct {
	mu   sync.Mutex
	data map[string]any
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		data: make(map[string]any),
	}
}

func (m *MemoryStore) Execute(ctx context.Context, scriptName, scriptBody string, keys []string, args ...any) (any, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	// What to keep in mind: In-memory stores can't execute Lua natively!
	// You would use a switch statement on 'scriptName' here to trigger
	// the equivalent Golang logic for TokenBucket, LeakyBucket, etc.
	switch scriptName {
	case "token_bucket":
		// TODO: Implement local token bucket math using m.data
		return []interface{}{int64(1), int64(9), int64(1700000000)}, nil
	default:
		return nil, fmt.Errorf("unsupported local script: %s", scriptName)
	}
}
