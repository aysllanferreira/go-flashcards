package controller

import (
	"flashcards/model"
	"flashcards/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type DeckController struct {
	deckService service.DeckService
}

func NewDeckController(decks service.DeckService) *DeckController {
	return &DeckController{
		deckService: decks,
	}
}

func (dc *DeckController) CreateDeck(ctx *gin.Context) {
	var deck model.Deck
	err := ctx.BindJSON(&deck)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := dc.deckService.CreateDeck(deck)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"id": id})
}

func (dc *DeckController) GetDecks(ctx *gin.Context) {
	decks, err := dc.deckService.GetDecks()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, decks)
}

func (dc *DeckController) GetDeckByID(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := uuid.Parse(idStr)
	deck, err := dc.deckService.GetDeckByID(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, deck)
}

func (dc *DeckController) UpdateDeck(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
		return
	}

	var deck model.Deck
	err = ctx.BindJSON(&deck)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	deck.ID = id

	err = dc.deckService.UpdateDeck(deck)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Deck updated successfully"})
}

func (dc *DeckController) DeleteDeck(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
		return
	}

	err = dc.deckService.DeleteDeck(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Deck deleted successfully"})
}
