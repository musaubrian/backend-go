// handles the server routes
package server

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/musaubrian/backend-go/models"
)


//gets all the redirect urls
func getRedirects(ctx *fiber.Ctx)error  {
    links, err := models.GetAllLinks()
    if err != nil {
        log.Fatal("Error getting all links:", err)
        return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "message": "error getting all the links"+err.Error(),
        }))
    }

    return ctx.Status(fiber.StatusOK).JSON(links)
}

func SetupAndListen()  {
    router := fiber.New()

    router.Use(cors.New(cors.Config{
        AllowOrigin: "*",
        AllowHeaders: "Origin, Content-Type, Accept",
    }))
}
