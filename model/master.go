package model

import "gorm.io/gorm"

type Master struct {
	gorm.Model
	Code     string
	Name     string
	Balance  float64
	MinFloat float64
	MaxFloat float64
}
