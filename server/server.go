// handles the server routes
package server

import (
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/musaubrian/backend-go/models"
)

//gets all the redirect urls
func getRedirects(c *fiber.Ctx) error{
    links, err := models.GetAllLinks()
    if err != nil {
        log.Fatal("Error getting all links:", err) 
    }

    return c.Status(fiber.StatusOK).JSON(links)
}

func getLink(c *fiber.Ctx) error  {
    id, err := strconv.ParseUint(c.Params("id"), 10, 64)
    if err != nil {
        log.Fatal("Error reading link id: ", err) 
    }

    link, err := models.GetLink(id)
    if err != nil {
        log.Fatal("Error getting link from db", err)
    }

    return c.Status(fiber.StatusOK).JSON(link)

}

func SetupAndListen()  {
    router := fiber.New()

    router.Use(cors.New(cors.Config{
        AllowOrigins: "*",
        AllowHeaders: "Origin, Content-Type, Accept",
    }))
    router.Get("/lns", getRedirects)
    router.Get("/l/:id", getLink)

    router.Listen(":8000")
}
