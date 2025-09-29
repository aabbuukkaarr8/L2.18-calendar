package token

import (
	"github.com/google/uuid"
	"time"

	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwt"
)

var jwtSecret = []byte("your_secret_key")

func GenerateJWT(userID uuid.UUID, role string, email string) (string, error) {
	tok := jwt.New()
	_ = tok.Set("userID", userID)
	_ = tok.Set("role", role)
	_ = tok.Set("email", email)
	_ = tok.Set("exp", time.Now().Add(24*time.Hour).Unix())
	signed, err := jwt.Sign(tok, jwt.WithKey(jwa.HS256, jwtSecret))
	if err != nil {
		return "", err

	}
	return string(signed), nil
}
