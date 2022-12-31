// handles the server routes
package server

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/musaubrian/backend-go/models"
)

//gets all the redirect urls
func getRedirects(c *fiber.Ctx) error{
    links, err := models.GetAllLinks()
    if err != nil {
        log.Fatal("Error getting all links:", err)
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "message": "error getting all the links" + err.Error(),
        })
    }

    return c.Status(fiber.StatusOK).JSON(links)
}

func SetupAndListen()  {
    router := fiber.New()

    router.Use(cors.New(cors.Config{
        AllowOrigins: "*",
        AllowHeaders: "Origin, Content-Type, Accept",
    }))
    router.Get("/lnks", getRedirects)

    router.Listen(":8000")
}
