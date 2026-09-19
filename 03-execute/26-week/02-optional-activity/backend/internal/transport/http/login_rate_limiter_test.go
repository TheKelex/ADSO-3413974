package http

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"workshop/internal/domain"
)

func fixedClockTime() func() time.Time {
	moment := time.Date(2026, time.March, 1, 9, 0, 0, 0, time.UTC)
	return func() time.Time { return moment }
}

func loginRequest(username string) *http.Request {
	body := `{"username":"` + username + `","password":"wrong"}`
	request := httptest.NewRequest(http.MethodPost, "/api/session", strings.NewReader(body))
	request.RemoteAddr = "192.168.1.10:40000"
	return request
}

// advancingClock lets a test move the limiter clock past its window.
type advancingClock struct {
	moment time.Time
}

func (c *advancingClock) current() time.Time { return c.moment }

func (c *advancingClock) advance(duration time.Duration) { c.moment = c.moment.Add(duration) }

func TestLoginRateLimiterBlocksAfterTheConfiguredFailures(t *testing.T) {
	limiter := NewLoginRateLimiter(5, 15*time.Minute, fixedClockTime())
	handler := limiter.Wrap(http.HandlerFunc(func(writer http.ResponseWriter, _ *http.Request) {
		failure(writer, domain.ErrUnauthorized)
	}))

	for attempt := 1; attempt <= 5; attempt++ {
		recorder := httptest.NewRecorder()
		handler.ServeHTTP(recorder, loginRequest("admin"))
		if recorder.Code != http.StatusUnauthorized {
			t.Fatalf("attempt %d must be a 401, got %d", attempt, recorder.Code)
		}
	}

	recorder := httptest.NewRecorder()
	handler.ServeHTTP(recorder, loginRequest("admin"))
	if recorder.Code != http.StatusTooManyRequests {
		t.Fatalf("the sixth failure must be blocked with 429, got %d", recorder.Code)
	}
	var payload errorPayload
	if err := json.Unmarshal(recorder.Body.Bytes(), &payload); err != nil {
		t.Fatalf("the block must answer JSON: %v", err)
	}
	if payload.Code != "too_many_attempts" {
		t.Fatalf("the block must carry the too_many_attempts code, got %q", payload.Code)
	}
}

func TestLoginRateLimiterResetsTheIdentityOnASuccess(t *testing.T) {
	attempts := 0
	limiter := NewLoginRateLimiter(5, 15*time.Minute, fixedClockTime())
	handler := limiter.Wrap(http.HandlerFunc(func(writer http.ResponseWriter, _ *http.Request) {
		attempts++
		if attempts == 3 {
			respond(writer, http.StatusOK, map[string]string{"token": "succeeded"})
			return
		}
		failure(writer, domain.ErrUnauthorized)
	}))

	for attempt := 1; attempt <= 2; attempt++ {
		recorder := httptest.NewRecorder()
		handler.ServeHTTP(recorder, loginRequest("admin"))
		if recorder.Code != http.StatusUnauthorized {
			t.Fatalf("attempt %d must be a 401, got %d", attempt, recorder.Code)
		}
	}

	recorder := httptest.NewRecorder()
	handler.ServeHTTP(recorder, loginRequest("admin"))
	if recorder.Code != http.StatusOK {
		t.Fatalf("a correct credential must succeed, got %d", recorder.Code)
	}

	for attempt := 1; attempt <= 2; attempt++ {
		recorder = httptest.NewRecorder()
		handler.ServeHTTP(recorder, loginRequest("admin"))
		if recorder.Code != http.StatusUnauthorized {
			t.Fatalf("after a success the budget must reset and allow a wrong 401, got %d", recorder.Code)
		}
	}
}

func TestLoginRateLimiterKeepsOneIdentityPerIPAndUsername(t *testing.T) {
	limiter := NewLoginRateLimiter(1, 15*time.Minute, fixedClockTime())
	handler := limiter.Wrap(http.HandlerFunc(func(writer http.ResponseWriter, _ *http.Request) {
		failure(writer, domain.ErrUnauthorized)
	}))

	recorder := httptest.NewRecorder()
	handler.ServeHTTP(recorder, loginRequest("admin"))
	if recorder.Code != http.StatusUnauthorized {
		t.Fatalf("the first failure must be a 401, got %d", recorder.Code)
	}

	recorder = httptest.NewRecorder()
	handler.ServeHTTP(recorder, loginRequest("admin"))
	if recorder.Code != http.StatusTooManyRequests {
		t.Fatalf("a second failure of the blocked identity must be 429, got %d", recorder.Code)
	}

	recorder = httptest.NewRecorder()
	handler.ServeHTTP(recorder, loginRequest("jperez"))
	if recorder.Code != http.StatusUnauthorized {
		t.Fatalf("a different identity must keep its own budget, got %d", recorder.Code)
	}
}

func TestLoginRateLimiterResetsInsideAnElapsedWindow(t *testing.T) {
	clock := &advancingClock{moment: time.Date(2026, time.March, 1, 9, 0, 0, 0, time.UTC)}
	limiter := NewLoginRateLimiter(1, 15*time.Minute, clock.current)
	handler := limiter.Wrap(http.HandlerFunc(func(writer http.ResponseWriter, _ *http.Request) {
		failure(writer, domain.ErrUnauthorized)
	}))

	recorder := httptest.NewRecorder()
	handler.ServeHTTP(recorder, loginRequest("admin"))
	if recorder.Code != http.StatusUnauthorized {
		t.Fatalf("the first failure must be a 401, got %d", recorder.Code)
	}

	recorder = httptest.NewRecorder()
	handler.ServeHTTP(recorder, loginRequest("admin"))
	if recorder.Code != http.StatusTooManyRequests {
		t.Fatalf("a second failure must be blocked inside the window, got %d", recorder.Code)
	}

	clock.advance(16 * time.Minute)

	recorder = httptest.NewRecorder()
	handler.ServeHTTP(recorder, loginRequest("admin"))
	if recorder.Code != http.StatusUnauthorized {
		t.Fatalf("after the window elapses the budget must reset, got %d", recorder.Code)
	}
}
