package domain

import (
	"fmt"
	"io"
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
	Content io.Reader
	Time    time.Time
}

// GetStorageName returns the name of the file for the storage system
func (u *Upload) GetStorageName() string {
	return fmt.Sprintf("%s_%s", u.Hash, u.Name)
}
