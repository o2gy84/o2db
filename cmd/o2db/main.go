package main

import (
	"time"
	"context"

	log "github.com/sirupsen/logrus"

	"github.com/o2gy84/o2db/pkg/lsm"
	"github.com/o2gy84/o2db/pkg/commit"
	"github.com/o2gy84/o2db/pkg/cmd"
)


func run(context context.Context) error {
	for {
		log.Printf("hi")
		time.Sleep(1000*time.Millisecond)
	}
	return nil
}

func main() {
	log.Printf("[commit] %s", commit.Commit)

	lsm, err := lsm.New()
	if err != nil {
		log.Errorf("lsm.Init() error: %s", err)
	}
	log.Printf("lsm is: %s", lsm)

	ctx := context.Background()
	opt := cmd.RunOpt{
		Timeout: 100*time.Millisecond,
		ExitWait: 100*time.Millisecond,
	}

	cmd.Run(ctx, opt, run)
}
