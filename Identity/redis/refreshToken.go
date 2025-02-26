package redis

import (
	"context"
	"time"
)

func (rc *dbRedisStore) SetRefreshTokenWithExpiry(ctx context.Context, username, token string, expiry time.Duration) error {
	cmd := rc.client.Set(ctx, username, token, expiry)
	if cmd.Err() != nil {
		return cmd.Err()
	}
	return nil
}

func (rc *dbRedisStore) GetRefreshToken(ctx context.Context, username string) (string, error) {
	cmd := rc.client.Get(ctx, username)
	if cmd.Err() != nil {
		return "", cmd.Err()
	}
	return cmd.Val(), nil
}

func (rc *dbRedisStore) DeleteRefreshToken(ctx context.Context, username string) error {
	cmd := rc.client.Del(ctx, username)
	if cmd.Err() != nil {
		return cmd.Err()
	}
	return nil
}
