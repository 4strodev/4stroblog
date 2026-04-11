package application

import (
	"context"

	"github.com/4strodev/4stroblog/site/features/user/domain"
	"github.com/4strodev/4stroblog/site/shared/domain/domainerrors"
	"github.com/google/uuid"
)

type RegisterService struct {
	UserRepository domain.UserRepository
	ProfileService *domain.ProfileService
}

type RegisterReqDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type UserRegisterResDTO struct {
	UserID uuid.UUID `json:"userId"`
}

func (s *RegisterService) Register(ctx context.Context, req RegisterReqDTO) (res UserRegisterResDTO, err error) {
	select {
	case <-ctx.Done():
		return res, domainerrors.WrapError(domainerrors.CONTEXT_FINISHED, ctx.Err())
	default:
	}

	user, err := domain.NewUser(req.Password)
	if err != nil {
		return res, err
	}

	if err = s.UserRepository.Save(ctx, user); err != nil {
		return res, err
	}

	_, err = s.ProfileService.CreateProfile(ctx, user.ID, req.Name, req.Email)
	if err != nil {
		return res, err
	}

	res = UserRegisterResDTO{UserID: user.ID}
	return res, nil
}
