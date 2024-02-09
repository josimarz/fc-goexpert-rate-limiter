package ratelimiter

type Settings struct {
	ratelimit      int
	expirationTime int
	limitByToken   bool
}

func NewSettings(ratelimit, expirationTime int, limitByToken bool) *Settings {
	return &Settings{ratelimit, expirationTime, limitByToken}
}
