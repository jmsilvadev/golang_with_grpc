package entities

import (
	"encoding/json"
	"testing"
)

func TestTransformVotable(t *testing.T) {
	configJSON := `
	{
		"uuid": "Uuid",
		"question": "Question",
		"answers": ["a", "b", "c"],
		"votes": {"0": 0, "1": 0, "2": 0}
	}
	`
	data := &Voteable{}
	err := json.Unmarshal([]byte(configJSON), data)
	if err != nil {
		t.Errorf("Got and Expected are not equals.\n got: %v, expected: %v ",
			err, nil)
	}
}

func TestTransformCreateVoteableRequest(t *testing.T) {
	configJSON := `
	{
		"question": "Question",
		"answers": ["a", "b", "c"]
	}
	`
	data := &CreateVoteableRequest{}
	err := json.Unmarshal([]byte(configJSON), data)
	if err != nil {
		t.Errorf("Got and Expected are not equals.\n got: %v, expected: %v ",
			err, nil)
	}
}

func TestTransformCreateVoteableResponse(t *testing.T) {
	configJSON := `
	{
		"uuid": "uuid"
	}
	`
	data := &CreateVoteableResponse{}
	err := json.Unmarshal([]byte(configJSON), data)
	if err != nil {
		t.Errorf("Got and Expected are not equals.\n got: %v, expected: %v ",
			err, nil)
	}
}

func TestTransformListVoteableRequest(t *testing.T) {
	configJSON := `
	{
		"page_size": 1,
		"page": "page"
	}
	`
	data := &CreateVoteableRequest{}
	err := json.Unmarshal([]byte(configJSON), data)
	if err != nil {
		t.Errorf("Got and Expected are not equals.\n got: %v, expected: %v ",
			err, nil)
	}
}

func TestTransformListVoteableResponse(t *testing.T) {
	configJSON := `
	{
		"page": "Uuid",
		"next_page": "next_page",
		"previous_page": "previous_page",
		"votables": [
			{
				"uuid": "Uuid",
				"question": "Question",
				"answers": ["a", "b", "c"],
				"votes": {"0": 0, "1": 0, "2": 0}
			}
		]
	}
	`
	data := &ListVoteableResponse{}
	err := json.Unmarshal([]byte(configJSON), data)
	if err != nil {
		t.Errorf("Got and Expected are not equals.\n got: %v, expected: %v ",
			err, nil)
	}
}

func TestTransformCastVoteRequest(t *testing.T) {
	configJSON := `
	{
		"uuid": "uuid",
		"answer_index": 0
	}
	`
	data := &CastVoteRequest{}
	err := json.Unmarshal([]byte(configJSON), data)
	if err != nil {
		t.Errorf("Got and Expected are not equals.\n got: %v, expected: %v ",
			err, nil)
	}
}

func TestTransformCastVoteResponse(t *testing.T) {
	configJSON := `
	{
		"answer": "uuid",
		"answer_votes": 10
	}
	`
	data := &CastVoteResponse{}
	err := json.Unmarshal([]byte(configJSON), data)
	if err != nil {
		t.Errorf("Got and Expected are not equals.\n got: %v, expected: %v ",
			err, nil)
	}
}
