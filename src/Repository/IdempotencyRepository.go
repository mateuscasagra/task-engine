package Repository

import (
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/task-engine/src"
)

type IdempotencyRepository struct {
	Instance *redis.Client
}

func GetInstanceIdempotency() *IdempotencyRepository {
	rdb := src.Redis()
	return &IdempotencyRepository{
		Instance: rdb,
	}
}

func (i *IdempotencyRepository) Get(ctx *gin.Context, key string) {

	i.Instance.Get(ctx, "idempotency").Result()
}
