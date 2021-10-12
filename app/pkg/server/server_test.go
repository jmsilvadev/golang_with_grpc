package server

import (
	"context"
	"testing"

	"github.com/jmsilvadev/golangtechtask/api"
	"github.com/jmsilvadev/golangtechtask/pkg/logger"
	"github.com/jmsilvadev/golangtechtask/pkg/providers"
)

var s = Server{
	Db:     providers.NewMockDB(ctx),
	Logger: logger.SetLogger(),
}

var ctx = context.Background()

func TestListVoteables(t *testing.T) {
	req := &api.ListVoteableRequest{
		PageSize: 1,
		Page:     "Page",
	}

	resp, _ := s.ListVoteables(ctx, req)
	if resp.Page != "Page" {
		t.Errorf("Got and Expected are not equals. Got: %v, expected: Page", resp.Page)
	}

	req = &api.ListVoteableRequest{
		PageSize: 1,
		Page:     "Error",
	}
	_, err := s.ListVoteables(ctx, req)
	if err == nil {
		t.Errorf("Got and Expected are not equals. Got: nil, expected: Error")
	}

	req = &api.ListVoteableRequest{
		PageSize: 1,
		Page:     "Nil",
	}
	resp, err = s.ListVoteables(ctx, req)
	if err != nil {
		t.Errorf("Got and Expected are not equals. Got: nil, expected: Error")
	}
}

func TestCreateVoteable(t *testing.T) {
	req := &api.CreateVoteableRequest{
		Question: "Question",
		Answers:  []string{"a", "b", "c"},
	}
	resp, _ := s.CreateVoteable(ctx, req)
	if resp.UUID != "uuid" {
		t.Errorf("Got and Expected are not equals. Got: %v, expected: uuid", resp.UUID)
	}

	req = &api.CreateVoteableRequest{
		Question: "Error",
		Answers:  []string{},
	}
	_, err := s.CreateVoteable(ctx, req)
	if err == nil {
		t.Errorf("Got and Expected are not equals. Got: nil, expected: Error")
	}
}

func TestCastVote(t *testing.T) {
	req := &api.CastVoteRequest{
		UUID:        "Uuid",
		AnswerIndex: 0,
	}
	resp, _ := s.CastVote(ctx, req)
	if resp.Answer != "Answer" {
		t.Errorf("Got and Expected are not equals. Got: %v, expected: Answer", resp.Answer)
	}

	req = &api.CastVoteRequest{
		UUID:        "Error",
		AnswerIndex: 0,
	}
	_, err := s.CastVote(ctx, req)
	if err == nil {
		t.Errorf("Got and Expected are not equals. Got: nil, expected: Error")
	}
}
