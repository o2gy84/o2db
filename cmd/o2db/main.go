package main

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/o2gy84/o2db/pkg/lsm"
	"github.com/o2gy84/o2db/pkg/commit"

)

func main() {

	log.Printf("[commit] %s", commit.Commit)

	var lsm lsm.LSM
	fmt.Printf("lsm is: %s\n", lsm)
}
