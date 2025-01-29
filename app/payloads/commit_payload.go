package payloads

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Date  string `json:"date"`
}

type Commit struct {
	Author    User   `json:"author"`
	Committer User   `json:"committer"`
	Message   string `json:"message"`
}

type CommitPayload struct {
	Sha    string `json:"sha"`
	NodeId string `json:"node_id"`
	Commit Commit `json:"commit"`
}
