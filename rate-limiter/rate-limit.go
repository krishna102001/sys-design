package ratelimiter

import (
	"context"

	"github.com/sys-design/rate-limiter/algorithm"
	"github.com/sys-design/rate-limiter/core"
	"github.com/sys-design/rate-limiter/store"
)

type Limiter struct {
	store store.Store
	algo  algorithm.Algorithm
}

// New sets up the limiter.
func New(s store.Store, a algorithm.Algorithm) *Limiter {
	return &Limiter{
		store: s,
		algo:  a,
	}
}

// Check is the single method the user calls to execute the rate limit.
func (l *Limiter) Check(ctx context.Context, key string) core.Result {
	return l.algo.Allow(ctx, l.store, key)
}
