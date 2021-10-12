package entities

// Voteable structure
type Voteable struct {
	UUID     string
	Question string
	Answers  []string
	Votes    map[string]int64
}

// CreateVoteableRequest structure
type CreateVoteableRequest struct {
	Question string
	Answers  []string
}

// CreateVoteableResponse structure
type CreateVoteableResponse struct {
	UUID string
}

// ListVoteableRequest structure
type ListVoteableRequest struct {
	PageSize int64
	Page     string
}

// ListVoteableResponse structure
type ListVoteableResponse struct {
	Page         string
	NextPage     string
	PreviousPage string
	Votables     []Voteable
}

// CastVoteRequest structure
type CastVoteRequest struct {
	UUID        string
	AnswerIndex int64
}

// CastVoteResponse structure
type CastVoteResponse struct {
	Answer      string
	AnswerVotes int64
}
