package main

import (
	"fmt"
	"myapp/config"
	"myapp/routes"

	"gorm.io/gen"
)

func main() {
	db, err := config.InitDatabase()
	if err != nil {
		fmt.Print(err)
	}
	g := gen.NewGenerator(gen.Config{
		OutPath: "myapp",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
	})
	g.UseDB(db)
	// g.GenerateModel("roles")
	// g.Execute()
	routes.InitRoute()

}
