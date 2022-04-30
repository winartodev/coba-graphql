package usecase

import (
	"context"
	"winartodev/coba-graphql/entity"
	"winartodev/coba-graphql/repository"
)

type UserUsecase interface {
	GetUsers(ctx context.Context) ([]entity.User, error)
	GetUserByID(ctx context.Context, ID int) (*entity.User, error)
	CreateUser(ctx context.Context, user entity.User) (*entity.User, error)
	UpdateUserByID(ctx context.Context, ID int, user entity.User) (*entity.User, error)
	DeleteUserByID(ctx context.Context, ID int) error
}

type UserUsecaseProvider struct {
	UserRepo repository.UserRepository
}

func NewUserUsecase(ur repository.UserRepository) UserUsecase {
	return &UserUsecaseProvider{UserRepo: ur}
}

func (ur *UserUsecaseProvider) GetUsers(ctx context.Context) ([]entity.User, error) {
	res, err := ur.UserRepo.GetUsers(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (ur *UserUsecaseProvider) GetUserByID(ctx context.Context, ID int) (*entity.User, error) {
	res, err := ur.UserRepo.GetUserByID(ctx, ID)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (ur *UserUsecaseProvider) CreateUser(ctx context.Context, user entity.User) (*entity.User, error) {
	res, err := ur.UserRepo.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (ur *UserUsecaseProvider) UpdateUserByID(ctx context.Context, ID int, user entity.User) (*entity.User, error) {
	res, err := ur.UserRepo.UpdateUserByID(ctx, ID, user)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (ur *UserUsecaseProvider) DeleteUserByID(ctx context.Context, ID int) error {
	err := ur.UserRepo.DeleteUserByID(ctx, ID)
	if err != nil {
		return err
	}
	return nil
}
