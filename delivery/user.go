package delivery

import (
	"errors"
	"net/http"
	"winartodev/coba-graphql/handler"
	"winartodev/coba-graphql/response"
	"winartodev/coba-graphql/schema"
	"winartodev/coba-graphql/schema/mutation"
	"winartodev/coba-graphql/schema/query"
	"winartodev/coba-graphql/usecase"

	"github.com/graphql-go/graphql"
	"github.com/julienschmidt/httprouter"
)

type UserHandler struct {
	uc usecase.UserUsecase
	uq query.UserQuery
	um mutation.UserMutation
	us schema.UserSchema
}

func NewUserHandler(userUsecase usecase.UserUsecase, userQuery query.UserQuery, userMutation mutation.UserMutation, userSchema schema.UserSchema) UserHandler {
	return UserHandler{
		uc: userUsecase,
		uq: userQuery,
		um: userMutation,
		us: userSchema,
	}
}

func (uh *UserHandler) Register(r *httprouter.Router) error {
	if r == nil {
		return errors.New("router cannot be empty")
	}

	r.GET("/users", handler.Decorate(uh.User))
	r.POST("/users", handler.Decorate(uh.User))
	return nil
}

func (uh *UserHandler) User(w http.ResponseWriter, r *http.Request, _ httprouter.Params) error {
	schema, _ := graphql.NewSchema(
		graphql.SchemaConfig{
			Query:    uh.uq.Query(),
			Mutation: uh.um.Mutation(),
		},
	)

	res := uh.us.ExeuteUserQuery(r.URL.Query().Get("query"), schema)
	if res.Errors != nil {
		response.HttpResponseError(w, r, http.StatusInternalServerError, res.Errors)
	} else {
		response.HttpResponseSuccess(w, r, res.Data)
	}
	return nil
}
