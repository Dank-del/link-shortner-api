package types

import "time"

type Error struct {
	ErrData    error
	OccouredAt time.Time
}
