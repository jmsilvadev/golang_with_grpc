package e2e_test

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jmsilvadev/golangtechtask/api"
	"github.com/jmsilvadev/golangtechtask/pkg/logger"
	"github.com/jmsilvadev/golangtechtask/pkg/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var s = server.Server{
	Logger: logger.SetLogger(),
}

var ctx = context.Background()
var uuid string

const port = ":9000"

var client api.VotingServiceClient
var conn *grpc.ClientConn

func TestCreateVoteable(t *testing.T) {
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = ":4000"
	}

	creds, err := credentials.NewClientTLSFromFile("/certs/server.crt", "")
	if err != nil {
		log.Fatal(err)
	}
	conn, err = grpc.Dial("localhost"+port, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatal(err)
	}

	data := &api.CreateVoteableRequest{
		Question: "Question 1",
		Answers:  []string{"a", "b", "c"},
	}

	client = api.NewVotingServiceClient(conn)
	resp, err := client.CreateVoteable(ctx, data)

	client = api.NewVotingServiceClient(conn)
	resp, err = client.CreateVoteable(ctx, data)

	client = api.NewVotingServiceClient(conn)
	resp, err = client.CreateVoteable(ctx, data)

	client = api.NewVotingServiceClient(conn)
	resp, err = client.CreateVoteable(ctx, data)

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
	req := &api.ListVoteableRequest{}

	resp, err := client.ListVoteables(ctx, req)
	if err != nil {
		t.Errorf("Got and Expected are not equals. Got: %v, expected: nil", err.Error())
	} else {
		if len(resp.Votables) == 0 {
			t.Errorf("Got and Expected are not equals. Got: 0, expected: >0")
		}
	}
}

func TestCastVote(t *testing.T) {
	defer conn.Close()
	req := &api.CastVoteRequest{
		UUID:        uuid,
		AnswerIndex: 0,
	}
	resp, err := client.CastVote(ctx, req)
	if err == nil {
		if resp.Answer != "a" {
			t.Errorf("Got and Expected are not equals. Got: %v, expected: Answer", resp.Answer)
		}
	}
}
