package core

import "time"

// Result will contain the outcome of the ratelimit check
type Result struct {
	Allowed   bool      // whether request should paased or stop
	Remaining int       // requests remaining in the current window
	ResetAt   time.Time // when the limit resets
	Error     error     // capture store failure for example redis down
}
