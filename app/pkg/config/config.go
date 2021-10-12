package config

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/jmsilvadev/golangtechtask/pkg/providers"
	"go.uber.org/zap"
)

// Config portable confiuration of a server
type Config struct {
	Context    context.Context
	ServerPort string
	DBProvider providers.Storage
	Logger     *zap.Logger
}

// AWSConfigConnection provides information to connect with AWS Services
type AWSConfigConnection struct {
	Endpoint string
	Region   string
	ID       string
	Secret   string
}

// NewConfig creates a new custom configuration
func NewConfig(ctx context.Context, port string, db providers.Storage, logger *zap.Logger) *Config {
	return &Config{
		Context:    ctx,
		ServerPort: port,
		DBProvider: db,
		Logger:     logger,
	}
}

// NewAwsConfig creates a new custom AWS configuration
func NewAwsConfig(ctx context.Context, config AWSConfigConnection) (*session.Session, error) {
	return session.NewSessionWithOptions(
		session.Options{
			Config: aws.Config{
				Credentials:      credentials.NewStaticCredentials(config.ID, config.Secret, ""),
				Region:           aws.String(config.Region),
				Endpoint:         aws.String(config.Endpoint),
				S3ForcePathStyle: aws.Bool(true),
			},
		},
	)
}
