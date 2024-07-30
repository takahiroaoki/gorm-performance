package dao

import (
	"github.com/takahiroaoki/gorm-performance/app/entity"
	"gorm.io/gorm"
)

type RecordDao struct{}

func (r *RecordDao) InsertRecord(db *gorm.DB, rec entity.Record) (*entity.Record, error) {
	if err := db.Create(&rec).Error; err != nil {
		return nil, err
	}
	return &rec, nil
}

func NewRecordDao() *RecordDao {
	return &RecordDao{}
}