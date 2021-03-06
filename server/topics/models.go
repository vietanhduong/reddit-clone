package topics

type (
	Topic struct {
		ID       int    `json:"id"`
		Content  string `json:"content"`
		UpVote   int    `json:"up_vote"`
		DownVote int    `json:"down_vote"`
	}

	Vote struct {
		ID int `json:"id"`
		UpVote int `json:"up_vote"`
		DownVote int `json:"down_vote"`
	}
)
