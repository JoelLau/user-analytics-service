package server

import "time"

type Nower interface {
	Now() time.Time
}
