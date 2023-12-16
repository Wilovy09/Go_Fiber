package main

import (
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/cors"
    "github.com/gofiber/fiber/v2/middleware/logger"
    "github.com/google/uuid"
)

type User struct {
    ID        uuid.UUID `json:"id"`
    Firstname string    `json:"firstname"`
    Lastname  string    `json:"lastname"`
}

func handleUser(c *fiber.Ctx) error {
    ID := uuid.New()
    user := User{
        ID:        ID,
        Firstname: "John",
        Lastname:  "Doe",
    }
    return c.Status(fiber.StatusOK).JSON(user)
}

func handleCreateUser(c *fiber.Ctx) error {
    user := User{}
    user.ID = uuid.New()
    if err := c.BodyParser(&user); err != nil {
        return err
    }
    return c.Status(fiber.StatusOK).JSON(user)
}

func main() {
    // Creamos una nueva instancia de Fiber
    app := fiber.New()

    // Middlewares
    app.Use(logger.New())

    app.Use(cors.New())
    app.Use(cors.New(cors.Config{
        // los AllowOrigins son separados por ','
        AllowOrigins: "https://wilovy.com, https:/gofiber.io",
        AllowMethods: "GET,POST,HEAD,PUT,PATCH,DELETE",
    }))

    // Ruta principal
    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, World!")
    })

    // Agregamos prefijo /api a todas las rutas que queramos
    apiGroup := app.Group("/api")

    // /api/user
    apiGroup.Get("/user", handleUser)

    // /api/users
    apiGroup.Post("/users", handleCreateUser)

    app.Listen(":8000")
}
