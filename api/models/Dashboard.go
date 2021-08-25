package models



type Dashboard   struct {
	UserID    uint32         `json:"user_id"`
	Widgets   []Widget       `json:"widgets"`
}