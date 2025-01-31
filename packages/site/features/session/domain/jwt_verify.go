package domain

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/cristalhq/jwt/v5"
)

type JwtVerify struct {
	secret   string
	verifier jwt.Verifier
}

func NewJwtVerify(secret string) JwtVerify {
	return JwtVerify{
		secret: secret,
	}
}

const SIGNER_ALGORITHM = jwt.HS256

// Verify checks if token is valid. If is valid a nil error is returned.
// In other case an error is returned
func (v JwtVerify) Verify(rawToken string) error {
	var err error
	err = v.ensureVerifierIsSet()
	if err != nil {
		return err
	}

	token, err := jwt.Parse([]byte(rawToken), v.verifier)
	if err != nil {
		return err
	}

	err = v.verifier.Verify(token)
	if err != nil {
		return err
	}

	// Check claims
	var claims jwt.RegisteredClaims
	err = json.Unmarshal(token.Claims(), &claims)
	if err != nil {
		return err
	}

	if !claims.IsValidAt(time.Now()) {
		return errors.New("jwt token is expired")
	}
	return nil
}

// ensureVerifierIsSet sets a [jwt.Verfier] if it is not set
// if it is already set does nothing
func (v *JwtVerify) ensureVerifierIsSet() error {
	if v.verifier != nil {
		return nil
	}

	verifier, err := jwt.NewVerifierHS(SIGNER_ALGORITHM, []byte(v.secret))
	if err != nil {
		return err
	}

	v.verifier = verifier

	return nil
}
