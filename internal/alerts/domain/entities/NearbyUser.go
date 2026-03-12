package entities

type UserEntry struct {
	UserID int `json:"user_id"`
}

type AlertPayload struct {
	SenderID     int         `json:"sender_id"`
	SenderName   string      `json:"sender_name"`
	Latitude     float64     `json:"latitude"`
	Longitude    float64     `json:"longitude"`
	UsersFamily  []UserEntry `json:"users_family"`
	UsersNetwork []UserEntry `json:"users_network"`
}
