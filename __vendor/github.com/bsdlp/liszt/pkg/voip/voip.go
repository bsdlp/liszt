package voip

import "context"

// Caller makes calls
type Caller interface {
	CallUnit(ctx context.Context, unitID string) (err error)
	CallResident(ctx context.Context, residentID string) (err error)
}
