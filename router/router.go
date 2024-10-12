package router

import (
	"flashcards/controller"
	"flashcards/repository"
	"flashcards/service"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"database/sql"

	"github.com/gin-gonic/gin"

	_ "flashcards/docs"
)

func SetupRouter(dbConnection *sql.DB) *gin.Engine {
	router := gin.Default()

	// Swagger route
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Initialize repositories
	deckRepository := repository.NewDeckRepository(dbConnection)
	cardRepository := repository.NewCardRepository(dbConnection)

	// Initialize services
	deckService := service.NewDeckService(deckRepository)
	cardService := service.NewCardService(cardRepository)

	// Initialize controllers
	deckController := controller.NewDeckController(deckService)
	cardController := controller.NewCardController(cardService)

	// Deck routes
	deckRoutes := router.Group("/decks")
	{
		deckRoutes.POST("/", deckController.CreateDeck)
		deckRoutes.GET("/", deckController.GetDecks)
		deckRoutes.GET("/:id", deckController.GetDeckByID)
		deckRoutes.PUT("/:id", deckController.UpdateDeck)
		deckRoutes.DELETE("/:id", deckController.DeleteDeck)
	}

	// Card routes
	cardRoutes := router.Group("/cards")
	{
		cardRoutes.POST("/", cardController.CreateCard)
		cardRoutes.GET("/", cardController.GetCards)
		cardRoutes.GET("/:id", cardController.GetCardByID)
		cardRoutes.PUT("/:id", cardController.UpdateCard)
		cardRoutes.DELETE("/:id", cardController.DeleteCard)
	}

	return router
}
