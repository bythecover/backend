package model

type Vote struct {
	Selection   int    `json:"selection"`
	PollEventId int    `json:"poll_event_id"`
	Source      string `json:"source"`
	UserId      int    `json:"user_id"`
}
