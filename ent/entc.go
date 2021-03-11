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

	templates = append(templates, gen.MustParse(
		gen.NewTemplate("pulid.tmpl").ParseFiles("./schema/pulid/template/pulid.tmpl")),
	)

	err := entc.Generate("./schema", &gen.Config{
		Templates: templates,
	})

	if err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}