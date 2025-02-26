package redis

import "context"

func (rc *dbRedisStore) SetBlackListedAccessToken(ctx context.Context, token string) error {
	// For immediate force logout, Will implement later
	return nil
}

func (rc *dbRedisStore) IsAccessTokenBlackListed(ctx context.Context, token string) (bool, error) {
	// For immediate force logout, Will implement later
	return false, nil
}
