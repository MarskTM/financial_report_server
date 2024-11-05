package dao

import (
	"gorm.io/gorm"
)

type CommonDAO struct {
	db *gorm.DB
}

func (dao *CommonDAO) AdvancedFilter(query string, param interface{}) (interface{}, error) {
	// Implement this method to fetch the object with the given ID from the database
	return nil, nil
}

func (dao *CommonDAO) BasicQuery(query string, param interface{}) (interface{}, error) {
	// Implement this method to execute a basic query and return the result
	return nil, nil
}

func NewCommonDAO(db *gorm.DB) *CommonDAO {
	return &CommonDAO{
		db: db,
	}
}
