package application

import (
	"context"

	"github.com/google/uuid"
)

type SessionDeleteReq struct {
	ID uuid.UUID `json:"id"`
}

func (s *SessionAppService) Delete(ctx context.Context, req SessionDeleteReq) error {
	return s.SessionService.Repository.Delete(ctx, req.ID)
}
