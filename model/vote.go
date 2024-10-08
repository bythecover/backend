package model

type Vote struct {
	Selection   int    `json:"selection"`
	PollEventId int    `json:"poll_event_id"`
	Source      string `json:"source"`
	UserId      string `json:"user_id"`
}

type VoteRepo interface {
	SubmitVote(Vote) error
	HasUserVoted(string, int) bool
	GetResults(int) []PollResult
}
