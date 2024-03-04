package models

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/redis/go-redis/v9"
)

var rdb *redis.Client
var ctx = context.Background()

func InitRedis(host string) {
	rdb = redis.NewClient(&redis.Options{Addr: host})
}

func CacheGoodsToRedis(key string, value []byte, expiration time.Duration)error{
	return rdb.Set(ctx, key, value, expiration).Err()
}
func GetGoodsFromRedis(key string)(string,error) {
	val, err := rdb.Get(ctx, key).Result()
	if err != nil && err != redis.Nil {
		return "", err
	}
	return val, nil
}

func RemoveGoodFromRedis(id_remove int, project_id_remove int)error{
	keys,err := rdb.Keys(ctx,"*").Result()
	if err != nil{
		return err
	}
	for _,key := range keys{
		if key != ""{
			val,_ := rdb.Get(ctx,key).Result()
			var good map[string]interface{}
			if err:=json.Unmarshal([]byte(val),&good);err!=nil{
				return err
			}
			for _, v := range good["goods"].([]interface{}) {
				if item, ok := v.(map[string]interface{}); ok {
					id, ok := item["id"].(float64)
					if ok {
						project_id := item["projectId"].(float64)
						if id == float64(id_remove) && project_id == float64(project_id_remove){
							rdb.Del(ctx,key)
						}
					}else{
						return errors.New("проблема с поиском")
					}
				}
			}
		}
	}
	return nil
}