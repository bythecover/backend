package model

type Poll struct {
	Id        int      `json:"id"`
	CreatedBy string   `json:"created_by"`
	Options   []Option `json:"options"`
	Title     string   `json:"title"`
	Expired   bool     `json:"expired"`
}

type Option struct {
	Image string `json:"image"`
	Name  string `json:"name"`
	Id    int
}

type PollRepo interface {
	GetById(int, string) (Poll, error)
	CreatePoll(Poll) error
	GetCreatedBy(string) ([]Poll, error)
	ExpirePoll(int) error
}

type PollResult struct {
	Option
	Total int
}
