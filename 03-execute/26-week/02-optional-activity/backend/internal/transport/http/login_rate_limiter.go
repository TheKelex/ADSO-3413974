package http

import (
	"bytes"
	"encoding/json"
	"io"
	"net"
	"net/http"
	"sync"
	"time"
)

// loginRateLimit tracks the failed sign-in attempts of one identity.
type loginRateLimit struct {
	first  time.Time
	failed int
}

// LoginRateLimiter caps the consecutive failed sign-in attempts per IP and
// username. A window that has elapsed resets naturally; a successful sign-in
// resets immediately. The state lives in memory: a restart starts a clean
// window and no sensitive data is kept.
type LoginRateLimiter struct {
	mu        sync.Mutex
	maxFailed int
	window    time.Duration
	attempts  map[string]*loginRateLimit
	now       func() time.Time
}

// NewLoginRateLimiter builds a limiter that allows maxFailed failures inside
// window before blocking the identity with 429.
func NewLoginRateLimiter(maxFailed int, window time.Duration, now func() time.Time) *LoginRateLimiter {
	return &LoginRateLimiter{
		maxFailed: maxFailed,
		window:    window,
		attempts:  map[string]*loginRateLimit{},
		now:       now,
	}
}

// Wrap guards one sign-in handler: it blocks an identity that exhausted its
// attempts and counts the failures of the guarded handler by status code.
func (r *LoginRateLimiter) Wrap(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		username := usernameFrom(request)
		key := clientIP(request) + "|" + username
		if r.isBlocked(key) {
			time.Sleep(300 * time.Millisecond)
			tooManyAttempts(writer)
			return
		}
		recorder := &statusRecorder{ResponseWriter: writer}
		next.ServeHTTP(recorder, request)
		r.observe(key, recorder.code)
	})
}

// isBlocked reports whether the identity is inside a window with its failure
// budget exhausted.
func (r *LoginRateLimiter) isBlocked(key string) bool {
	r.mu.Lock()
	defer r.mu.Unlock()
	entry, ok := r.attempts[key]
	if !ok {
		return false
	}
	if r.now().Sub(entry.first) > r.window {
		delete(r.attempts, key)
		return false
	}
	return entry.failed >= r.maxFailed
}

// observe records a failure or resets the identity on success.
func (r *LoginRateLimiter) observe(key string, code int) {
	r.mu.Lock()
	defer r.mu.Unlock()
	switch code {
	case http.StatusOK:
		delete(r.attempts, key)
	default:
		if code != http.StatusUnauthorized {
			return
		}
		now := r.now()
		entry, ok := r.attempts[key]
		if !ok || now.Sub(entry.first) > r.window {
			r.attempts[key] = &loginRateLimit{first: now, failed: 1}
			return
		}
		entry.failed++
	}
}

// statusRecorder captures the status code the guarded handler wrote.
type statusRecorder struct {
	http.ResponseWriter
	code int
}

func (r *statusRecorder) WriteHeader(code int) {
	r.code = code
	r.ResponseWriter.WriteHeader(code)
}

func (r *statusRecorder) Write(payload []byte) (int, error) {
	if r.code == 0 {
		r.code = http.StatusOK
	}
	return r.ResponseWriter.Write(payload)
}

// usernameFrom reads the username the client sent without consuming the body
// of the request, so the guarded handler can read it again.
func usernameFrom(request *http.Request) string {
	body, err := io.ReadAll(request.Body)
	if err != nil {
		return ""
	}
	request.Body = io.NopCloser(bytes.NewReader(body))
	var payload struct {
		Username string `json:"username"`
	}
	if err := json.Unmarshal(body, &payload); err != nil {
		return ""
	}
	return payload.Username
}

// clientIP extracts the client address from the socket, never from a header a
// client controls.
func clientIP(request *http.Request) string {
	host, _, err := net.SplitHostPort(request.RemoteAddr)
	if err != nil {
		return request.RemoteAddr
	}
	return host
}
