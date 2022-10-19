package methods

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"
)

func GracefulShutdown(log *logrus.Logger) {
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	sig := <-sigChan
	log.Fatalf("Got signal: %v, exiting.", sig)
}
