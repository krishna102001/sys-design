package algorithm

import (
	"context"
	_ "embed"
	"time"

	"github.com/google/uuid"
	"github.com/sys-design/rate-limiter/core"
	"github.com/sys-design/rate-limiter/store"
)

//go:embed scripts/sliding-window-log.lua
var slidingWindowLogScript string

type SlidingWindowLog struct {
	Limit      int64
	WindowSize float64
}

func NewSlidingWindowLog(limit int64, window_size float64) *SlidingWindowLog {
	return &SlidingWindowLog{
		Limit:      limit,
		WindowSize: window_size,
	}
}

func (swl *SlidingWindowLog) Allow(ctx context.Context, store store.Store, key string) core.Result {
	args := []any{swl.Limit, swl.WindowSize, time.Now().Unix(), uuid.New()}

	rawResult, err := store.Execute(ctx, "sliding_window_log", slidingWindowLogScript, []string{key}, args...)
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
