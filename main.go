package main

import (
	"fmt"
	"graphql-todo/handlers/mutation"
	"graphql-todo/handlers/query"
	"graphql-todo/handlers/schema"

	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
)

func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("errors: %v", result.Errors)
	}
	return result
}

func main() {
	router := gin.Default()

	schema := schema.NewTodoSchema(query.QueryType, mutation.MutationType)

	router.POST("/todo", func(c *gin.Context) {

		Query := c.PostForm("query")

		// fmt.Printf("Query: %s;", Query)

		result := executeQuery(Query, schema)

		c.JSON(200, result)
	})
	router.Run(":8080")
}
