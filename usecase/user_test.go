package usecase_test

import (
	"context"
	"errors"
	"testing"
	"winartodev/coba-graphql/entity"
	"winartodev/coba-graphql/mocks"
	"winartodev/coba-graphql/usecase"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockUserProvider struct {
	UserRepo *mocks.UserRepository
}

func newUserProvider() mockUserProvider {
	return mockUserProvider{
		UserRepo: new(mocks.UserRepository),
	}
}

func newUserMocksUsecase(repo *usecase.UserUsecaseProvider) usecase.UserUsecase {
	return usecase.NewUserUsecase(repo)
}

func Test_GetUsers(t *testing.T) {
	testcases := []struct {
		name      string
		data      []entity.User
		wantError bool
		err       error
	}{
		{
			name: "success",
			data: []entity.User{},
		},
		{
			name:      "failed",
			data:      nil,
			wantError: true,
			err:       errors.New("failed test"),
		},
	}

	for _, test := range testcases {
		t.Run(test.name, func(t *testing.T) {
			userMock := newUserProvider()
			newUserUsecase := newUserMocksUsecase(&usecase.UserUsecaseProvider{UserRepo: userMock.UserRepo})

			userMock.UserRepo.On("GetUsers", mock.Anything).Return(test.data, test.err)

			res, err := newUserUsecase.GetUsers(context.Background())
			if err != nil && test.wantError {
				assert.Nil(t, res)
				assert.Error(t, err)
			} else {
				assert.NotNil(t, res)
				assert.NoError(t, err)
			}
		})
	}
}

func Test_GetUserByID(t *testing.T) {
	testcases := []struct {
		name      string
		ID        int
		data      *entity.User
		wantError bool
		err       error
	}{
		{
			name: "success",
			ID:   1,
			data: &entity.User{},
		},
		{
			name:      "failed test",
			ID:        1,
			data:      &entity.User{},
			wantError: true,
			err:       errors.New("failed test"),
		},
	}

	for _, test := range testcases {
		t.Run(test.name, func(t *testing.T) {
			userMock := newUserProvider()
			newUserUsecase := newUserMocksUsecase(&usecase.UserUsecaseProvider{UserRepo: userMock.UserRepo})

			userMock.UserRepo.On("GetUserByID", mock.Anything, mock.AnythingOfType("int")).Return(test.data, test.err)

			res, err := newUserUsecase.GetUserByID(context.Background(), test.ID)
			if err != nil && test.wantError {
				assert.Nil(t, res)
				assert.Error(t, err)
			} else {
				assert.NotNil(t, res)
				assert.NoError(t, err)
			}
		})
	}
}

func Test_CreateUser(t *testing.T) {
	testcases := []struct {
		name      string
		data      entity.User
		wantError bool
		err       error
	}{
		{
			name: "success",
			data: entity.User{},
		},
		{
			name:      "failed",
			data:      entity.User{},
			wantError: true,
			err:       errors.New("failed test"),
		},
	}

	for _, test := range testcases {
		t.Run(test.name, func(t *testing.T) {
			userMock := newUserProvider()
			newUserUsecase := newUserMocksUsecase(&usecase.UserUsecaseProvider{UserRepo: userMock.UserRepo})

			userMock.UserRepo.On("CreateUser", mock.Anything, test.data).Return(&test.data, test.err)
			res, err := newUserUsecase.CreateUser(context.Background(), test.data)
			if err != nil && test.wantError {
				assert.Nil(t, res)
				assert.Error(t, err)
			} else {
				assert.NotNil(t, res)
				assert.NoError(t, err)
			}
		})
	}
}

func Test_UpdateUserByID(t *testing.T) {
	testcases := []struct {
		name      string
		ID        int
		data      entity.User
		wantError bool
		err       error
	}{
		{
			name: "success",
			ID:   1,
			data: entity.User{},
		},
		{
			name:      "failed",
			ID:        1,
			data:      entity.User{},
			wantError: true,
			err:       errors.New("failed test"),
		},
	}

	for _, test := range testcases {
		t.Run(test.name, func(t *testing.T) {
			userMock := newUserProvider()
			userUsecase := newUserMocksUsecase(&usecase.UserUsecaseProvider{UserRepo: userMock.UserRepo})

			userMock.UserRepo.On("UpdateUserByID", mock.Anything, mock.AnythingOfType("int"), test.data).Return(&test.data, test.err)
			res, err := userUsecase.UpdateUserByID(context.Background(), test.ID, test.data)
			if err != nil && test.wantError {
				assert.Nil(t, res)
				assert.Error(t, err)
			} else {
				assert.NotNil(t, res)
				assert.NoError(t, err)
			}
		})
	}
}

func Test_DeleteUserByID(t *testing.T) {
	testcases := []struct {
		name      string
		ID        int
		wantError bool
		err       error
	}{
		{
			name: "success",
			ID:   1,
		},
		{
			name:      "failed",
			ID:        1,
			wantError: true,
			err:       errors.New("failed test"),
		},
	}

	for _, test := range testcases {
		t.Run(test.name, func(t *testing.T) {
			userMock := newUserProvider()
			userUsecase := newUserMocksUsecase(&usecase.UserUsecaseProvider{UserRepo: userMock.UserRepo})

			userMock.UserRepo.On("DeleteUserByID", mock.Anything, mock.AnythingOfType("int")).Return(test.err)

			err := userUsecase.DeleteUserByID(context.Background(), test.ID)
			if err != nil && test.wantError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
