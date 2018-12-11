package mysql

import (
	"errors"

	"github.com/ethereum/go-ethereum/log"
)

// ===========================Basic=================================
func Add(item interface{}) error {
	if item == nil {
		return errors.New("Mysql Add param[item] is nil!")
	}
	if err := DB.Add(item); err != nil {
		log.Error("Mysql Add err:%v", err)
		return err
	}
	return nil
}

func NewRecord(value interface{}) bool {
	return DB.DB.NewRecord(value)
}

// ===========================Test==================================
type Rest struct {
	ID   int64  `gorm:"column:id;" json:"id"`
	Name string `gorm:"column:name;" json:"name"`
}

func TestMysql() ([]Rest, error) {
	var result []Rest
	if err := DB.DB.Raw("select * from d4d_test").Find(&result).Error; err != nil {
		log.Error("TestMysql err:%v", err)
		return result, err
	}
	return result, nil
}
