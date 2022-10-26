package repository

import (
	"github.com/mazeem91/trackman-poc/infrastructure/logger"
	"gorm.io/gorm"
)

type SQLite struct {
	DB *gorm.DB
}

func (repo *SQLite) Save(model interface{}) error {
	err := repo.DB.Create(model).Error
	if err != nil {
		logger.Errorf("error, not save data %v", err)
	}
	return err
}

func (repo *SQLite) Get(model interface{}) error {
	err := repo.DB.Find(model).Error
	return err
}

func (repo *SQLite) GetWith(model interface{}, join string) interface{} {
	err := repo.DB.Joins(join).Find(model).Error
	return err
}

func (repo *SQLite) GetBy(model interface{}, conditions interface{}) error {
	err := repo.DB.Where(conditions).First(&model).Error
	return err
}

func (repo *SQLite) GetOne(model interface{}) interface{} {
	err := repo.DB.Last(model).Error
	return err
}

func (repo *SQLite) Update(model interface{}) interface{} {
	err := repo.DB.Find(model).Error
	return err
}
