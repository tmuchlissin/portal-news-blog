package config

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/rs/zerolog/log"
)

func (cfg Config) LoadAwsConfig() aws.Config {
	conf, err := awsConfig.LoadDefaultConfig(context.TODO(), 
	awsConfig.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
		cfg.R2.ApiKey, cfg.R2.ApiSecret, "",
		)), awsConfig.WithRegion("auto"), )
		if err != nil {
			log.Fatal().Err(err).Msg("unable to load AWS config")

}
	log.Info().Msg("AWS config loaded successfully")

	return conf
}