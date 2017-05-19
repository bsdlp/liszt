package registry

import (
	"math/rand"
	"time"

	"github.com/oklog/ulid"
)

func getULID() ulid.ULID {
	return ulid.MustNew(ulid.Now(), rand.New(rand.NewSource(time.Now().UnixNano())))
}
