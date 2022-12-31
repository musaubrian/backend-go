// handles the server routes
package server

import (
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/musaubrian/backend-go/models"
    "github.com/musaubrian/backend-go/utils"
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


func createLink(c *fiber.Ctx) error {
    c.Accepts("application/json")

    var link models.TinyUrl
    err := c.BodyParser(&link)
    if err != nil {
        log.Fatal("Error parsing JSON: ", err)
    }

    link.ShortUrl = utils.GenerateUrl()

    err = models.CreateLink(link)
    if err != nil {
        log.Fatal("Could not create new link")
    }

    return c.Status(fiber.StatusOK).JSON(link)
}


func SetupAndListen()  {
    router := fiber.New()

    router.Use(cors.New(cors.Config{
        AllowOrigins: "*",
        AllowHeaders: "Origin, Content-Type, Accept",
    }))
    router.Get("/l", getRedirects)
    router.Get("/l/:id", getLink)
    router.Post("/n", createLink)

    router.Listen(":8000")
}
