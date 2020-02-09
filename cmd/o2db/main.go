package main

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/o2gy84/o2db/pkg/cmd"
	"github.com/o2gy84/o2db/pkg/commit"
	"github.com/o2gy84/o2db/pkg/lsm"
)

func run(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			log.Printf("context is canceled")
			return nil
		default:
			log.Printf("hi")
			time.Sleep(1000 * time.Millisecond)
		}
	}
}

func main() {
	log.Printf("[commit] %s", commit.Commit)

	l, err := lsm.New()
	if err != nil {
		log.Errorf("lsm.Init() error: %s", err)
	}
	log.Printf("lsm is: %s", l)

	ctx := context.Background()
	opt := cmd.RunOpt{
		Timeout:  100 * time.Millisecond,
		ExitWait: 100 * time.Millisecond,
	}

	cmd.Run(ctx, opt, run)
}
