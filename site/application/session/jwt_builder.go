package session

import (
	"fmt"
	"time"

	"github.com/cristalhq/jwt/v5"
)

type JWTBuilder struct {
	secret string
}

type JWTBuild struct {
	AccessToken  string
	RefreshToken string
}

func (b *JWTBuilder) SetSecret(secret string) *JWTBuilder {
	b.secret = secret
	return b
}

// Build creates an access token and a refresh token
func (b *JWTBuilder) Build(session Session) (JWTBuild, error) {
	var build JWTBuild
	if b.secret == "" {
		return build, fmt.Errorf("jwt secret not set")
	}

	signer, err := jwt.NewSignerHS(jwt.HS256, []byte(b.secret))
	if err != nil {
		return build, fmt.Errorf("error creating jwt signer: %w", err)
	}

	refreshTokenClaims := &jwt.RegisteredClaims{
		ID:        session.ID.String(),
		ExpiresAt: jwt.NewNumericDate(session.ExpriationTime),
	}
	accessTokenClaims := &jwt.RegisteredClaims{
		ID:        session.ID.String(),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 15)),
	}

	builder := jwt.NewBuilder(signer)
	refreshToken, err := builder.Build(refreshTokenClaims)
	if err != nil {
		return build, fmt.Errorf("error signing refresh token: %w", err)
	}
	accessToken, err := builder.Build(accessTokenClaims)
	if err != nil {
		return build, fmt.Errorf("error signing access token: %w", err)
	}

	build = JWTBuild{
		AccessToken:  accessToken.String(),
		RefreshToken: refreshToken.String(),
	}
	return build, nil
}
