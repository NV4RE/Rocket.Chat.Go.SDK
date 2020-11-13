package models

import "time"

type Group struct {
	ID    string `json:"_id"`
	Name  string `json:"name"`
	Fname string `json:"fname,omitempty"`
	Type  string `json:"t"`
	Msgs  int    `json:"msgs"`

	ReadOnly  bool `json:"ro,omitempty"`
	SysMes    bool `json:"sysMes,omitempty"`
	Default   bool `json:"default"`
	Broadcast bool `json:"broadcast,omitempty"`

	Timestamp *time.Time `json:"ts,omitempty"`
	UpdatedAt *time.Time `json:"_updatedAt,omitempty"`

	User        *User    `json:"u,omitempty"`
	LastMessage *Message `json:"lastMessage,omitempty"`

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
