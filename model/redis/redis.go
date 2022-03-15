package redis

import (
	"context"
	"time"

	"github.com/goworkeryyt/go-core/global"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"

)

type Store struct {
	Expiration time.Duration
	PreKey     string
	Context    context.Context
}

// NewDefaultRedisStore 初始化默认的验证码redis共享存储器
func NewDefaultRedisStore() *Store {
	return &Store{
		Expiration: time.Second * 180,
		PreKey:     "CAPTCHA_",
	}
}

func (rs *Store) UseWithCtx(ctx context.Context) base64Captcha.Store {
	rs.Context = ctx
	return rs
}

func (rs *Store) Set(id string, value string)  error{
	err := global.REDIS.Set(rs.Context, rs.PreKey+id, value, rs.Expiration).Err()
	if err != nil {
		global.LOG.Error("RedisStoreSetError!", zap.Error(err))
		return err
	}
	return nil
}

func (rs *Store) Get(key string, clear bool) string {
	val, err := global.REDIS.Get(rs.Context, key).Result()
	if err != nil {
		global.LOG.Error("RedisStoreGetError!", zap.Error(err))
		return ""
	}
	if clear {
		err := global.REDIS.Del(rs.Context, key).Err()
		if err != nil {
			global.LOG.Error("RedisStoreClearError!", zap.Error(err))
			return ""
		}
	}
	return val
}

func (rs *Store) Verify(id, answer string, clear bool) bool {
	key := rs.PreKey + id
	v := rs.Get(key, clear)
	return v == answer
}