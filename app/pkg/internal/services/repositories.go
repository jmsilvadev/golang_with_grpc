package services

import (
	"context"

	"github.com/jmsilvadev/golangtechtask/api"
	"github.com/jmsilvadev/golangtechtask/pkg/providers"
)

// Repo is a point receiver of repository pattern
type Repo struct{}

// ListAllVoteables gets all voteables in the database
func (r *Repo) ListAllVoteables(ctx context.Context, data *api.ListVoteableRequest, db providers.Storage) (*api.ListVoteableResponse, error) {
	resp, err := db.List(ctx, r.convertMessageToListVoteable(data))
	if err != nil {
		return nil, err
	}
	return r.convertListVoteableToMessage(resp), err
}

// CreateNewVoteable creates a New Votable in the database
func (r *Repo) CreateNewVoteable(ctx context.Context, data *api.CreateVoteableRequest, db providers.Storage) (*api.CreateVoteableResponse, error) {
	resp, err := db.Create(ctx, r.convertMessageToVoteable(data))
	if err != nil {
		return nil, err
	}
	return r.convertCreateVoteableToMessage(resp), err
}

// UpdateCastVote a votable in the database with one vote in the answer passed
func (r *Repo) UpdateCastVote(ctx context.Context, data *api.CastVoteRequest, db providers.Storage) (*api.CastVoteResponse, error) {
	resp, err := db.Update(ctx, r.convertMessageToCastVote(data))
	if err != nil {
		return nil, err
	}
	return r.convertCastVoteToMessage(resp), err
}
