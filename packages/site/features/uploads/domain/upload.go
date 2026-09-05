package domain

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type Upload struct {
	// Id used for internal logic only, not used to identify
	// object in object storage
	ID uuid.UUID
	// Hash is used to prevent duplication. Used to name and identify
	// object in object storage.
	Hash []byte
	// string encoding for [Hash]. Do not update this value manually
	hashString string
	// size of the blob do not update
	Size uint64
	// The mime type of the uploaded file
	MimeType string
	// A human readable name
	Name string
	// Upload time
	Time time.Time
}

// DigestBlob takes a reader and calculates its intrinsic meta-data like hash and size
func (u *Upload) DigestBlob(blob io.Reader) error {
	hash := sha256.New()
	bs := hash.BlockSize()

	const MIME_BS = 512
	var mimeBlock = make([]byte, MIME_BS)
	var n_mimeBlock = 0
	var size uint64 = 0
	buf := make([]byte, bs)
	for {
		n, err := blob.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		if n_mimeBlock < MIME_BS {
			n_mimeBlock = copy(mimeBlock[n_mimeBlock:], buf)
		}

		_, err = hash.Write(buf[:n])
		if err != nil {
			return err
		}
		size += uint64(n)
	}

	u.MimeType = http.DetectContentType(mimeBlock)

	u.Hash = hash.Sum(nil)
	u.Size = size
	return nil
}

// StringHash returns hash represented in hex
// if hash is not set returns an empty string
func (u Upload) StringHash() string {
	if u.Hash != nil && u.hashString == "" {
		u.hashString = hex.EncodeToString(u.Hash)
	}

	return u.hashString
}
