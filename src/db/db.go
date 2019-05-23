package db

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"time"
)

func NewPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:   80,
		MaxActive: 12000, // max number of connections
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", ":6379")
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},
	}

}

var pool = NewPool()
var Conn = pool.Get()

func Get(key string) (string, error) {
	Conn = pool.Get()
	value, err := redis.Bytes(Conn.Do("get", key))
	defer Conn.Close()
	return string(value), err
}

func Set(key string, value string) error {
	Conn = pool.Get()
	_, err := Conn.Do("set", key, value)
	defer Conn.Close()
	return err
}

func Del(key string) error {
	Conn = pool.Get()
	_, err := Conn.Do("del", key)
	defer Conn.Close()
	return err
}

func HGet(hash string, key string) (string, error) {
	Conn = pool.Get()
	value, err := redis.Bytes(Conn.Do("hget", hash, key))
	defer Conn.Close()
	return string(value), err
}

func HSet(hash string, key string, value string) error {
	Conn = pool.Get()
	_, err := Conn.Do("hset", hash, key, value)
	defer Conn.Close()
	return err
}

func HDel(hash string, key string) error {
	Conn = pool.Get()
	_, err := Conn.Do("hdel", hash, key)
	defer Conn.Close()
	return err
}

func HGetAll(hash string) (map[string]string, error) {
	Conn = pool.Get()
	values, err := redis.StringMap(Conn.Do("hgetall", hash))
	defer Conn.Close()
	return values, err

}

func ZRangeByScore(hash string) ([][]byte, error) {
	Conn = pool.Get()
	logsRaw, err := redis.ByteSlices(Conn.Do("zRangeByScore", hash, "-inf", "+inf"))
	defer Conn.Close()
	return logsRaw, err
}

func ZAdd(hash string, value string) error {
	Conn = pool.Get()
	_, err := Conn.Do("zadd", hash, time.Now().Unix(), value)
	defer Conn.Close()
	return err
}

func ZRem(hash string, key string) error {
	Conn = pool.Get()
	_, err := Conn.Do("zrem", hash, key)
	defer Conn.Close()
	return err
}

func GenerateTime(hash string) (int64, error) {
	Conn = pool.Get()
	value, err := redis.Int64(Conn.Do("get", hash))
	defer Conn.Close()
	return value, err
}

func LPush(listname string, list []string) {
	Conn = pool.Get()
	defer Conn.Close()
	for _, listval := range list {
		_, err := Conn.Do("lpush", listname, listval)
		if err != nil {
			fmt.Println("error in pushing list", err)
		}
	}
}

func LRange(list string) []string {
	Conn = pool.Get()
	defer Conn.Close()
	charlist, err := redis.Strings(Conn.Do("lrange", list, 0, -1))
	if err != nil {
		fmt.Println("Lrange err", err)
	}
	return charlist
}
