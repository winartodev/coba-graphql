package resolver

import (
	"context"
	"errors"
	"fmt"
	"winartodev/coba-graphql/entity"
	"winartodev/coba-graphql/usecase"

	"github.com/graphql-go/graphql"
)

type UserResolver interface {
	GetUsers(params graphql.ResolveParams) (interface{}, error)
	GetUserByID(params graphql.ResolveParams) (interface{}, error)
	CreateUser(params graphql.ResolveParams) (interface{}, error)
	UpdateUserByID(params graphql.ResolveParams) (interface{}, error)
	DeleteUserByID(params graphql.ResolveParams) (interface{}, error)
}

type UserResolverProvider struct {
	UserUsecase usecase.UserUsecase
}

func NewUserResolver(userUsecase usecase.UserUsecase) UserResolver {
	return &UserResolverProvider{
		UserUsecase: userUsecase,
	}
}

func (ur *UserResolverProvider) GetUsers(params graphql.ResolveParams) (interface{}, error) {
	data, err := ur.UserUsecase.GetUsers(context.Background())
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (ur *UserResolverProvider) GetUserByID(params graphql.ResolveParams) (interface{}, error) {
	id, ok := params.Args["id"].(int)
	if !ok {
		err := errors.New("id must integer ex: 1, 2, etc")
		return nil, err
	} else {
		data, err := ur.UserUsecase.GetUserByID(context.Background(), id)
		if err != nil {
			return nil, err
		}
		return data, nil
	}
}

func (ur *UserResolverProvider) CreateUser(params graphql.ResolveParams) (interface{}, error) {
	user := entity.User{
		Name: params.Args["name"].(string),
	}
	data, err := ur.UserUsecase.CreateUser(context.Background(), user)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (ur *UserResolverProvider) UpdateUserByID(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"].(int)
	user := entity.User{
		Name: params.Args["name"].(string),
	}
	data, err := ur.UserUsecase.UpdateUserByID(context.Background(), id, user)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (ur *UserResolverProvider) DeleteUserByID(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"].(int)
	err := ur.UserUsecase.DeleteUserByID(context.Background(), id)
	if err != nil {
		return nil, err
	}
	return fmt.Sprintf("user is %d successfully deleted", id), nil
}
