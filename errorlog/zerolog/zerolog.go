package zerolog

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func zeroLog() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Print("Hey, I'm a log message!")
}

func zeroLog1() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	log.Debug().Int("EmployeeId", 1001).Msg("Getting employee information")
	log.Debug().Str("Name", "John").Send()
}
