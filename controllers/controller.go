package controllers

import (
	"fmt"
	database "go-fiber-test/databases"
	models "go-fiber-test/models"
	"log"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func AddUser(c *fiber.Ctx) error {
	validate := validator.New()
	//Connect to database
	user := new(models.User)
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	errors := validate.Struct(user)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors.Error())
	}
	return c.JSON(user)
}

func HelloWorld(ctx *fiber.Ctx) error {
	// Send custom error page
	err := ctx.Status(500).SendFile(fmt.Sprintf("./%d.html", 500))
	if err != nil {
		// In case the SendFile fails
		return ctx.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}
	// Return from handler
	return nil
}

/*
func Error(ctx *fiber.Ctx, err error) error {
	// Status code defaults to 500
	code := fiber.StatusInternalServerError

	// Retrieve the custom status code if it's a *fiber.Error
	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}

	// Send custom error page
	err = ctx.Status(code).SendFile(fmt.Sprintf("./%d.html", code))
	if err != nil {
		// In case the SendFile fails
		return ctx.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	// Return from handler
	return nil
}
*/

func Name(c *fiber.Ctx) error {
	p := new(models.Person)
	if err := c.BodyParser(p); err != nil {
		return err
	}
	log.Println(p.Name)
	log.Println(p.Pass)
	return c.JSON(p)
}

func GetUser(c *fiber.Ctx) error {
	a := c.Params("name")
	str := "my name is " + a
	return c.SendString(str)
}

func Fact(c *fiber.Ctx) error {
	a := c.Params("n")
	n, _ := strconv.Atoi(a)
	fact := 1
	if n < 0 {
		str := "Worng Input"
		return c.SendString(str)
	} else {
		for i := 1; i <= n; i++ {
			fact *= i
		}
	}

	//break 31 - 32
	result := strconv.Itoa(fact)
	return c.SendString(result)
}

func Search(c *fiber.Ctx) error {
	c.Query("search") // "fenny"

	a := c.Query("search")
	str := "my search is  " + a
	return c.JSON(str)
}

func Params(c *fiber.Ctx) error {
	fmt.Fprintf(c, "%s\n", c.Params("name"))
	fmt.Fprintf(c, "%s", c.Params("title"))
	return nil
}

func GetList(c *fiber.Ctx) error {
	return c.SendString("string: api/v1/list")
}

func V1GetUser(c *fiber.Ctx) error {
	return c.SendString("string: api/v1/user")
}

func V2List(c *fiber.Ctx) error {
	return c.SendString("string: api/v2/list")
}

func V2User(c *fiber.Ctx) error {
	return c.SendString("string: api/v2/user")
}

func GetDogs(c *fiber.Ctx) error {
	db := database.DBConn
	var dogs []models.Dogs

	db.Find(&dogs)
	return c.Status(200).JSON(dogs)
}

func GetDog(c *fiber.Ctx) error {
	db := database.DBConn
	search := strings.TrimSpace(c.Query("search"))
	var dog []models.Dogs

	result := db.Find(&dog, "dog_id = ?", search)

	// returns found records count, equals `len(users)
	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}
	return c.Status(200).JSON(&dog)
}

func AddDog(c *fiber.Ctx) error {
	//twst3
	db := database.DBConn
	var dog models.Dogs

	if err := c.BodyParser(&dog); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.Create(&dog)
	return c.Status(201).JSON(dog)
}

func UpdateDog(c *fiber.Ctx) error {
	db := database.DBConn
	var dog models.Dogs
	id := c.Params("id")

	if err := c.BodyParser(&dog); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.Where("id = ?", id).Updates(&dog)
	return c.Status(200).JSON(dog)
}

func RemoveDog(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")
	var dog models.Dogs

	result := db.Delete(&dog, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.SendStatus(200)
}
