package dao

import (
	"encoding/json"
	"github.com/gomodule/redigo/redis"
	"log"
	"single-sentence/repository/model"
)

func PostSentenceRedis(sentence *model.Xyq777) error {
	jsonData, err := json.Marshal(sentence)
	if err != nil {
		log.Panicln("JSON绑定失败", err)
		return err
	}
	_, err = RedisDao.Do("HSET", sentence.Type, sentence.ID, jsonData)
	if err != nil {
		log.Panicln("插入REDIS数据失败", err)
		return err
	}
	return nil
}
func GetSentenceRedis(param string) (sentence *model.Xyq777, err error) {
	sentence = &model.Xyq777{}
	sentenceId, err := redis.Strings(RedisDao.Do("HRANDFIELD", param, 1))
	if err != nil {
		log.Println("获得REDIS数据失败", err)
		return nil, err

	}
	sentenceJson, err := redis.Bytes(RedisDao.Do("HGET", param, sentenceId[0]))
	if err != nil {
		log.Println("获得REDIS数据失败", err)
		return nil, err
	}
	err = json.Unmarshal(sentenceJson, sentence)
	if err != nil {
		log.Println("解析REDIS数据失败", err)
		return nil, err
	}
	return sentence, nil

}
func DeleteSentenceRedis(sentenceId uint, sentenceType string) (err error) {
	_, err = RedisDao.Do("HDEL", sentenceType, sentenceId)
	if err != nil {
		log.Panicln("删除REDIS数据失败", err)
		return err

	}
	return nil
}
