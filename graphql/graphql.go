package graphql

import (
	"fmt"
	"sync"

	"github.com/graphql-go/graphql"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(getGraphQLSchema),
	fx.Invoke(initiaizeSchema),
)

var schema graphql.Schema
var once sync.Once

func initiaizeSchema(s graphql.Schema) graphql.Schema {
	once.Do(func() {
		schema = s
	})
	return schema
}

func getGraphQLSchema() (graphql.Schema, error) {
	return graphql.NewSchema(
		graphql.SchemaConfig{
			Query:    NewGraphQLQueryObject(),
			Mutation: NewGraphQLMutationObject(),
		},
	)
}

func ExecuteQuery(query string) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("errors: %v", result.Errors)
	}
	return result
}
