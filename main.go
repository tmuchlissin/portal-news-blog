package main

import (
	"portal-news-blog/cmd"

	"github.com/rs/zerolog/log"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal().Err(err).Msg("failed to execute command")
	}
}
