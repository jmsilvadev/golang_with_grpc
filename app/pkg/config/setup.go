package config

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/jmsilvadev/golangtechtask/pkg/logger"
	"github.com/jmsilvadev/golangtechtask/pkg/providers"
)

// GetAwsConfig gets the current AWS configuration
func GetAwsConfig(ctx context.Context, awsConfig AWSConfigConnection) (*session.Session, error) {

	conn, err := NewAwsConfig(ctx, awsConfig)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

// GetTimeout gets the current timeout duration to wait for database response
func GetTimeout(timeout string) time.Duration {
	duration, err := time.ParseDuration(timeout)
	if err != nil {
		duration = 100 * time.Millisecond
		log.Println("Timeout info not found, setting default: ", duration)
	}

	return duration
}

// GetDeaultConfig gets the default configuration based in the environment variables
func GetDeaultConfig() *Config {
	ctx := context.Background()

	logger := logger.SetLogger()

	awsConfig := AWSConfigConnection{
		Endpoint: os.Getenv("AWS_ENDPOINT"),
		Region:   os.Getenv("AWS_REGION"),
		ID:       os.Getenv("AWS_ACCESS_KEY_ID"),
		Secret:   os.Getenv("AWS_SECRET_ACCESS_KEY"),
	}
	conn, err := GetAwsConfig(ctx, awsConfig)
	if err != nil {
		logger.Fatal("Invalid Aws Configuration")
	}

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = ":4000"
	}

	timeout := GetTimeout(os.Getenv("TIMEOUT"))
	db := providers.NewDynamoDB(ctx, conn, timeout)
	config := NewConfig(ctx, port, db, logger)

	return config
}
