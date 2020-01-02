package main

import(
	"github.com/abenbyy/graphql-api/api"
	"github.com/abenbyy/graphql-api/graphql/mutations"
	"github.com/abenbyy/graphql-api/graphql/query"
	"github.com/abenbyy/graphql-api/middleware"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"log"
	"net/http"
)

func main(){
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: query.GetRoot(),
		Mutation: mutations.GetRoot(),
	})

	if err != nil{
		panic(err.Error())
	}

	h := handler.New(&handler.Config{
		Schema: &schema,
		Pretty:	true,
		GraphiQL: true,
		Playground: true,
	})

	wrapped := middleware.CorsMiddleware(h)

	router := api.NewRouter()
	router.Handle("/api",wrapped)
	log.Fatalln(http.ListenAndServe(":2000",router))
}
