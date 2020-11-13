package models

type Group struct {
	Channel
	Users []string `json:"usernames"`
}

type CreateGroupRequest struct {
	Name     string   `json:"name"`
	Members  []string `json:"members"`
	ReadOnly bool     `json:"readOnly"`
}

type InviteGroupRequest struct {
	RoomID string `json:"roomId"`
	UserID string `json:"userId"`
}
