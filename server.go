package main

import (
	"alwayslate/ent"
	"alwayslate/ent/migrate"
	_ "alwayslate/ent/runtime"
	"alwayslate/graph"
	"alwayslate/graph/generated"
	"context"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}

	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background(), migrate.WithGlobalUniqueID(true)); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
		Client: client,
	}}))

	project, _ := client.Project.Create().SetName("project 1").Save(context.Background())
	activity1, _ := client.Activity.Create().SetName("activity 1").SetProject(project).Save(context.Background())
	activity2, _ := client.Activity.Create().SetName("activity 2").SetProject(project).Save(context.Background())

	client.Timesheet.Create().SetActivity(activity1).SetDuration(60).Save(context.Background())
	client.Timesheet.Create().SetActivity(activity1).SetDuration(160).Save(context.Background())
	client.Timesheet.Create().SetActivity(activity1).SetDuration(360).Save(context.Background())
	client.Timesheet.Create().SetActivity(activity2).SetDuration(70).Save(context.Background())

	log.Printf("project %v, activity: %v, %v", project, activity1, activity2)

	found, _ := client.Activity.Get(context.Background(), activity1.ID)
	log.Printf("project: %v", found.Edges.Project)

	router := chi.NewRouter()
	// project: 4294967297
	// activity: 1, 2, 3, 4, 5
	// timesheet: 8589934593, 8589934594, 8589934595, 8589934596
	// Add CORS middleware around every request
	// See https://github.com/rs/cors for full option listing
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080", "http://localhost:3000", "http://localhost:5000"},
		AllowCredentials: true,
		//Debug:            true,
	}).Handler)

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
