package main

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"

	"github.com/graphql-go/graphql"
	"github.com/leomfelicissimo/biblionend/schemas"
)

func TestMain(t *testing.T) {
	schema, err := schemas.CreateSchema()
	if err != nil {
		log.Fatalln("Erro ao criar schema", err)
	}

	query := `
        {
            book {
                name
            }
        }
    `

	params := graphql.Params{Schema: schema, RequestString: query}
	r := graphql.Do(params)
	if len(r.Errors) > 0 {
		log.Fatalf("Failed to execute graphql operation, errors: %+v", r.Errors)
	}

	rJSON, _ := json.Marshal(r)
	fmt.Printf("%s \n", rJSON)
}
