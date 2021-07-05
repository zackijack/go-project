package helpers

import (
	"fmt"
	zlog "github.com/rs/zerolog/log"
	"os"
)

func CheckErr(err error, msg string, exit bool) {

	if err != nil {
		zlog.Error().Err(err).Msg(msg)

		if exit {
			os.Exit(1)
		}
	}
}

func ErrMsg(s string) string {
	return fmt.Sprintf("There is something wrong with the %s", s)
}
