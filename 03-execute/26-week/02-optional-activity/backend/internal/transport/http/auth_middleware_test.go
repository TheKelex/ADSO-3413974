package http

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"workshop/internal/domain"
)

// testMoment keeps every API test independent from the wall clock.
var testMoment = time.Date(2026, time.March, 1, 9, 0, 0, 0, time.UTC)

func testClock() func() time.Time {
	return func() time.Time { return testMoment }
}

func testIssuer() TokenIssuer {
	return NewTokenIssuer("a-test-signing-key-that-is-long-enough", time.Hour)
}

// tokenFor signs a token for a role so a test can call a protected route.
func tokenFor(t *testing.T, role domain.Role) string {
	t.Helper()
	token, _, err := testIssuer().Issue("user-1", role, testMoment)
	if err != nil {
		t.Fatalf("issuing the test token failed: %v", err)
	}
	return token
}

// protectedProbe is a handler that only reports who reached it.
func protectedProbe(reached *bool) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, _ *http.Request) {
		*reached = true
		writer.WriteHeader(http.StatusOK)
	})
}

func TestAuthMiddlewareRejectsARequestWithoutAToken(t *testing.T) {
	reached := false
	handler := authMiddleware(testIssuer(), testClock())(protectedProbe(&reached))

	recorder := httptest.NewRecorder()
	handler.ServeHTTP(recorder, httptest.NewRequest(http.MethodGet, "/api/customer", nil))

	if recorder.Code != http.StatusUnauthorized {
		t.Fatalf("a request without a token must be rejected with 401, got %d", recorder.Code)
	}
	if reached {
		t.Fatal("the protected handler must never run for an unauthenticated request")
	}
}

func TestAuthMiddlewareAcceptsAValidToken(t *testing.T) {
	reached := false
	handler := authMiddleware(testIssuer(), testClock())(protectedProbe(&reached))

	request := httptest.NewRequest(http.MethodGet, "/api/customer", nil)
	request.Header.Set("Authorization", "Bearer "+tokenFor(t, domain.RoleAdministrator))
	recorder := httptest.NewRecorder()
	handler.ServeHTTP(recorder, request)

	if recorder.Code != http.StatusOK || !reached {
		t.Fatalf("a valid token must reach the handler, got %d reached=%v", recorder.Code, reached)
	}
}

func TestAuthMiddlewareRejectsATamperedToken(t *testing.T) {
	reached := false
	handler := authMiddleware(testIssuer(), testClock())(protectedProbe(&reached))

	tampered := tokenFor(t, domain.RoleAdministrator) + "x"
	request := httptest.NewRequest(http.MethodGet, "/api/customer", nil)
	request.Header.Set("Authorization", "Bearer "+tampered)
	recorder := httptest.NewRecorder()
	handler.ServeHTTP(recorder, request)

	if recorder.Code != http.StatusUnauthorized || reached {
		t.Fatalf("a tampered signature must be rejected, got %d reached=%v", recorder.Code, reached)
	}
}

func TestAuthMiddlewareRejectsAnExpiredToken(t *testing.T) {
	reached := false
	expiredClock := func() time.Time { return testMoment.Add(2 * time.Hour) }
	handler := authMiddleware(testIssuer(), expiredClock)(protectedProbe(&reached))

	request := httptest.NewRequest(http.MethodGet, "/api/customer", nil)
	request.Header.Set("Authorization", "Bearer "+tokenFor(t, domain.RoleAdministrator))
	recorder := httptest.NewRecorder()
	handler.ServeHTTP(recorder, request)

	if recorder.Code != http.StatusUnauthorized || reached {
		t.Fatalf("an expired token must be rejected, got %d reached=%v", recorder.Code, reached)
	}
}

func TestCorsMiddlewareNeverAnswersWithAWildcard(t *testing.T) {
	reached := false
	handler := corsMiddleware("http://localhost:4173")(protectedProbe(&reached))

	request := httptest.NewRequest(http.MethodGet, "/api/health", nil)
	request.Header.Set("Origin", "http://evil.example")
	recorder := httptest.NewRecorder()
	handler.ServeHTTP(recorder, request)

	if recorder.Header().Get("Access-Control-Allow-Origin") != "" {
		t.Fatal("an unknown origin must receive no allow origin header")
	}

	allowed := httptest.NewRequest(http.MethodGet, "/api/health", nil)
	allowed.Header.Set("Origin", "http://localhost:4173")
	allowedRecorder := httptest.NewRecorder()
	handler.ServeHTTP(allowedRecorder, allowed)

	if allowedRecorder.Header().Get("Access-Control-Allow-Origin") != "http://localhost:4173" {
		t.Fatal("the configured origin must be echoed back exactly, never as a wildcard")
	}
}
