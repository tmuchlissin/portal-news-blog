package app

import (
	"portal-news-blog/config"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/rs/zerolog/log"
)

func RunServer() {
	cfg := config.NewConfig()
	_, err := cfg.ConnectionPostgres()
	if err != nil {
		log.Fatal().Err(err).Msg("error connecting to database")
	}

	// Cloudflare R2
	cfgR2 := cfg.LoadAwsConfig()
	_ = s3.NewFromConfig(cfgR2)
}
