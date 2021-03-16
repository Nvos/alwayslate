// +build ignore

package main

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"log"
)


func main() {
	templates := entgql.AllTemplates
	opts := []entc.Option{
		entc.FeatureNames("privacy"),
		entc.FeatureNames("entql"),
	}

	err := entc.Generate("./schema", &gen.Config{
		Templates: templates,
	}, opts...)

	if err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}