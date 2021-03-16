package repositories

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	di "github.com/klasrak/data-integration"
	"github.com/klasrak/data-integration/auth"
)

type RedisAuthRepository struct {
	client *redis.Client
}

var _ auth.AuthInterface = &RedisAuthRepository{}

func NewAuthRepository(client *redis.Client) *RedisAuthRepository {
	return &RedisAuthRepository{client: client}
}

func (rr *RedisAuthRepository) CreateAuth(userID string, td *di.TokenDetails) error {
	at := time.Unix(td.AtExpires, 0) // converting Unix to UTC(to Time object)
	rt := time.Unix(td.RtExpires, 0)
	now := time.Now()

	atCreated, err := rr.client.Set(context.TODO(), td.TokenUUID, userID, at.Sub(now)).Result()

	if err != nil {
		return err
	}

	rtCreated, err := rr.client.Set(context.TODO(), td.RefreshUUID, userID, rt.Sub(now)).Result()

	if err != nil {
		return err
	}

	if atCreated == "0" || rtCreated == "0" {
		return errors.New("no record inserted")
	}

	return nil
}

func (rr *RedisAuthRepository) FetchAuth(tokenUUID string) (string, error) {
	userID, err := rr.client.Get(context.TODO(), tokenUUID).Result()

	if err != nil {
		return "", nil
	}

	return userID, nil
}

func (rr *RedisAuthRepository) DeleteTokens(authD *di.AccessDetails) error {
	// get the refresh uuid
	refreshUUID := fmt.Sprintf("%s++%s", authD.TokenUUID, authD.UserID)
	//delete access token
	deletedAt, err := rr.client.Del(context.TODO(), authD.TokenUUID).Result()

	if err != nil {
		return err
	}

	deletedRt, err := rr.client.Del(context.TODO(), refreshUUID).Result()

	if err != nil {
		return err
	}

	if deletedAt != 1 || deletedRt != 1 {
		return errors.New("something went wrong")
	}

	return nil
}

func (rr *RedisAuthRepository) DeleteRefresh(refreshUUID string) error {
	deleted, err := rr.client.Del(context.TODO(), refreshUUID).Result()

	if err != nil || deleted == 0 {
		return err
	}

	return nil
}
