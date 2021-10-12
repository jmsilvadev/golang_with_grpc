package server_test

import (
	"context"
	"os"
	"testing"

	"github.com/jmsilvadev/golangtechtask/api"
	"github.com/jmsilvadev/golangtechtask/pkg/config"
	"github.com/jmsilvadev/golangtechtask/pkg/logger"
	"github.com/jmsilvadev/golangtechtask/pkg/providers"
	"github.com/jmsilvadev/golangtechtask/pkg/server"
)

var s = server.Server{
	Logger: logger.SetLogger(),
}

var ctx = context.Background()
var uuid string

func TestCreateVoteable(t *testing.T) {

	awsConfig := config.AWSConfigConnection{
		Endpoint: os.Getenv("AWS_ENDPOINT"),
		Region:   os.Getenv("AWS_REGION"),
		ID:       os.Getenv("AWS_ACCESS_KEY_ID"),
		Secret:   os.Getenv("AWS_SECRET_ACCESS_KEY"),
	}

	conn, _ := config.GetAwsConfig(ctx, awsConfig)
	timeout := config.GetTimeout("1000ms")

	s.Db = providers.NewDynamoDB(ctx, conn, timeout)

	req := &api.CreateVoteableRequest{
		Question: "Question",
		Answers:  []string{"a", "b", "c"},
	}
	resp, err := s.CreateVoteable(ctx, req)
	if err != nil {
		t.Errorf("Got and Expected are not equals. Got: %v, expected: nil", err.Error())
	} else {
		uuid = resp.UUID
		if resp.UUID == "" {
			t.Errorf("Got and Expected are not equals. Got: %v, expected: any uuid", resp.UUID)
		}
	}
}

func TestListVoteables(t *testing.T) {
	req := &api.ListVoteableRequest{
		PageSize: 1,
	}

	resp, _ := s.ListVoteables(ctx, req)
	if len(resp.Votables) == 0 {
		t.Errorf("Got and Expected are not equals. Got: 0, expected: >0")
	}
}

func TestCastVote(t *testing.T) {
	req := &api.CastVoteRequest{
		UUID:        uuid,
		AnswerIndex: 0,
	}
	resp, _ := s.CastVote(ctx, req)
	if resp.Answer != "a" {
		t.Errorf("Got and Expected are not equals. Got: %v, expected: Answer", resp.Answer)
	}
}
