package services

import (
	"context"
	"testing"

	"github.com/jmsilvadev/golangtechtask/api"
	"github.com/jmsilvadev/golangtechtask/pkg/providers"
)

func TestListAllVoteables(t *testing.T) {
	s := &Repo{}
	req := &api.ListVoteableRequest{
		PageSize: 1,
		Page:     "Page",
	}
	db := providers.NewMockDB(context.Background())
	resp, _ := s.ListAllVoteables(context.Background(), req, db)
	if resp.Page != "Page" {
		t.Errorf("Got and Expected are not equals. Got: %v, expected: Page", resp.Page)
	}

	req = &api.ListVoteableRequest{
		PageSize: 1,
		Page:     "Error",
	}
	_, err := s.ListAllVoteables(context.Background(), req, db)
	if err == nil {
		t.Errorf("Got and Expected are not equals. Got: nil, expected: Error")
	}

	req = &api.ListVoteableRequest{
		PageSize: 1,
		Page:     "Nil",
	}
	resp, err = s.ListAllVoteables(context.Background(), req, db)
	if err != nil {
		t.Errorf("Got and Expected are not equals. Got: nil, expected: Error")
	}
}

func TestCreateNewVoteable(t *testing.T) {
	s := &Repo{}
	req := &api.CreateVoteableRequest{
		Question: "Question",
		Answers:  []string{"a", "b", "c"},
	}
	db := providers.NewMockDB(context.Background())
	resp, _ := s.CreateNewVoteable(context.Background(), req, db)
	if resp.UUID != "uuid" {
		t.Errorf("Got and Expected are not equals. Got: %v, expected: uuid", resp.UUID)
	}

	req = &api.CreateVoteableRequest{
		Question: "Error",
		Answers:  []string{},
	}
	_, err := s.CreateNewVoteable(context.Background(), req, db)
	if err == nil {
		t.Errorf("Got and Expected are not equals. Got: nil, expected: Error")
	}
}

func TestUpdateCastVote(t *testing.T) {
	s := &Repo{}
	req := &api.CastVoteRequest{
		UUID:        "Uuid",
		AnswerIndex: 0,
	}
	db := providers.NewMockDB(context.Background())
	resp, _ := s.UpdateCastVote(context.Background(), req, db)
	if resp.Answer != "Answer" {
		t.Errorf("Got and Expected are not equals. Got: %v, expected: Answer", resp.Answer)
	}

	req = &api.CastVoteRequest{
		UUID:        "Error",
		AnswerIndex: 0,
	}
	_, err := s.UpdateCastVote(context.Background(), req, db)
	if err == nil {
		t.Errorf("Got and Expected are not equals. Got: nil, expected: Error")
	}
}
