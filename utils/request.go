package utils

import (
	"context"
	"github.com/go-redis/redis/v8"
)

type RequestInfo struct {
	Ip    string
	Api   string
	Redis *redis.Client
	Ctx   context.Context
}

// IsRepeat 是否在短时间内重复请求
func (r *RequestInfo) IsRepeat() bool {
	//r.Redis.Set(r.Ctx, r.Ip+r.Api, 1, 30)
	//return r.Redis.Get(r.Ctx, r.Ip+r.Api).Val() == "1"
	oldVal, err := r.Redis.Get(r.Ctx, r.Ip+r.Api).Int()

	if err != nil {
		r.Redis.Set(r.Ctx, r.Ip+r.Api, 1, 30)
		return false
	}
	if oldVal <= 1 {
		r.Redis.Set(r.Ctx, r.Ip+r.Api, oldVal+1, 30)
		return false
	} else if oldVal >= 10 {
		//redis 发布订阅
		r.Redis.Subscribe(r.Ctx, r.Ip+r.Api)

		return true
	}
	return false
}
