package lsm

import (
	"fmt"
)

type LSM struct {
	A int64
	B int64
}

func (lsm LSM) String() string {
	return fmt.Sprintf("TEST XXX: %d, YYY: %d", lsm.A, lsm.B)
}
