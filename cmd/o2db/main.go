package main

import (
	log "github.com/sirupsen/logrus"

	"github.com/o2gy84/o2db/pkg/lsm"
	"github.com/o2gy84/o2db/pkg/commit"

)

func main() {

	log.Printf("[commit] %s", commit.Commit)

	lsm, err := lsm.New()
	if err != nil {
		log.Fatalf("lsm.Init() error: %s", err)
	}
	log.Printf("lsm is: %s", lsm)

}
