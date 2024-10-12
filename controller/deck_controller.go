package controller

import (
	"flashcards/model"
	"flashcards/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// DeckController manages deck-related endpoints
type DeckController struct {
	deckService service.DeckService
}

// NewDeckController initializes the deck controller
func NewDeckController(decks service.DeckService) *DeckController {
	return &DeckController{
		deckService: decks,
	}
}

// CreateDeck godoc
// @Summary Create a new deck
// @Description Create a new deck with a name
// @Tags Decks
// @Accept  json
// @Produce  json
// @Param   deck  body  model.Deck  true  "Deck JSON"
// @Success 201 {object} model.Response
// @Failure 400 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /decks [post]
func (dc *DeckController) CreateDeck(ctx *gin.Context) {
	var deck model.Deck
	err := ctx.BindJSON(&deck)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
		return
	}

	id, err := dc.deckService.CreateDeck(deck)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, model.Response{Message: id.String()})
}

// GetDecks godoc
// @Summary Get all decks
// @Description Retrieve all decks
// @Tags Decks
// @Produce  json
// @Success 200 {array} model.Deck
// @Failure 500 {object} model.ErrorResponse
// @Router /decks [get]
func (dc *DeckController) GetDecks(ctx *gin.Context) {
	decks, err := dc.deckService.GetDecks()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, decks)
}

// GetDeckByID godoc
// @Summary Get a deck by ID
// @Description Retrieve a deck by its UUID
// @Tags Decks
// @Produce  json
// @Param id path string true "Deck UUID"
// @Success 200 {object} model.Deck
// @Failure 400 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /decks/{id} [get]
func (dc *DeckController) GetDeckByID(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Invalid UUID"})
		return
	}

	deck, err := dc.deckService.GetDeckByID(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, deck)
}

// UpdateDeck godoc
// @Summary Update a deck by ID
// @Description Update a deck by its UUID
// @Tags Decks
// @Accept  json
// @Produce  json
// @Param id path string true "Deck UUID"
// @Param deck body model.Deck true "Updated Deck JSON"
// @Success 200 {object} model.Response
// @Failure 400 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /decks/{id} [put]
func (dc *DeckController) UpdateDeck(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Invalid UUID"})
		return
	}

	var deck model.Deck
	err = ctx.BindJSON(&deck)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
		return
	}

	deck.ID = id

	err = dc.deckService.UpdateDeck(deck)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, model.Response{Message: "Deck updated successfully"})
}

// DeleteDeck godoc
// @Summary Delete a deck by ID
// @Description Delete a deck by its UUID
// @Tags Decks
// @Produce  json
// @Param id path string true "Deck UUID"
// @Success 200 {object} model.Response
// @Failure 400 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /decks/{id} [delete]
func (dc *DeckController) DeleteDeck(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Invalid UUID"})
		return
	}

	err = dc.deckService.DeleteDeck(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, model.Response{Message: "Deck deleted successfully"})
}
