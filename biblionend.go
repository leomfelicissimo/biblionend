package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/graphql-go/handler"
	"github.com/leomfelicissimo/biblionend/schemas"
)

func main() {
	schema, err := schemas.CreateSchema()
	if err != nil {
		log.Fatalln("Erro ao criar schema", err)
	}

	h := handler.New(&handler.Config{
		Schema:     &schema,
		Pretty:     true,
		GraphiQL:   true,
		Playground: true,
	})

	http.Handle("/graphql", h)

	fmt.Println("Listening and serving at: http://localhost:5000/graphql")
	http.ListenAndServe(":5000", nil)
}
