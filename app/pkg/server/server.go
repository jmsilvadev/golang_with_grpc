package server

import (
	"context"

	"github.com/jmsilvadev/golangtechtask/api"
	"github.com/jmsilvadev/golangtechtask/pkg/internal/services"
	"github.com/jmsilvadev/golangtechtask/pkg/providers"
	"go.uber.org/zap"
)

// Server is the structure to handle the gRPC commands
type Server struct {
	api.UnimplementedVotingServiceServer
	services.Repo
	Db     providers.Storage
	Logger *zap.Logger
}

// CreateVoteable is the gRPC command to create a voteable
func (s *Server) CreateVoteable(ctx context.Context, in *api.CreateVoteableRequest) (*api.CreateVoteableResponse, error) {
	res, err := s.CreateNewVoteable(ctx, in, s.Db)
	if err != nil {
		s.Logger.Error(err.Error())
	}
	return res, err
}

// ListVoteables is the gRPC command to list voteables with pagination or not
func (s *Server) ListVoteables(ctx context.Context, in *api.ListVoteableRequest) (*api.ListVoteableResponse, error) {
	res, err := s.ListAllVoteables(ctx, in, s.Db)
	if err != nil {
		s.Logger.Error(err.Error())
	}
	return res, err
}

// CastVote  is the gRPC command to add a vote in a voteable answer
func (s *Server) CastVote(ctx context.Context, in *api.CastVoteRequest) (*api.CastVoteResponse, error) {
	res, err := s.UpdateCastVote(ctx, in, s.Db)
	if err != nil {
		s.Logger.Error(err.Error())
	}
	return res, err
}
