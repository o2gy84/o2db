package cmd

import (
	"context"
	"log/syslog"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/o2gy84/o2db/config"
)

func NewLogger(cfg *config.Logger, tag string) *logrus.Logger {
	return setupLogger(logrus.New(), cfg, tag)
}

func setupLogger(log *logrus.Logger, cfg *config.Logger, tag string) *logrus.Logger {
	if log == nil {
		log = logrus.StandardLogger()
	}
	lvl, err := logrus.ParseLevel(*cfg.Level)
	if err != nil {
		log.WithError(err).Panic("Failed to parse log level")
	}
	log.SetLevel(lvl)
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   *cfg.FullTimestamps,
		TimestampFormat: *cfg.TimestampFormat,
	})

	if *cfg.Syslog {
		syslogWriter, err := syslog.New(syslog.LOG_USER|syslog.LOG_INFO, tag)
		if err != nil {
			log.WithError(err).Panic("Failed to init syslog")
		}
		log.SetOutput(syslogWriter)
		log.SetFormatter(&logrus.TextFormatter{
			DisableTimestamp: true,
			DisableColors:    true,
		})
	}
	return log
}

func SetupLog(cfg *config.Logger, tag string) *logrus.Logger {
	return setupLogger(logrus.StandardLogger(), cfg, tag)
}

func Seed() {
	rand.Seed(time.Now().UnixNano())
}

type RunOpt struct {
	Timeout  time.Duration // graceful shutdown hard timeout
	ExitWait time.Duration // wait before exit
}

func Run(ctx context.Context, opt RunOpt, f func(ctx context.Context) error) {
	// Listening for ^C or SIGTERM signal that will start graceful shutdown.
	runCtx, cancel := context.WithCancel(ctx)

	done := make(chan struct{})
	go func() {
		defer close(done)
		if err := f(runCtx); err != context.Canceled {
			logrus.WithError(err).Fatal(err)
		}
	}()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)
	s := <-signals
	logrus.Infof("Stopping on signal %s", s)
	cancel()
	forcedCancel := make(chan struct{})
	go func() {
		// Handling second signal as forced cancel.
		<-signals
		close(forcedCancel)
	}()

	select {
	case <-done:
		logrus.Info("Graceful shutdown OK")
	case <-time.After(opt.Timeout):
		logrus.Fatal("Graceful shutdown failed (timed out)")
	case <-forcedCancel:
		logrus.Warn("Shutting down (forced)")
	}
	if opt.ExitWait > 0 {
		logrus.Info("Waiting before exit")
		select {
		case <-time.After(opt.ExitWait):
			logrus.Info("Exiting")
		case <-forcedCancel:
			logrus.Warn("Not waiting before exit (forced)")
		}
	}
}
