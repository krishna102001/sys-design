package algorithm

import (
	"context"
	_ "embed"
	"time"

	"github.com/sys-design/rate-limiter/core"
	"github.com/sys-design/rate-limiter/store"
)

//go:embed scripts/leaky-bucket.lua
var leakyBucketScript string

type LeakyBucket struct {
	Capacity int
	LeakRate float64
}

func NewLeakyBucket(capacity int, leaky_rate float64) *LeakyBucket {
	return &LeakyBucket{
		Capacity: capacity,
		LeakRate: leaky_rate,
	}
}

func (lb *LeakyBucket) Allow(ctx context.Context, store store.Store, key string) core.Result {
	args := []any{lb.Capacity, lb.LeakRate, time.Now().Unix()}

	rawResult, err := store.Execute(ctx, "leaky_bucket", leakyBucketScript, []string{key}, args...)
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
