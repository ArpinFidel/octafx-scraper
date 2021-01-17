package service

import (
	"github.com/ArpinFidel/octafx-scraper/model"
	"github.com/ArpinFidel/octafx-scraper/repo"
)

type Master interface {
	Insert(master model.Master) (model.Master, error)
	Update(cond model.Master, master model.Master) (model.Master, error)
}

func (svc *master) Insert(master model.Master) (model.Master, error) {
	return svc.repo.Insert(master)
}

func (svc *master) Update(cond model.Master, master model.Master) (model.Master, error) {
	return svc.repo.Update(cond, master)
}

type master struct{ repo repo.MasterHandler }

var masterSvc *master

func NewMaster(repo repo.MasterHandler) master {
	if masterSvc == nil {
		masterSvc = &master{repo}
	}
	return *masterSvc
}
