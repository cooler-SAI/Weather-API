package tools

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog/log"
)

func StopApp() {

	stop := make(chan os.Signal, 1)

	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop

	log.Info().Msg("Shutting down application gracefully...")
}
