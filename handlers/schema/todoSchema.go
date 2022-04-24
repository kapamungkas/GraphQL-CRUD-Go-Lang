package schema

import (
	"github.com/graphql-go/graphql"
)

func NewTodoSchema(queryType *graphql.Object, mutationType *graphql.Object) graphql.Schema {
	value, _ := graphql.NewSchema(
		graphql.SchemaConfig{
			Query:    queryType,
			Mutation: mutationType,
		},
	)
	return value
}
