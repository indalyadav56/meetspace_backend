package types

import (
	"time"

	"github.com/google/uuid"
)

type ChatContactResponse struct {
    RoomId uuid.UUID `json:"room_id"`
    IsGroup bool `json:"is_group"`
    RoomName string `json:"room_name,omitempty"`
    UserId *uuid.UUID `json:"user_id,omitempty"`
    Email string `json:"email,omitempty"`
    FirstName string `json:"first_name,omitempty"`
    LastName string `json:"last_name,omitempty"`
    IsActive bool `json:"is_active,omitempty"`
    UpdatedAt time.Time `json:"updated_at,omitempty"`
    LastMessage string `json:"last_message,omitempty"`
    MessageUnSeenCount int `json:"message_unseen_count,omitempty"`
    Status string `json:"status,omitempty"`
}

type CallRequestBody struct {
    RoomId string `json:"room_id"`
}
