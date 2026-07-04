package domain

import (
	"time"

	"github.com/google/uuid"
)

type Upload struct {
	ID   uuid.UUID
	Hash string
	// The mime type of the uploaded file
	MimeType string
	// A human readable name
	Name    string
	// Upload time
	Time    time.Time
}
