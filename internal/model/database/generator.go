package database

import (
	"github.com/laravel2004/seracan-auth-service/internal/config"
	"gorm.io/gen"
	"log"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath:      "internal/model/database",
		OutFile:      "gen.go",
		Mode:         gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
		ModelPkgPath: "database",
	})

	g.UseDB(config.GetDB())

	g.GenerateAllTable()

	g.Execute()
	log.Printf("Database model generation complete!")
}
