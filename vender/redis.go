package vender

import (
	"bert/shopbert/database"
	"strconv"
)

func RedisSet(key string, e string) (err error) {
	err = database.Rdb.Set(key, e, 0).Err()
	return
}

func RedisGet(key string) (str string, err error) {
	str, err = database.Rdb.Get(key).Result()
	return
}

func RedisLPush(key string, val int) (err error) {
	err = database.Rdb.LPush(key, val).Err()
	return
}

func RedisExists(key string) string {
	str, _ := database.Rdb.Exists(key).Result()
	return strconv.FormatInt(str, 10)

}
