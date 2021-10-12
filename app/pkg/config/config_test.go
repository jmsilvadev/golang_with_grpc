package config

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	"github.com/jmsilvadev/golangtechtask/pkg/providers"
	"go.uber.org/zap"
)

func TestNewConfig(t *testing.T) {
	got := NewConfig(context.Background(), ":4000", providers.NewMockDB(context.Background()), &zap.Logger{})
	if got.ServerPort != ":4000" {
		t.Errorf("Got and Expected are not equals. Got: %v, expected: :4000", got.ServerPort)
	}
}

func TestNewAwsConfig(t *testing.T) {

	awsConfig := AWSConfigConnection{
		Endpoint: "endpoint",
		Region:   "region",
		ID:       "id",
		Secret:   "secret",
	}

	_, err := NewAwsConfig(context.Background(), awsConfig)
	if err != nil {
		t.Errorf("Got and Expected are not equals. Got: %v, expected: nil", err.Error())
	}
}

func TestGetAwsConfig(t *testing.T) {
	awsConfig := AWSConfigConnection{
		Endpoint: "endpoint",
		Region:   "region",
		ID:       "id",
	}
	_, err := GetAwsConfig(context.Background(), awsConfig)
	if err != nil {
		t.Errorf("Got and Expected are not equals. Got: %v, nil", err.Error())
	}
}

func TestGetTimeout(t *testing.T) {
	duration := 100 * time.Millisecond
	got := GetTimeout("100")
	if got != duration {
		t.Errorf("Got and Expected are not equals. Got: %v, %v", got, duration)
	}
}

func TestGetTimeoutWithError(t *testing.T) {
	duration := 100 * time.Millisecond
	got := GetTimeout("")
	if got != duration {
		t.Errorf("Got and Expected are not equals. Got: %v, %v", got, duration)
	}
}

func TestTransformConfigStruct(t *testing.T) {

	configJSON := `
	{
		"ServerPort": ":4000"
	}
	`
	data := &Config{}
	err := json.Unmarshal([]byte(configJSON), data)
	if err != nil {
		t.Errorf("Got and Expected are not equals.\n got: %v, expected: %v ",
			err, nil)
	}
}

func TestTransformAWSConfigConnection(t *testing.T) {
	configJSON := `
	{
		"Endpoint": "Endpoint",
		"Region": "Region",
		"ID": "ID",
		"Secret": "Secret"
	}
	`
	data := &Config{}
	err := json.Unmarshal([]byte(configJSON), data)
	if err != nil {
		t.Errorf("Got and Expected are not equals.\n got: %v, expected: %v ",
			err, nil)
	}
}

func TestGetDeaultConfig(t *testing.T) {
	config := GetDeaultConfig()
	if config.ServerPort == "" {
		t.Errorf("Got and Expected are not equals. got: '', expected: !''")
	}
}
