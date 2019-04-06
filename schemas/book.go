package schemas

import (
	"log"

	gql "github.com/graphql-go/graphql"
	"github.com/leomfelicissimo/biblionend/repositories"
)

/*
type Book {
    name: String
    abbreviation: String
}
*/
func bookObject() *gql.Object {
	objectConfig := gql.ObjectConfig{
		Name: "Book",
		Fields: gql.Fields{
			"name":         &gql.Field{Type: gql.String},
			"abbreviation": &gql.Field{Type: gql.String},
		},
	}

	return gql.NewObject(objectConfig)
}

func bookField() *gql.Field {
	return &gql.Field{
		Type: &gql.List{OfType: bookObject()},
		Resolve: func(p gql.ResolveParams) (interface{}, error) {
			log.Println("Resolving book.name")
			bookRepo := &repositories.BookRepository{}
			return bookRepo.GetAll()
		},
	}
}
