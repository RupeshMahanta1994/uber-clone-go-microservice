package repository

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type ITokenRepsitory interface {
	BlackListToken(ctx context.Context, token string, expiry time.Duration) error
	IsBlackListed(ctx context.Context, token string) (bool, error)
}
type TokenRepository struct {
	client *redis.Client
}

func NewTokenRepository(client *redis.Client) *TokenRepository {
	return &TokenRepository{
		client: client,
	}
}

func (r *TokenRepository) BlackListToken(ctx context.Context, token string, expiry time.Duration) error {
	return r.client.Set(ctx, token, "blacklisted", expiry).Err()
}
func (r *TokenRepository) IsBlackListed(ctx context.Context, token string) (bool, error) {
	val, err := r.client.Get(ctx, token).Result()
	if err == redis.Nil {
		return false, nil //Add a proper error message
	}
	if err != nil {
		return false, err // Add a proper error message
	}
	return val == "blacklisted", nil
}
