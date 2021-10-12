package providers_test

import (
	"context"
	"os"
	"testing"

	"github.com/jmsilvadev/golangtechtask/pkg/config"
	"github.com/jmsilvadev/golangtechtask/pkg/providers"
)

func TestNewDynamoDB(t *testing.T) {

	awsConfig := config.AWSConfigConnection{
		Endpoint: os.Getenv("AWS_ENDPOINT"),
		Region:   os.Getenv("AWS_REGION"),
		ID:       os.Getenv("AWS_ACCESS_KEY_ID"),
		Secret:   os.Getenv("AWS_SECRET_ACCESS_KEY"),
	}

	conn, _ := config.GetAwsConfig(context.Background(), awsConfig)
	timeout := config.GetTimeout("1000ms")

	got := providers.NewDynamoDB(context.Background(), conn, timeout)
	if got.TableName != "Voteables" {
		t.Errorf("Got and Expected are not equals. Got: %v, Exp: Votables", got.TableName)
	}
}

func TestNewDynamoDBWithLowTimeout(t *testing.T) {

	awsConfig := config.AWSConfigConnection{
		Endpoint: os.Getenv("AWS_ENDPOINT"),
		Region:   os.Getenv("AWS_REGION"),
		ID:       os.Getenv("AWS_ACCESS_KEY_ID"),
		Secret:   os.Getenv("AWS_SECRET_ACCESS_KEY"),
	}

	conn, _ := config.GetAwsConfig(context.Background(), awsConfig)
	timeout := config.GetTimeout("0ms")

	got := providers.NewDynamoDB(context.Background(), conn, timeout)
	if got.TableName != "Voteables" {
		t.Errorf("Got and Expected are not equals. Got: %v, Exp: Votables", got.TableName)
	}
}
