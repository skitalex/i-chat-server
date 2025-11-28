package chatdomain

type Chat struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type CreateChatRequest struct {
	Name string `json:"name"`
	Type int    `json:"type"`
}
