package presenter

type Message struct {
	ID             string          `json:"id"`
	User           *User           `json:"user"`
	Content        string          `json:"content"`
	ReplyToMessage *ReplyToMessage `json:"reply_to_message"`
	CreatedAt      string          `json:"created_at"`
	UpdatedAt      string          `json:"updated_at"`
}

type ReplyToMessage struct {
	ID      string `json:"id"`
	User    *User  `json:"user"`
	Content string `json:"content"`
}
