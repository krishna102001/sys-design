package store

import "context"

// its define the contract for backend for any database
type Store interface {
	// Execute runs an atomic operation.
	// scriptName helps in-memory stores know which Go logic to run,
	// while scriptBody is the actual Lua code for Redis.
	Execute(ctx context.Context, scriptName, scriptBody string, keys []string, args ...any) (any, error)
}
