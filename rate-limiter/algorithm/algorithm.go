package algorithm

import (
	"context"

	"github.com/sys-design/rate-limiter/core"
	"github.com/sys-design/rate-limiter/store"
)

type Algorithm interface {
	Allow(ctx context.Context, s store.Store, key string) core.Result
}
