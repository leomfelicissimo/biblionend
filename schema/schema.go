package schema

import (
	gql "github.com/graphql-go/graphql"
)

func helloField() *gql.Field {
	return &gql.Field{
		Type: gql.String,
		Resolve: func(p gql.ResolveParams) (interface{}, error) {
			return "world", nil
		},
	}
}

/*
type Query {
    hello: String
}
*/
func queryObject() *gql.Object {
	objectConfig := gql.ObjectConfig{
		Name: "Query",
		Fields: gql.Fields{
			"hello": helloField(),
		},
	}

	return gql.NewObject(objectConfig)
}

// CreateSchema creates a new Graphql-Go Schema base on a sort of defined objects
func CreateSchema() (gql.Schema, error) {
	schemaConfig := gql.SchemaConfig{
		Query: queryObject(),
	}

	schema, err := gql.NewSchema(schemaConfig)
	return schema, err
}
