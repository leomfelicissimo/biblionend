package schemas

import gql "github.com/graphql-go/graphql"

func helloField() *gql.Field {
	return &gql.Field{
		Type: gql.String,
		Resolve: func(p gql.ResolveParams) (interface{}, error) {
			return "world", nil
		},
	}
}
