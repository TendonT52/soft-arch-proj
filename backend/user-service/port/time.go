package port

import "time"

type TimeProvider interface {
	Now() time.Time
}
