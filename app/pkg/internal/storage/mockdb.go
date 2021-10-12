package storage

import (
	"context"
	"errors"

	"github.com/jmsilvadev/golangtechtask/pkg/internal/entities"
)

// MockDb is a mock structure used by depency injectio to run tests
type MockDb struct{}

// HasTables checks if there are tables in the database
func (db *MockDb) HasTables(ctx context.Context) (bool, error) {
	return true, nil
}

// CreateTables creates a new database structure
func (db *MockDb) CreateTables(ctx context.Context) error {
	return nil
}

// Create create a new votable in the database
func (db *MockDb) Create(ctx context.Context, data *entities.Voteable) (*entities.CreateVoteableResponse, error) {
	if data.Question == "Error" {
		return nil, errors.New("Error")
	}

	resp := &entities.CreateVoteableResponse{
		UUID: "uuid",
	}
	return resp, nil
}

// Update adds a new vote in a votable answer in the database
func (db *MockDb) Update(ctx context.Context, data *entities.CastVoteRequest) (*entities.CastVoteResponse, error) {
	if data.UUID == "Error" {
		return nil, errors.New("Error")
	}

	resp := &entities.CastVoteResponse{
		Answer:      "Answer",
		AnswerVotes: 0,
	}
	return resp, nil
}

// List lists votables from the database using or not pagination
func (db *MockDb) List(ctx context.Context, data *entities.ListVoteableRequest) (*entities.ListVoteableResponse, error) {
	if data.Page == "Error" {
		return nil, errors.New("Error")
	}

	if data.Page == "Nil" {
		return nil, nil
	}

	ent := entities.Voteable{
		UUID:     "1",
		Question: "q",
		Answers:  []string{"a", "b", "c"},
		Votes:    map[string]int64{},
	}
	votes := []entities.Voteable{}
	votes = append(votes, ent)
	resp := &entities.ListVoteableResponse{
		Page:         "Page",
		NextPage:     "NextPage",
		PreviousPage: "PreviousPage",
		Votables:     votes,
	}
	return resp, nil
}
