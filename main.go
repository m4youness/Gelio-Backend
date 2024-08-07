package main

import (
	services "gelio/m/Services"
	"gelio/m/controllers"
	"gelio/m/initializers"

	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
)

func init() {
	initializers.CloudinaryConnect()
	initializers.DbConnect()
	initializers.InitRedis()
}

func main() {

	r := gin.Default()
	c := cors.New(cors.Options{
		// AllowedOrigins: []string{"http://localhost:4200"},
		AllowedOrigins:   []string{"https://gelio.vercel.app"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "OPTIONS"},
		AllowCredentials: true,
	})
	r.StaticFile("/test", "./test.html")

	r.Use(corsHandler(c))

	// User
	userController := controllers.NewUserController(services.UserService{}, services.PersonService{})
	userController.InitializeRoutes(r)

	// People
	peopleController := controllers.NewPersonController(services.PersonService{})
	peopleController.InitializeRoutes(r)

	// Country
	countryController := controllers.CountryController()
	countryController.InitializeRoutes(r)

	// Image
	imageController := controllers.ImageController()
	imageController.InitializeRoutes(r)

	// Message
	messageController := controllers.MessageController(services.UserService{})
	messageController.InitializeRoutes(r)

	// Post
	postController := controllers.PostController()
	postController.InitializeRoutes(r)

	// Post Likes
	postLikesController := controllers.PostLikesController()
	postLikesController.InitializeRoutes(r)

	// Comments
	commentsController := controllers.CommentsController()
	commentsController.InitializeRoutes(r)

	// Websockets
	websocketController := controllers.NewServer()
	websocketController.InitializeRoutes(r)

	r.Run()

}

func corsHandler(corsMiddleware *cors.Cors) gin.HandlerFunc {
	return func(c *gin.Context) {
		corsMiddleware.HandlerFunc(c.Writer, c.Request)
		c.Next()
	}
}
