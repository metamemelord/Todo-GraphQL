package graphql

import (
	"errors"
	"time"
	"todo-graph/db"

	"github.com/graphql-go/graphql"
)

var todoType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Todo",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"title": &graphql.Field{
				Type: graphql.String,
			},
			"body": &graphql.Field{
				Type: graphql.String,
			},
			"time": &graphql.Field{
				Type: graphql.DateTime,
			},
		},
	},
)

func NewGraphQLQueryObject() *graphql.Object {
	query := graphql.NewObject(
		graphql.ObjectConfig{
			Name: "RootQuery",
			Fields: graphql.Fields{
				"todo": &graphql.Field{
					Type:        todoType,
					Description: "Get product by id",
					Args: graphql.FieldConfigArgument{
						"id": &graphql.ArgumentConfig{
							Type: graphql.Int,
						},
					},
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						if id, ok := p.Args["id"].(int); ok {
							return db.FindByID(id)
						}
						return nil, nil
					},
				},
				"todos": &graphql.Field{
					Type:        graphql.NewList(todoType),
					Description: "Get all todos",
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						return db.Find()
					},
				},
			},
		},
	)
	return query
}

func NewGraphQLMutationObject() *graphql.Object {
	mutation := graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Mutation",
			Fields: graphql.Fields{
				"add": &graphql.Field{
					Type:        todoType,
					Description: "Add new todo",
					Args: graphql.FieldConfigArgument{
						"title": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.String),
						},
						"body": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.String),
						},
						"time": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.DateTime),
						},
					},
					Resolve: func(params graphql.ResolveParams) (interface{}, error) {
						t := params.Args["time"].(time.Time)
						todo := &db.Todo{
							Title: params.Args["title"].(string),
							Body:  params.Args["body"].(string),
							Time:  t,
						}

						return db.AddTodo(todo)
					},
				},
				"update": &graphql.Field{
					Type:        todoType,
					Description: "Update an existing todo",
					Args: graphql.FieldConfigArgument{
						"id": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.Int),
						},
						"title": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"body": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"time": &graphql.ArgumentConfig{
							Type: graphql.DateTime,
						},
					},
					Resolve: func(params graphql.ResolveParams) (interface{}, error) {
						t := params.Args["time"].(time.Time)
						var todo *db.Todo
						id, ok := params.Args["id"].(int)
						if ok {
							todo = &db.Todo{
								ID:    id,
								Title: params.Args["title"].(string),
								Body:  params.Args["body"].(string),
								Time:  t,
							}
						}
						return todo, db.UpdateTodo(id, todo)
					},
				},
				"delete": &graphql.Field{
					Type:        todoType,
					Description: "Delete a todo by ID",
					Args: graphql.FieldConfigArgument{
						"id": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.Int),
						},
					},
					Resolve: func(params graphql.ResolveParams) (interface{}, error) {
						id, ok := params.Args["id"].(int)
						if ok {
							return nil, db.DeleteTodo(id)
						}
						return nil, errors.New("Error while parsing the todo ID")
					},
				},
			},
		},
	)
	return mutation
}
