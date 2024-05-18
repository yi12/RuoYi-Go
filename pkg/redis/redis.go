// Copyright (c) [2024] K. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Author: K.
// Email: hot_kun@hotmail.com or BusinessCallKun@gmail.com

package ryredis

import (
	"context"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisStruct struct {
	rdb     *redis.Client
	Options *redis.Options

	mu sync.Mutex
}

func (rs *RedisStruct) NewClient() {
	rs.rdb = redis.NewClient(rs.Options)
}

// Get Redis `GET key` command. It returns redis.Nil error when key does not exist.
func (rs *RedisStruct) Get(key string) (string, error) {
	rs.mu.Lock()
	defer rs.mu.Unlock()

	return rs.rdb.Get(context.Background(), key).Result()
}

func (rs *RedisStruct) Set(key string, value interface{}, expiration time.Duration) error {
	rs.mu.Lock()
	defer rs.mu.Unlock()

	return rs.rdb.Set(context.Background(), key, value, expiration).Err()
}

func (rs *RedisStruct) Close() error {
	rs.mu.Lock()
	defer rs.mu.Unlock()

	if rs.rdb != nil {
		// 关闭连接
		return rs.rdb.Close()
	}
	return nil
}