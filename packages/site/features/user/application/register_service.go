package application

import (
	"context"

	"github.com/4strodev/4stroblog/site/features/user/domain"
	"github.com/4strodev/4stroblog/site/shared/domain/domainerrors"
	"github.com/google/uuid"
)

type RegisterService struct {
	UserService *domain.UserService
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
	user, err := domain.NewUser(req.Name, req.Email, req.Password)
	if err != nil {
		return res, err
	}
	select {
	case <-ctx.Done():
		return res, domainerrors.WrapError(domainerrors.CONTEXT_FINISHED, ctx.Err())
	default:
		// Proceed
	}

	// Check if context is done before persisting user
	err = s.UserService.CreateUser(ctx, user)
	res = UserRegisterResDTO{
		UserID: user.ID,
	}
	return
}
