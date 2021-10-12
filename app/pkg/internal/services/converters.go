package services

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/jmsilvadev/golangtechtask/api"
	"github.com/jmsilvadev/golangtechtask/pkg/internal/entities"
)

func (r *Repo) convertMessageToListVoteable(data *api.ListVoteableRequest) *entities.ListVoteableRequest {
	return &entities.ListVoteableRequest{
		PageSize: data.PageSize,
		Page:     data.Page,
	}
}

func (r *Repo) convertMessageToVoteable(data *api.CreateVoteableRequest) *entities.Voteable {
	votes := map[string]int64{}
	for k := range data.Answers {
		votes[fmt.Sprint(k)] = 0
	}
	return &entities.Voteable{
		UUID:     uuid.New().String(),
		Question: data.Question,
		Answers:  data.Answers,
		Votes:    votes,
	}
}

func (r *Repo) convertMessageToCastVote(data *api.CastVoteRequest) *entities.CastVoteRequest {
	return &entities.CastVoteRequest{
		UUID:        data.UUID,
		AnswerIndex: data.AnswerIndex,
	}
}

func (r *Repo) convertVoteableToMessage(data entities.Voteable) *api.Voteable {
	return &api.Voteable{
		UUID:     data.UUID,
		Question: data.Question,
		Answers:  data.Answers,
		Votes:    data.Votes,
	}
}

func (r *Repo) convertListVoteableToMessage(data *entities.ListVoteableResponse) *api.ListVoteableResponse {
	result := []*api.Voteable{}
	if data == nil {
		return &api.ListVoteableResponse{}
	}
	for _, vote := range data.Votables {
		result = append(result, r.convertVoteableToMessage(vote))
	}
	return &api.ListVoteableResponse{
		Page:         data.Page,
		NextPage:     data.NextPage,
		PreviousPage: data.PreviousPage,
		Votables:     result,
	}
}

func (r *Repo) convertCreateVoteableToMessage(data *entities.CreateVoteableResponse) *api.CreateVoteableResponse {
	return &api.CreateVoteableResponse{
		UUID: data.UUID,
	}
}

func (r *Repo) convertCastVoteToMessage(data *entities.CastVoteResponse) *api.CastVoteResponse {
	return &api.CastVoteResponse{
		Answer:      data.Answer,
		AnswerVotes: data.AnswerVotes,
	}
}
