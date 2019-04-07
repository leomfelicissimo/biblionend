package schemas

import (
	gql "github.com/graphql-go/graphql"
)

/*
type Query {
    hello: String
}
*/
func queryObject() *gql.Object {
	objectConfig := gql.ObjectConfig{
		Name: "Query",
		Fields: gql.Fields{
			"hello":     helloField(),
			"book":      bookField(),
			"bibleText": bibleTextField(),
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
