package model

import (
	"time"

	"github.com/google/uuid"
)

type Card struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Front     string    `json:"front"`
	Back      string    `json:"back"`
	AudioURL  *string   `json:"audio_url,omitempty"`
	ImageURL  *string   `json:"image_url,omitempty"`
	DeckID    uuid.UUID `json:"deck_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
