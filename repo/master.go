package repo

import (
	"github.com/ArpinFidel/octafx-scraper/model"
	"gorm.io/gorm"
)

type MasterHandler interface {
	Insert(master model.Master) (model.Master, error)
	Update(cond model.Master, master model.Master) (model.Master, error)
}

func (repo *masterSQL) Insert(master model.Master) (model.Master, error) {
	return master, repo.db.Begin().Create(&master).Error
}

func (repo *masterSQL) Update(cond model.Master, master model.Master) (model.Master, error) {
	err := repo.db.Begin().Model(cond).Updates(&master).Error
	return master, err
}

type masterSQL struct{ db *gorm.DB }

var masterSQLRepo *masterSQL

func NewMasterSQL(db *gorm.DB) masterSQL {
	if masterSQLRepo == nil {
		masterSQLRepo = &masterSQL{db}
		db.AutoMigrate(&model.Master{})
	}
	return *masterSQLRepo
}
