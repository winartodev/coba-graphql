package schema

import (
	"fmt"

	"github.com/graphql-go/graphql"
)

var User = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
	},
})

type UserSchema interface {
	ExeuteUserQuery(query string, schema graphql.Schema) *graphql.Result
}

type userSchema struct{}

func NewUserSchema() UserSchema {
	return &userSchema{}
}

func (us *userSchema) ExeuteUserQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})

	if len(result.Errors) > 0 {
		fmt.Printf("errors: %v\n", result.Errors)
	}
	return result
}
