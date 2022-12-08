package main

import (
	"fmt"
	database "go-fiber-test/databases"

	m "go-fiber-test/models"
	"go-fiber-test/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func initDataBase() {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		"root",
		"",
		"127.0.0.1",
		"3306",
		"golang_test",
	)
	var err error
	database.DBConn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	// log.Println(db)
	fmt.Println("Database Connect Successfully")
	database.DBConn.AutoMigrate(&m.Dogs{})

}

func main() {
	app := fiber.New()
	// Provide a minimal config
	app.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{
			// "john":  "doe",
			// "admin": "123456",
			"gofiber": "15112565",
		},
	}))

	routes.Route(app)

	initDataBase()

	app.Listen(":3000")

}
