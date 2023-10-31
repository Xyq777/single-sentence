package dao

import (
	"log"
	"single-sentence/repository/model"
)

func PostSentence(sentence *model.Xyq777) (err error) {
	if err = MysqlDao.Create(sentence).Error; err != nil {
		log.Panicln("获取数据库数据失败", err)
		return err
	}
	return nil
}
func GetAllSentences(username string) ([]model.Xyq777, error) {
	var sentences []model.Xyq777
	if err := MysqlDao.Find(&sentences, "creator = ?", username).Error; err != nil {
		log.Println("获取数据库数据失败", err)
		return nil, err
	}

	return sentences, nil
}
func GetSentenceById(sentenceId uint) (*model.Xyq777, error) {
	var sentence = &model.Xyq777{}
	if err := MysqlDao.Find(&sentence, "id = ?", sentenceId).Error; err != nil {
		log.Println("获取数据库数据失败", err)
		return nil, err
	}

	return sentence, nil
}
func UpdateSentence(sentence *model.Xyq777) (sentenceType string, err error) {
	//从redis删数据需要id和句子类型。。。句子类型还得从数据库中拿，感觉有点麻烦，但目前感觉只能这样了。。
	var oldSentence model.Xyq777
	oldSentence.ID = sentence.ID
	if err = MysqlDao.Find(&oldSentence).Error; err != nil {
		log.Println("寻找数据库数据失败", err)
		return "", err
	}
	oldSentence.From = sentence.From
	oldSentence.Type = sentence.Type
	oldSentence.Xyq777 = sentence.Xyq777

	if err = MysqlDao.Save(&oldSentence).Error; err != nil {
		log.Panicln("更新数据库数据失败", err)
		return "", err
	}

	return oldSentence.Type, nil
}
func DeleteSentence(sentenceId uint) (sentenceType string, err error) {
	var sentence model.Xyq777
	sentence.ID = sentenceId
	if err = MysqlDao.First(&sentence).Error; err != nil {
		log.Println("获取数据库数据失败", err)
		return "", err
	}
	sentenceType = sentence.Type

	if err = MysqlDao.Delete(&sentence).Error; err != nil {
		log.Panicln("删除数据库数据失败", err)
		return "", err
	}

	return sentenceType, nil
}
