# Primer contacto con Fiber, backend para Go

Este es un primer contacto con Fiber, no es un CRUD

Solo es para ir conociendo las bases para luego hacer un CRUD

Primero, tenemos que iniciar nuestro proyecto en Go

```sh
go mod init github.com/USER/PROYECTO
```

Luego creamos nuestro archivo principal `main.go`

```go
package main

import "fmt"

func main(){
    fmt.Println("Hola Fiber!")
}
```

Ahora ya podemos instalar `Fiber`!

```sh
go get github.com/gofiber/fiber/v2
```

Tambien necesitaremos `UUID`

```sh
go get github.com/google/uuid
```

Ahora si en nuestro archivo main, empezaremos a construir nuestro esqueleto

```go
package main

// Importamos lo que instalamos
import (
    "github.com/gofiber/fiber/v2"
    "github.com/google/uuid"
)

// Creamos una Struct para nuestro usuario
type User struct {
    ID        uuid.UUID `json:"id"`
    Firstname string    `json:"firstname"`
    Lastname  string    `json:"lastname"`
}

// Creamos una funcion handleUser
func handleUser(c *fiber.Ctx) error {
    ID := uuid.New()
    user := User{
        ID:        ID,
        Firstname: "John",
        Lastname:  "Doe",
    }
    return c.Status(fiber.StatusOK).JSON(user)
}

// Creamos una funcion para crear usuarios
func handleCreateUser(c *fiber.Ctx) error {
    user := User{}
    user.ID = uuid.New()
    if err := c.BodyParser(&user); err != nil {
        return err
    }
    return c.Status(fiber.StatusOK).JSON(user)
}

// Funcion main
func main() {
    // Creamos una nueva instancia de Fiber
    app := fiber.New()

    // Ruta principal
    //               \/ Aqui tenemos que poner el Ctx 
    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, World!")
    })

    // Agregamos prefijo /api a todas las rutas que queramos
    apiGroup := app.Group("/api")
    // /api/user
    apiGroup.Get("/user", handleUser)
    // /api/users
    apiGroup.Post("/users", handleCreateUser)

    // El puerto en el que queremos lanzar el servidor
    app.Listen(":8000")
}
```

## Middlewares

```go
package main

import (
    "github.com/gofiber/fiber/v2"
    // Importamos el middleware/logger
    // Este middleware nos hare un log de todas las peticiones que se hagan al servidor, el log se mostrara en consola
    "github.com/gofiber/fiber/v2/middleware/logger"
    "github.com/google/uuid"
)

// ...

func main() {
    app := fiber.New()

    // MIDDLEWARES
    // Los middlewares afectan a las ursl que esten por debajo de ellos
    // En este caso queremos que afecte a todas entonces lo pondremos lo mas arriba posible
    app.Use(logger.New())

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, World!")
    })

    //...
}

```

## CORS

Recomiendo leer la web oficial para entender mejor este concepto.

```go
package main

import (
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/cors"
    "github.com/gofiber/fiber/v2/middleware/logger"
    "github.com/google/uuid"
)

// ...

func main() {
    app := fiber.New()

    // MIDDLEWARES
    // Los middlewares afectan a las ursl que esten por debajo de ellos
    // En este caso queremos que afecte a todas entonces lo pondremos lo mas arriba posible
    app.Use(logger.New())
    
    app.Use(cors.New())
    app.Use(cors.New(cors.Config{
        // los AllowOrigins son separados por ','
        AllowOrigins: "https://wilovy.com, https:/gofiber.io",
        AllowMethods: "GET,POST,HEAD,PUT,PATCH,DELETE",
    }))

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, World!")
    })

    //...
}
```

Asi debio finalizar nuestro primer contacto con Fiber

```go
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
```
