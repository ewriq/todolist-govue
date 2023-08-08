package main

import (
	"api/Database"

	"api/Utils"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Request struct {
	Title string `json:"Title"`
}

func main() {
	app := fiber.New()
	Database.Connect()

	app.Use(cors.New(cors.Config{
        AllowOrigins: "http://localhost:5173",
        AllowHeaders:  "Origin, Content-Type, Accept",
        AllowCredentials: true,
    }))
	
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	app.Get("/todos", func(c *fiber.Ctx) error {
		listelerim, err := Database.List()
		if err != nil {
			return c.Status(404).SendString("404")
		}
		return c.JSON(fiber.Map{
			"data": listelerim,
		})
	})

	app.Get("/del/:token", func(c *fiber.Ctx) error {
		token := c.Params("token")
		Database.Delete(token)
		return c.JSON(fiber.Map{
			"status": "OK",
			"msg":    "Başarılı bir şekilde silindi !",
			"token":  token,
		})
	})

	app.Get("/upt/:token", func(c *fiber.Ctx) error {
		token := c.Params("token")
		Database.Update(token)
		return c.JSON(fiber.Map{
			"status": "OK",
			"msg":    "Başarılı bir şekilde yapıldı !",
			"token":  token,
		})
	})

	app.Post("/add", func(c *fiber.Ctx) error {
		var request Request
		if err := c.BodyParser(&request); err != nil {
			fmt.Println("Error parsing request:", err)
			return err
		}

		title := request.Title
		GENtoken := Utils.Token(31)
		Database.Insert(GENtoken, title)

		fmt.Println("Salamben geldim 41")
		return c.JSON(fiber.Map{
			"token": GENtoken,
			"title": title,
			"msg":   "Başarıyla eklendi!",
		})
	})

	app.Listen(":3000")
}
