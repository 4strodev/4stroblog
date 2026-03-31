package application

import (
	"context"

	"github.com/4strodev/4stroblog/site/features/session/application/dto"
	"github.com/google/uuid"
)

func (s *SessionAppService) FindById(ctx context.Context, id uuid.UUID) (dto.SessionDto, error) {
	var sessionDto dto.SessionDto

	session, err := s.SessionService.FindById(ctx, id)
	if err != nil {
		return sessionDto, err
	}

	sessionDto = dto.SessionDto{
		ID:             session.ID,
		UserID:         session.UserID,
		ExpirationTime: session.ExpriationTime,
		ProfileID:      session.ProfileID,
	}

	return sessionDto, nil
}
