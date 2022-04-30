package repository

import (
	"context"
	"winartodev/coba-graphql/entity"

	"github.com/jinzhu/gorm"
)

type UserRepository interface {
	GetUsers(ctx context.Context) ([]entity.User, error)
	GetUserByID(ctx context.Context, ID int) (*entity.User, error)
	CreateUser(ctx context.Context, user entity.User) (*entity.User, error)
	UpdateUserByID(ctx context.Context, ID int, user entity.User) (*entity.User, error)
	DeleteUserByID(ctx context.Context, ID int) error
}

type storeData struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &storeData{
		db: db,
	}
}

func (db *storeData) GetUsers(ctx context.Context) ([]entity.User, error) {
	users := []entity.User{}

	err := db.db.Model(&entity.User{}).Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (db *storeData) GetUserByID(ctx context.Context, ID int) (*entity.User, error) {
	user := entity.User{}

	err := db.db.Model(&entity.User{}).Where("id = ?", ID).Find(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (db *storeData) CreateUser(ctx context.Context, user entity.User) (*entity.User, error) {
	err := db.db.Model(&entity.User{}).Create(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (db *storeData) UpdateUserByID(ctx context.Context, ID int, user entity.User) (*entity.User, error) {
	err := db.db.Model(&entity.User{}).Select("name").Where("id = ?", ID).Updates(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (db *storeData) DeleteUserByID(ctx context.Context, ID int) error {
	err := db.db.Model(&entity.User{}).Where("id = ?", ID).Delete(&entity.User{}).Error
	if err != nil {
		return err
	}
	return nil
}
