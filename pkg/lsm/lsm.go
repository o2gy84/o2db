package lsm

import (
	"fmt"

	"golang.org/x/xerrors"
)

type LSM struct {
	A int64
	B int64
}

func (lsm LSM) String() string {
	return fmt.Sprintf("TEST XXX: %d, YYY: %d", lsm.A, lsm.B)
}

func New() (LSM, error) {
	var lsm LSM
	return lsm, xerrors.Errorf("test init error")
}
