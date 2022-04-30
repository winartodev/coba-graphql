package mutation

import (
	"winartodev/coba-graphql/schema"
	"winartodev/coba-graphql/schema/resolver"

	"github.com/graphql-go/graphql"
)

type UserMutation struct {
	UserResolver resolver.UserResolver
}

func NewUserMutation(userResolver resolver.UserResolver) UserMutation {
	return UserMutation{
		UserResolver: userResolver,
	}
}

func (um *UserMutation) Mutation() *graphql.Object {
	objectConfig := graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"Create": &graphql.Field{
				Type:        schema.User,
				Description: "Create New",
				Args: graphql.FieldConfigArgument{
					"name": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: um.UserResolver.CreateUser,
			},
			"Update": &graphql.Field{
				Type:        schema.User,
				Description: "Update User",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"name": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: um.UserResolver.UpdateUserByID,
			},
			"Delete": &graphql.Field{
				Type:        schema.User,
				Description: "Delete User",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: um.UserResolver.DeleteUserByID,
			},
		},
	}

	return graphql.NewObject(objectConfig)
}
