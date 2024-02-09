package gateway

import (
	"context"
)

type Token struct {
	token string
	limit int
}

type TokenInMemoryGateway struct {
	tokens []*Token
}

func NewTokenInMemoryGateway() *TokenInMemoryGateway {
	return &TokenInMemoryGateway{tokens: []*Token{
		{
			token: "p7eWgd0PvJcqB3ea45pw3k5thpWaqpI12RGYU3MiP91Kgao5MCXtlFtL2rwISxYL",
			limit: 100,
		},
		{
			token: "65aYJmkHf8QC52s10HYjVV5xtfwaKf3qC2J79cKaFZarStmbT6Mueic195OXLXVy",
			limit: 200,
		},
		{
			token: "DUc0K3ojDA0kQgCVl2SEfT5evimjAEojGs5QOxMfA3JAgdrF6I5l8hHFXDDtpxqv",
			limit: 300,
		},
		{
			token: "sEUMDAjUhxGzTcXwerMiCwdSfe0vp24q9ISuZrYra0Hj65gtGpGeA5zZt4bjG0gz",
			limit: 400,
		},
		{
			token: "IB9BB71TsBk8QevTskkFusRrZwfUBjmqACA7vvsc1TCpS5FMAmZZTZx1R1OhA12y",
			limit: 500,
		},
	}}
}

func (g *TokenInMemoryGateway) GetLimit(ctx context.Context, token string) (int, error) {
	for _, t := range g.tokens {
		if t.token == token {
			return t.limit, nil
		}
	}
	return 0, nil
}
