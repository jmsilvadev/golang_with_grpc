package storage_test

import (
	"context"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/jmsilvadev/golangtechtask/pkg/config"
	"github.com/jmsilvadev/golangtechtask/pkg/internal/entities"
	"github.com/jmsilvadev/golangtechtask/pkg/logger"
	"github.com/jmsilvadev/golangtechtask/pkg/providers"
	"github.com/jmsilvadev/golangtechtask/pkg/server"
)

var s = server.Server{
	Logger: logger.SetLogger(),
}

var ctx = context.Background()
var db providers.Storage
var newUuid string

func TestHasTable(t *testing.T) {
	awsConfig := config.AWSConfigConnection{
		Endpoint: os.Getenv("AWS_ENDPOINT"),
		Region:   os.Getenv("AWS_REGION"),
		ID:       os.Getenv("AWS_ACCESS_KEY_ID"),
		Secret:   os.Getenv("AWS_SECRET_ACCESS_KEY"),
	}

	conn, _ := config.GetAwsConfig(ctx, awsConfig)
	timeout := config.GetTimeout("1000ms")

	db = providers.NewDynamoDB(ctx, conn, timeout)
	ok, _ := db.HasTables(context.Background())
	if !ok {
		err := db.CreateTables(context.Background())
		if err != nil {
			t.Errorf("Got and Expected are not equals. Got: !nil, expected: nil")
		}
	} else {
		err := db.CreateTables(context.Background())
		if err == nil {
			t.Errorf("Got and Expected are not equals. Got: nil, expected: !nil")
		}
	}
}
func TestCreate(t *testing.T) {
	votes := map[string]int64{}
	votes["0"] = 0
	votes["1"] = 0
	votes["2"] = 0
	newUuid = uuid.NewString()
	req := &entities.Voteable{
		UUID:     newUuid,
		Question: "Question",
		Answers:  []string{"a", "b", "c"},
		Votes:    votes,
	}
	resp, err := db.Create(context.Background(), req)
	if err != nil {
		t.Errorf("Error occurred %v. Uuid: %v", err.Error(), newUuid)
	} else {
		if resp.UUID != newUuid {
			t.Errorf("Got and Expected are not equals. Got: %v, expected: uuid", resp.UUID)
		}
	}
}

func TestListAllVoteables(t *testing.T) {

	insertVoteable()
	insertVoteable()

	req := &entities.ListVoteableRequest{
		PageSize: 2,
	}
	resp, _ := db.List(context.Background(), req)
	if len(resp.Votables) == 0 {
		t.Errorf("Got and Expected are not equals. Got: %v, expected: > 0", len(resp.Votables))
	}

	req = &entities.ListVoteableRequest{
		PageSize: 2,
		Page:     resp.NextPage,
	}
	resp, _ = db.List(context.Background(), req)
	if len(resp.Votables) == 0 {
		t.Errorf("Got and Expected are not equals. Got: %v, expected: > 0", len(resp.Votables))
	}

	req = &entities.ListVoteableRequest{}
	resp, _ = db.List(context.Background(), req)
	if len(resp.Votables) == 0 {
		t.Errorf("Got and Expected are not equals. Got: %v, expected: > 0", len(resp.Votables))
	}

	for _, v := range resp.Votables {
		if len(v.Votes) > 0 {
			newUuid = v.UUID
		}
	}
}

func TestUpdateCastVote(t *testing.T) {
	insertVoteable()
	req := &entities.CastVoteRequest{
		UUID:        newUuid,
		AnswerIndex: 0,
	}
	resp, err := db.Update(context.Background(), req)
	if err != nil {
		t.Errorf("Error occurred %v. Uuid: %v", err.Error(), newUuid)
	} else {
		if resp.Answer != "a" {
			t.Errorf("Got and Expected are not equals. Got: %v, expected: a", resp)
		}

	}

	req = &entities.CastVoteRequest{
		UUID:        newUuid,
		AnswerIndex: 100,
	}
	resp, err = db.Update(context.Background(), req)
	if err.Error() != "Answer index not found" {
		t.Errorf("Got and Expected are not equals. Got: %v, expected: Answer index not found", err.Error())
	}

	req = &entities.CastVoteRequest{
		UUID:        uuid.NewString(),
		AnswerIndex: 0,
	}
	resp, err = db.Update(context.Background(), req)
	if err.Error() != "Not Found" {
		t.Errorf("Got and Expected are not equals. Got: %v, expected: Not Found", err.Error())
	}
}

func insertVoteable() {
	votes := map[string]int64{}
	votes["0"] = 0
	votes["1"] = 0
	votes["2"] = 0
	newUuid = uuid.NewString()
	req := &entities.Voteable{
		UUID:     newUuid,
		Question: "Question",
		Answers:  []string{"a", "b", "c"},
		Votes:    votes,
	}
	db.Create(context.Background(), req)
}
