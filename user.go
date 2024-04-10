package main

type User struct {
	ID            int64  `json:"id,omitempty"`
	Username      string `json:"username"`
	IsProvider    bool   `json:"isProvider"`
	Email         string `json:"email"`
	Password      string `json:"password,omitempty"` // Hashed password, omit from JSON
	ProfilePicURL string `json:"profilePicURL"`
	City          string `json:"city"`
	PCSpecs       string `json:"pcSpecs"`
	Description   string `json:"description"`
	CloudService  string `json:"cloudService"`
	CreatedAt     string `json:"createdAt,omitempty"`
}

type Message struct {
	MessageID  int64  `json:"messageID,omitempty"`
	SenderID   int64  `json:"senderID"`
	ReceiverID int64  `json:"receiverID"`
	Message    string `json:"message"`
	SentAt     string `json:"sentAt,omitempty"`
}
