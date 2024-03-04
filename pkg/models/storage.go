package models

import (
	"encoding/json"
	"fmt"
	"time"
)

func GetGoods(limit int,offset int)([]byte,error){
	key := fmt.Sprintf("goods:%d:%d",limit,offset)
	data,_ := GetGoodsFromRedis(key)
	if data != ""{
		return []byte(data),nil
	}
	goods,meta,err := GetGoodsFromDB(limit,offset)
	if err !=nil{
		return nil,err
	}
    jsonData, err := json.Marshal(map[string]interface{}{
        "goods": goods,
        "meta":  meta,
    })
    if err != nil {
        return nil,err
    }
	if err := CacheGoodsToRedis(key, jsonData, time.Minute); err != nil {
        fmt.Println("Error caching data in Redis:", err)
    }
	return jsonData, nil
}
func RemoveGood(id int, project_id int) error {
	if err := RemoveGoodFromDb(id,project_id); err != nil{
		return err
	}
	if err := RemoveGoodFromRedis(id, project_id); err!=nil{
		return err
	}
	return nil
}