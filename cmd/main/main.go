package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/iamganeshagrawal/go-gql-upload-demo/graph/generated"
	"github.com/iamganeshagrawal/go-gql-upload-demo/graph/resolver"
)

const defaultPort = "8080"
const mb int64 = 1 << 20

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	gqlResolvers := &resolver.Resolver{}
	gqlConfig := generated.Config{
		Resolvers: gqlResolvers,
	}
	gqlExeSchema := generated.NewExecutableSchema(gqlConfig)

	srv := handler.NewDefaultServer(gqlExeSchema)
	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.MultipartForm{
		MaxUploadSize: 32 * mb,
		MaxMemory:     64 * mb,
	})
	srv.Use(extension.Introspection{})

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)
	http.Handle("/upload/", http.StripPrefix("/upload/", http.FileServer(http.Dir("./uploads"))))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
