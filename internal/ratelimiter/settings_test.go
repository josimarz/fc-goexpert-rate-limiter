package ratelimiter

import "testing"

func TestNewSettings(t *testing.T) {
	rateLimit := 10
	expirationTime := 60
	limitByToken := true
	settings := NewSettings(rateLimit, expirationTime, limitByToken)
	if settings.ratelimit != rateLimit {
		t.Errorf("settings.rateLimit = %v, want %v", settings.ratelimit, rateLimit)
	}
	if settings.expirationTime != expirationTime {
		t.Errorf("settings.expirationTime = %v, want %v", settings.expirationTime, expirationTime)
	}
	if settings.limitByToken != limitByToken {
		t.Errorf("settings.limitByToken = %v, want %v", settings.limitByToken, limitByToken)
	}
}
