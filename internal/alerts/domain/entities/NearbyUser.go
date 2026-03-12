package entities

type UserEntry struct {
	UserID int `json:"user_id"`
}

type AlertPayload struct {
	SenderID    int         `json:"sender_id"`
	SenderName  string      `json:"sender_name"`
	UsersFamily []UserEntry `json:"users_family"`
	UsersNetwork []UserEntry `json:"users_network"`
}