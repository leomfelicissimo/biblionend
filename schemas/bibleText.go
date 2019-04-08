package schemas

import (
	"log"
	"regexp"

	"github.com/leomfelicissimo/biblionend/repositories"

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

func sanitize(ref string) string {
	r := regexp.MustCompile(`([\.:,;\s])`)
	return r.ReplaceAllString(ref, "")
}

func bibleTextField() *gql.Field {
	return &gql.Field{
		Type: bibleTextObject(),
		Resolve: func(p gql.ResolveParams) (interface{}, error) {
			reference := sanitize(p.Args["reference"].(string))
			// TODO: Identificar a abreviatura do livro
			log.Println("Reference:", reference)

			repository := &repositories.BibleTextRepository{}
			return repository.FindByReference(reference)
		},
		Args: gql.FieldConfigArgument{
			"reference": &gql.ArgumentConfig{Type: gql.String},
		},
	}
}
