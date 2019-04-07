package schemas

import (
	"log"

	gql "github.com/graphql-go/graphql"
)

func bibleTextObject() *gql.Object {
	objectConfig := gql.ObjectConfig{
		Name: "BibleText",
		Fields: gql.Fields{
			"reference": &gql.Field{Type: gql.String},
			"book":      &gql.Field{Type: gql.String},
			"chapter":   &gql.Field{Type: gql.Int},
			"verse":     &gql.Field{Type: gql.Int},
			"text":      &gql.Field{Type: gql.String},
		},
	}

	return gql.NewObject(objectConfig)
}

func bibleTextField() *gql.Field {
	return &gql.Field{
		Type: &gql.List{OfType: bibleTextObject()},
		Resolve: func(p gql.ResolveParams) (interface{}, error) {
			reference := p.Args["reference"]
			log.Println("Reference:", reference)
			return "Ok", nil
		},
		Args: gql.FieldConfigArgument{
			"reference": &gql.ArgumentConfig{Type: gql.String},
		},
	}
}
