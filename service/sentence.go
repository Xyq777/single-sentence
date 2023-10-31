package service

import (
	"log"
	"single-sentence/repository/dao"
	"single-sentence/repository/model"
	"single-sentence/serializer"
)

func SentenceGet(param string) (sentence *model.Xyq777, err error) {
	sentence, err = dao.GetSentenceRedis(param)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return sentence, nil
}
func SentencePost(req *serializer.SentencePost, username string) (*model.Xyq777, error) {
	var sentence = &model.Xyq777{}
	sentence.NewSentence(0, req.Xyq777, req.Type, req.From, username)

	err := dao.PostSentence(sentence)
	if err != nil {
		log.Panicln(err)
		return nil, err
	}
	err = dao.PostSentenceRedis(sentence)
	if err != nil {
		log.Panicln(err)
		return nil, err
	}
	return sentence, nil
}
func SentenceUpdate(req *serializer.SentenceUpdate, sentenceId uint) (*model.Xyq777, error) {
	var sentence = &model.Xyq777{}
	sentence.NewSentence(sentenceId, req.Xyq777, req.Type, req.From, "")
	sentenceType, err := dao.UpdateSentence(sentence)
	if err != nil {
		log.Panicln(err)
		return nil, err
	}
	err = dao.DeleteSentenceRedis(sentenceId, sentenceType)
	if err != nil {
		log.Panicln(err)
		return nil, err
	}
	err = dao.PostSentenceRedis(sentence)
	if err != nil {
		log.Panicln(err)
		return nil, err
	}
	return sentence, nil
}
func SentenceDelete(sentenceId uint) error {
	sentenceType, err := dao.DeleteSentence(sentenceId)
	if err != nil {
		log.Panicln(err)
		return err
	}
	err = dao.DeleteSentenceRedis(sentenceId, sentenceType)
	if err != nil {
		log.Panicln(err)
		return err
	}
	return nil
}
