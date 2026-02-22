package algorithm

import (
	"context"
	_ "embed"
	"time"

	"github.com/sys-design/rate-limiter/core"
	"github.com/sys-design/rate-limiter/store"
)

//go:embed scripts/token-bucket.lua
var tokenBucketScript string

type TokenBucket struct {
	Capacity   int
	RefillRate float64
}

func NewTokentBucket(capacity int, refillRate float64) *TokenBucket {
	return &TokenBucket{Capacity: capacity, RefillRate: refillRate}
}

func (tb *TokenBucket) Allow(ctx context.Context, s store.Store, key string) core.Result {
	args := []any{tb.Capacity, tb.RefillRate, time.Now().Unix()}

	// Pass both the name (for memory) and the script (for Redis)
	rawResult, err := s.Execute(ctx, "token_bucket", tokenBucketScript, []string{key}, args...)
	if err != nil {
		return core.Result{Error: err}
	}

	// Parse the raw []interface{} array returned by Redis
	resArray := rawResult.([]interface{})
	return core.Result{
		Allowed:   resArray[0].(int64) == 1,
		Remaining: int(resArray[1].(int64)),
		ResetAt:   time.Unix(resArray[2].(int64), 0),
	}
}
