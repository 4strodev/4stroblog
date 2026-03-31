package domain

import (
	"context"

	"github.com/4strodev/4stroblog/site/shared/domain/domainerrors"
	"github.com/google/uuid"
)

type SessionService struct {
	Repository SessionRepository
}

func (s *SessionService) Save(ctx context.Context, session Session) error {
	sessionFound, err := s.Repository.FindByProfileId(ctx, session.ProfileID)
	_, entityNotFound := domainerrors.Is(err, domainerrors.ENTITY_NOT_FOUND)
	if err != nil && !entityNotFound {
		return err
	}

	if !entityNotFound {
		err := s.Repository.Delete(ctx, sessionFound.ID)
		if err != nil {
			return err
		}
	}

	return s.Repository.Save(ctx, session)
}

func (s *SessionService) FindById(ctx context.Context, id uuid.UUID) (Session, error) {
	session, err := s.Repository.FindById(ctx, id)
	return session, err
}
