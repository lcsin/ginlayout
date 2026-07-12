package main

import (
	"ginlayout/internal/model"

	"gorm.io/gen"
	// "gorm.io/driver/mysql"
	// "gorm.io/gorm"
)

// To run this generator, you need to first define your models or connect to a live database.
// This is a basic template to get started with gorm-gen.
func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./internal/repository/query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
	})

	// Example: Generate from database
	// dsn := "root:123456@tcp(127.0.0.1:3306)/ginlayout?charset=utf8mb4&parseTime=True&loc=Local"
	// db, _ := gorm.Open(mysql.Open(dsn))
	// g.UseDB(db)

	// Example: Generate from struct models
	g.ApplyBasic(model.User{})

	// Generate the code
	g.Execute()
}
