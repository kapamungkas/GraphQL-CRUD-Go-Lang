package query

import (
	"graphql-todo/database"
	"graphql-todo/repositories"
	"graphql-todo/services"
	"graphql-todo/types"

	"github.com/graphql-go/graphql"
)

var db = database.ConnectDB()

var todo_repository = repositories.NewTodoRepository(db)
var todo_service = services.NewTodoService(todo_repository)

var QueryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			/* Get (read) single product by id
			   http://localhost:8080/product?query={product(id:1){name,info,price}}
			*/
			"product": &graphql.Field{
				Type:        types.ProductType,
				Description: "Get product by id",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id, ok := p.Args["id"].(int)
					if ok {
						// Find product
						return todo_service.GetByIDTodo(id)
					}
					return nil, nil
				},
			},
			/* Get (read) product list
			   http://localhost:8080/product?query={list{id,name,info,price}}
			*/
			"list": &graphql.Field{
				Type:        graphql.NewList(types.ProductType),
				Description: "Get product list",
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					return todo_service.GetAllTodo()
				},
			},
		},
	})
