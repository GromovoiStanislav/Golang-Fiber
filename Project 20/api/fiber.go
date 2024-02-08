package api
 
import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"fiber-example/database"
)
 
var (
   articles = map[string]database.Article{}
)

func createArticle(c *fiber.Ctx) error {
   article := new(database.Article)
 
   err := c.BodyParser(&article)
   if err != nil {
       return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
           "errors": err.Error(),
       })
 
   }
 
   article.ID = uuid.New().String()
  
   articles[article.ID] = *article
 
   c.Status(200).JSON(&fiber.Map{
       "article": article,
   })
 
   return nil
}


func readArticle(c *fiber.Ctx) error {
	id := c.Params("id")
  
	if article, ok := articles[id]; ok {
		c.Status(200).JSON(&fiber.Map{
			"article": article,
		})
	} else {
		c.Status(404).JSON(&fiber.Map{
			"error": "article not found",
		})
	}
	return nil
 }


 func readArticles(c *fiber.Ctx) error {
	c.Status(200).JSON(&fiber.Map{
		"articles": articles,
	})
	return nil
 }


 func updateArticle(c *fiber.Ctx) error {
	updateArticle := new(database.Article)
  
	err := c.BodyParser(updateArticle)
  
	if err != nil {
		c.Status(500).JSON(&fiber.Map{
			"error": err.Error(),
		})
		return err
	}
	id := c.Params("id")
	if article, ok := articles[id]; ok {
		article.Title = updateArticle.Title
		article.Description = updateArticle.Description
		article.Rate = updateArticle.Rate
		articles[id] = article
		c.Status(200).JSON(&fiber.Map{
			"article": article,
		})
	} else {
		c.Status(404).JSON(&fiber.Map{
			"error": "article not found",
		})
  
	}
	return nil
 }


 func deleteArticle(c *fiber.Ctx) error {
	id := c.Params("id")
  
	if _, ok := articles[id]; ok {
		delete(articles, id)
		c.Status(200).JSON(&fiber.Map{
			"message": "article deleted successfully",
		})
	} else {
		c.Status(404).JSON(&fiber.Map{
			"error": "article not found",
		})
	}
  
	return nil
 }

 func SetupRoute() *fiber.App {
	app := *fiber.New()
	app.Post("/api/v1/articles", createArticle)
	app.Get("/api/v1/articles/:id", readArticle)
	app.Get("/api/v1/articles/", readArticles)
	app.Put("/api/v1/articles/:id", updateArticle)
	app.Delete("/api/v1/articles/:id", deleteArticle)
	return &app
 }