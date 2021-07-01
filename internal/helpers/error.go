package helpers

import (
	"github.com/rs/zerolog/log"
	"os"
)

func CheckErr(err error, msg string, exit bool) {

	if err != nil {
		log.Error().Err(err).Msg(msg)

		if exit {
			os.Exit(1)
		}
	}
}
