package lsm

import (
	"fmt"
)

type LSM struct {
	A int64
}

func (lsm LSM) String() string {
	return fmt.Sprintf("TEST XXX: %d", lsm.A)
}
