package algorithm

import (
	"context"
	_ "embed"
	"time"

	"github.com/sys-design/rate-limiter/core"
	"github.com/sys-design/rate-limiter/store"
)

//go:embed scripts/fixed-window-counter.lua
var fixedWindowCounterScript string

type FixedWindowCounter struct {
	Limit      int
	WindowSize float64
}

func NewFixedWindowCounter(limit int, window_size float64) *FixedWindowCounter {
	return &FixedWindowCounter{
		Limit:      limit,
		WindowSize: window_size,
	}
}

func (fwc *FixedWindowCounter) Allow(ctx context.Context, store store.Store, key string) core.Result {
	args := []any{fwc.Limit, fwc.WindowSize, time.Now().Unix()}

	rawResult, err := store.Execute(ctx, "fixed_window_counter", fixedWindowCounterScript, []string{key}, args...)
	if err != nil {
		return core.Result{Error: err}
	}

	resArray := rawResult.([]interface{})
	return core.Result{
		Allowed:   resArray[0].(int64) == 1,
		Remaining: int(resArray[1].(int64)),
		ResetAt:   time.Unix(resArray[2].(int64), 0),
	}
}
