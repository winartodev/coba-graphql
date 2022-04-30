package repository

import (
	"context"
	"errors"
	"regexp"
	"testing"
	"winartodev/coba-graphql/entity"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_GetUsers(t *testing.T) {
	testcases := []struct {
		name      string
		query     string
		rows      []entity.User
		wantError bool
		err       error
	}{
		{
			name:  "success",
			query: `SELECT * FROM "users"`,
			rows: []entity.User{
				{
					ID:   1,
					Name: "budi",
				},
			},
		},
		{
			name:      "failed",
			query:     `SELECT * FROM "users"`,
			wantError: true,
			err:       errors.New("failed test"),
		},
	}

	for _, test := range testcases {
		t.Run(test.name, func(t *testing.T) {
			mockDB, mock, err := sqlmock.New()
			require.NoError(t, err)

			gormDB, err := gorm.Open("postgres", mockDB)
			require.NoError(t, err)

			repository := NewUserRepository(gormDB)

			if !test.wantError {
				rows := sqlmock.NewRows([]string{"id", "name"})
				for _, row := range test.rows {
					rows.AddRow(row.ID, row.Name)
				}
				mock.ExpectQuery(regexp.QuoteMeta(test.query)).WillReturnRows(rows)
			} else {
				mock.ExpectQuery(regexp.QuoteMeta(test.query)).WillReturnError(test.err)
			}

			res, err := repository.GetUsers(context.Background())
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
		query     string
		row       entity.User
		ID        int
		wantError bool
		err       error
	}{
		{
			name:  "success",
			query: `SELECT * FROM "users" WHERE (id = $1)`,
			ID:    1,
			row: entity.User{
				ID:   1,
				Name: "winarto",
			},
		},
		{
			name:      "failed",
			query:     `SELECT * FROM "users" WHERE (id = $1)`,
			ID:        1,
			row:       entity.User{},
			wantError: true,
			err:       errors.New("failed test"),
		},
	}

	for _, test := range testcases {
		t.Run(test.name, func(t *testing.T) {
			mockDB, mock, err := sqlmock.New()
			require.NoError(t, err)

			gormDB, err := gorm.Open("postgres", mockDB)
			require.NoError(t, err)

			repository := NewUserRepository(gormDB)

			if !test.wantError {
				mock.ExpectQuery(regexp.QuoteMeta(test.query)).WithArgs(test.ID).WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(test.row.ID, test.row.Name))
			} else {
				mock.ExpectQuery(regexp.QuoteMeta(test.query)).WithArgs(test.ID).WillReturnError(test.err)
			}

			res, err := repository.GetUserByID(context.Background(), test.ID)
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
		query     string
		row       entity.User
		wantError bool
		err       error
	}{
		{
			name:  "success",
			query: `INSERT INTO "users" ("id","name") VALUES ($1,$2) RETURNING "users"."id"`,
			row: entity.User{
				ID:   1,
				Name: "winarto",
			},
		},
		{
			name:      "failed",
			query:     `INSERT INTO "users" ("name") VALUES ($1) RETURNING "users"."id"`,
			row:       entity.User{},
			wantError: true,
			err:       errors.New("failed test"),
		},
	}

	for _, test := range testcases {
		t.Run(test.name, func(t *testing.T) {
			mockDB, mock, err := sqlmock.New()
			require.NoError(t, err)

			gormDB, err := gorm.Open("postgres", mockDB)
			require.NoError(t, err)

			repository := NewUserRepository(gormDB)

			mock.ExpectBegin()
			if !test.wantError {
				mock.ExpectQuery(regexp.QuoteMeta(test.query)).WithArgs(test.row.ID, test.row.Name).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(test.row.ID))
			} else {
				mock.ExpectQuery(regexp.QuoteMeta(test.query)).WithArgs(test.row.Name).WillReturnError(test.err)
			}
			mock.ExpectCommit()

			res, err := repository.CreateUser(context.Background(), test.row)
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
		query     string
		ID        int
		row       entity.User
		wantError bool
		err       error
	}{
		{
			name:  "success",
			query: `UPDATE "users" SET "name" = $1 WHERE (id = $2)`,
			ID:    1,
			row: entity.User{
				ID:   1,
				Name: "jono",
			},
		},
		{
			name:  "failed",
			query: `UPDATE "users" SET "name" = $1 WHERE (id = $2)`,
			ID:    1,
			row: entity.User{
				ID:   1,
				Name: "jono",
			},
			wantError: true,
			err:       errors.New("failed test"),
		},
	}

	for _, test := range testcases {
		t.Run(test.name, func(t *testing.T) {
			mockDB, mock, err := sqlmock.New()
			require.NoError(t, err)

			gormDB, err := gorm.Open("postgres", mockDB)
			require.NoError(t, err)

			repository := NewUserRepository(gormDB)

			mock.ExpectBegin()
			if !test.wantError {
				mock.ExpectExec(regexp.QuoteMeta(test.query)).WithArgs(test.row.Name, test.ID).WillReturnResult(sqlmock.NewResult(0, 1))
			} else {
				mock.ExpectExec(regexp.QuoteMeta(test.query)).WithArgs(test.row.Name, test.ID).WillReturnError(test.err)
			}
			mock.ExpectCommit()

			res, err := repository.UpdateUserByID(context.Background(), test.ID, test.row)
			if err != nil && test.wantError {
				assert.Error(t, err)
				assert.Nil(t, res)
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
		query     string
		ID        int
		wantError bool
		err       error
	}{
		{
			name:  "success",
			query: `DELETE FROM "users" WHERE (id = $1)`,
			ID:    1,
		},
		{
			name:      "failed",
			query:     `DELETE FROM "users" WHERE (id = $1)`,
			ID:        1,
			wantError: true,
			err:       errors.New("failed test"),
		},
	}

	for _, test := range testcases {
		t.Run(test.name, func(t *testing.T) {
			mockDB, mock, err := sqlmock.New()
			require.NoError(t, err)

			gormDB, err := gorm.Open("postgres", mockDB)
			require.NoError(t, err)

			repository := NewUserRepository(gormDB)

			mock.ExpectBegin()
			if !test.wantError {
				mock.ExpectExec(regexp.QuoteMeta(test.query)).WithArgs(test.ID).WillReturnResult(sqlmock.NewResult(0, 0))
			} else {
				mock.ExpectExec(regexp.QuoteMeta(test.query)).WithArgs(test.ID).WillReturnError(test.err)
			}
			mock.ExpectCommit()

			err = repository.DeleteUserByID(context.Background(), test.ID)

			if err != nil && test.wantError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
