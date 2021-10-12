package storage

import (
	"context"
	"testing"

	"github.com/jmsilvadev/golangtechtask/pkg/internal/entities"
)

func TestListAllVoteables(t *testing.T) {
	req := &entities.ListVoteableRequest{
		PageSize: 1,
		Page:     "Page",
	}

	db := &MockDb{}
	resp, _ := db.List(context.Background(), req)
	if resp.Page != "Page" {
		t.Errorf("Got and Expected are not equals. Got: %v, expected: Page", resp.Page)
	}

	req = &entities.ListVoteableRequest{
		PageSize: 1,
		Page:     "Error",
	}
	_, err := db.List(context.Background(), req)
	if err == nil {
		t.Errorf("Got and Expected are not equals. Got: nil, expected: Error")
	}

	req = &entities.ListVoteableRequest{
		PageSize: 1,
		Page:     "Nil",
	}
	resp, err = db.List(context.Background(), req)
	if err != nil {
		t.Errorf("Got and Expected are not equals. Got: nil, expected: Error")
	}
}

func TestCreateNewVoteable(t *testing.T) {
	db := &MockDb{}
	req := &entities.Voteable{
		UUID:     "",
		Question: "Question",
		Answers:  []string{"a", "b", "c"},
		Votes:    map[string]int64{},
	}
	resp, _ := db.Create(context.Background(), req)
	if resp.UUID != "uuid" {
		t.Errorf("Got and Expected are not equals. Got: %v, expected: uuid", resp.UUID)
	}

	req = &entities.Voteable{
		UUID:     "",
		Question: "Error",
		Answers:  []string{},
		Votes:    map[string]int64{},
	}
	_, err := db.Create(context.Background(), req)
	if err == nil {
		t.Errorf("Got and Expected are not equals. Got: nil, expected: Error")
	}
}

func TestUpdateCastVote(t *testing.T) {
	db := &MockDb{}
	req := &entities.CastVoteRequest{
		UUID:        "Uuid",
		AnswerIndex: 0,
	}
	resp, _ := db.Update(context.Background(), req)
	if resp.Answer != "Answer" {
		t.Errorf("Got and Expected are not equals. Got: %v, expected: Answer", resp.Answer)
	}

	req = &entities.CastVoteRequest{
		UUID:        "Error",
		AnswerIndex: 0,
	}
	_, err := db.Update(context.Background(), req)
	if err == nil {
		t.Errorf("Got and Expected are not equals. Got: nil, expected: Error")
	}
}

func TestHasTable(t *testing.T) {
	db := &MockDb{}
	_, err := db.HasTables(context.Background())
	if err != nil {
		t.Errorf("Got and Expected are not equals. Got: !nil, expected: nil")
	}
}

func TestCreateTables(t *testing.T) {
	db := &MockDb{}
	err := db.CreateTables(context.Background())
	if err != nil {
		t.Errorf("Got and Expected are not equals. Got: !nil, expected: nil")
	}
}
