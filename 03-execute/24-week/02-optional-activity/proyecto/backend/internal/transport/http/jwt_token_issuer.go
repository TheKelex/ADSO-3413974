package http

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"time"

	"workshop/internal/domain"
)

// tokenClaim is the payload carried by a session token. It holds no personal
// data: only the account identifier, its role and the expiry.
type tokenClaim struct {
	UserID    string `json:"userId"`
	Role      string `json:"role"`
	ExpiresAt int64  `json:"expiresAt"`
}

// TokenIssuer signs and verifies session tokens with HMAC-SHA256 over a
// compact JSON claim. The signature is compared in constant time, so a wrong
// token cannot be discovered byte by byte through timing.
type TokenIssuer struct {
	secret []byte
	ttl    time.Duration
}

// NewTokenIssuer wires the token issuer with the configured signing key.
func NewTokenIssuer(secret string, ttl time.Duration) TokenIssuer {
	return TokenIssuer{secret: []byte(secret), ttl: ttl}
}

// Issue signs a token for the account and returns it with its expiry.
func (t TokenIssuer) Issue(userID string, role domain.Role, issuedAt time.Time) (string, time.Time, error) {
	expiresAt := issuedAt.Add(t.ttl)
	payload, err := json.Marshal(tokenClaim{
		UserID:    userID,
		Role:      string(role),
		ExpiresAt: expiresAt.Unix(),
	})
	if err != nil {
		return "", time.Time{}, fmt.Errorf("signing the session token: %w", err)
	}
	encoded := base64.RawURLEncoding.EncodeToString(payload)
	return encoded + "." + t.sign(encoded), expiresAt, nil
}

// Verify returns the claim of a valid token and rejects a token whose
// signature does not match or whose expiry has passed.
func (t TokenIssuer) Verify(token string, now time.Time) (tokenClaim, error) {
	separator := -1
	for index := len(token) - 1; index >= 0; index-- {
		if token[index] == '.' {
			separator = index
			break
		}
	}
	if separator <= 0 || separator == len(token)-1 {
		return tokenClaim{}, domain.ErrUnauthorized
	}
	encoded, signature := token[:separator], token[separator+1:]
	if !hmac.Equal([]byte(signature), []byte(t.sign(encoded))) {
		return tokenClaim{}, domain.ErrUnauthorized
	}
	payload, err := base64.RawURLEncoding.DecodeString(encoded)
	if err != nil {
		return tokenClaim{}, domain.ErrUnauthorized
	}
	var claim tokenClaim
	if err := json.Unmarshal(payload, &claim); err != nil {
		return tokenClaim{}, domain.ErrUnauthorized
	}
	if now.Unix() >= claim.ExpiresAt {
		return tokenClaim{}, domain.ErrUnauthorized
	}
	return claim, nil
}

func (t TokenIssuer) sign(encoded string) string {
	mac := hmac.New(sha256.New, t.secret)
	mac.Write([]byte(encoded))
	return base64.RawURLEncoding.EncodeToString(mac.Sum(nil))
}
