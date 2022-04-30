package query

import (
	"winartodev/coba-graphql/schema"
	"winartodev/coba-graphql/schema/resolver"

	"github.com/graphql-go/graphql"
)

type UserQuery struct {
	userResolver resolver.UserResolver
}

func NewUserQuery(userResolver resolver.UserResolver) UserQuery {
	return UserQuery{
		userResolver: userResolver,
	}
}

func (uq UserQuery) Query() *graphql.Object {
	objectConfig := graphql.ObjectConfig{
		Name: "Quer",
		Fields: graphql.Fields{
			"FetchAll": &graphql.Field{
				Type:        graphql.NewList(schema.User),
				Description: "Fetch All",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"name": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: uq.userResolver.GetUsers,
			},
			"FetchByID": &graphql.Field{
				Type:        schema.User,
				Description: "Fetch All",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: uq.userResolver.GetUserByID,
			},
		},
	}

	return graphql.NewObject(objectConfig)
}
