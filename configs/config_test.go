package configs

import "testing"

func TestLoadConfig(t *testing.T) {
	path := "."
	config, err := LoadConfig(path)
	if err != nil {
		t.Errorf("err = %v, want %v", err, nil)
	}
	if config.Host != "localhost" {
		t.Errorf("config.Host = %v, want %v", config.Host, "localhost")
	}
	if config.Port != "6379" {
		t.Errorf("config.Port = %v, want %v", config.Port, "6379")
	}
	if config.Password != "lI6sI5dS8nZ0lG6p" {
		t.Errorf("config.Password = %v, want %v", config.Password, "lI6sI5dS8nZ0lG6p")
	}
	if config.DB != 0 {
		t.Errorf("config.DB = %v, want %v", config.DB, 0)
	}
	if config.RateLimit != 10 {
		t.Errorf("config.RateLimit = %v, want %v", config.RateLimit, 10)
	}
	if !config.LimitByToken {
		t.Errorf("config.LimitByToken = %v, want %v", config.LimitByToken, true)
	}
	if config.ExpirationTime != 60 {
		t.Errorf("config.ExpirationTime = %v, want %v", config.ExpirationTime, 60)
	}
}
