package db

import (
	"context"
	"os"
	"time"

	. "github.com/digitalcircle-com-br/restdb/pkg/log"
	"github.com/go-redis/redis/v8"
)

var cli *redis.Client

func Cli() *redis.Client {
	return cli
}

func Setup() error {
	var err error

	redisurl := os.Getenv("REDIS")
	if redisurl == "" {
		redisurl = "redis://redis:6379"
	}
	opts, err := redis.ParseURL(redisurl)
	if err != nil {
		return err
	}
	cli = redis.NewClient(opts)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	pingCmd := cli.Ping(ctx)

	if pingCmd.Err() != nil {
		return pingCmd.Err()
	}

	res, err := pingCmd.Result()

	if err != nil {
		return err
	}

	Log("Cache connected - %s - %s", redisurl, res)
	return nil
}
