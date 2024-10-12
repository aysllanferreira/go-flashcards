package controller

import (
	"flashcards/model"
	"flashcards/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CardController manages card-related endpoints
type CardController struct {
	cardService service.CardService
}

// NewCardController initializes the card controller
func NewCardController(cards service.CardService) *CardController {
	return &CardController{
		cardService: cards,
	}
}

// CreateCard godoc
// @Summary Create a new card
// @Description Create a new card with name, front, back, and optional audio/image URLs
// @Tags Cards
// @Accept  json
// @Produce  json
// @Param   card  body  model.Card  true  "Card JSON"
// @Success 201 {object} model.Response
// @Failure 400 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /cards [post]
func (cc *CardController) CreateCard(ctx *gin.Context) {
	var card model.Card
	err := ctx.BindJSON(&card)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
		return
	}

	id, err := cc.cardService.CreateCard(card)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, model.Response{Message: id.String()})
}

// GetCards godoc
// @Summary Get all cards
// @Description Retrieve all cards
// @Tags Cards
// @Produce  json
// @Success 200 {array} model.Card
// @Failure 500 {object} model.ErrorResponse
// @Router /cards [get]
func (cc *CardController) GetCards(ctx *gin.Context) {
	cards, err := cc.cardService.GetCards()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, cards)
}

// GetCardByID godoc
// @Summary Get a card by ID
// @Description Retrieve a card by its UUID
// @Tags Cards
// @Produce  json
// @Param id path string true "Card UUID"
// @Success 200 {object} model.Card
// @Failure 400 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /cards/{id} [get]
func (cc *CardController) GetCardByID(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Invalid UUID"})
		return
	}

	card, err := cc.cardService.GetCardByID(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, card)
}

// UpdateCard godoc
// @Summary Update a card
// @Description Update a card by its UUID
// @Tags Cards
// @Accept  json
// @Produce  json
// @Param card body model.Card true "Updated Card JSON"
// @Success 200 {object} model.Response
// @Failure 400 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /cards [put]
func (cc *CardController) UpdateCard(ctx *gin.Context) {
	var card model.Card
	err := ctx.BindJSON(&card)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
		return
	}

	err = cc.cardService.UpdateCard(card)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, model.Response{Message: "Card updated"})
}

// DeleteCard godoc
// @Summary Delete a card by ID
// @Description Delete a card by its UUID
// @Tags Cards
// @Produce  json
// @Param id path string true "Card UUID"
// @Success 200 {object} model.Response
// @Failure 400 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /cards/{id} [delete]
func (cc *CardController) DeleteCard(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Invalid UUID"})
		return
	}

	err = cc.cardService.DeleteCard(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, model.Response{Message: "Card deleted"})
}
