package controller

import (
	"flashcards/model"
	"flashcards/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CardController struct {
	cardService service.CardService
}

func NewCardController(cards service.CardService) *CardController {
	return &CardController{
		cardService: cards,
	}
}

func (cc *CardController) CreateCard(ctx *gin.Context) {
	var card model.Card
	err := ctx.BindJSON(&card)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := cc.cardService.CreateCard(card)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"id": id})
}

func (cc *CardController) GetCards(ctx *gin.Context) {
	cards, err := cc.cardService.GetCards()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, cards)
}

func (cc *CardController) GetCardByID(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := uuid.Parse(idStr)
	card, err := cc.cardService.GetCardByID(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, card)
}

func (cc *CardController) UpdateCard(ctx *gin.Context) {
	var card model.Card
	err := ctx.BindJSON(&card)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = cc.cardService.UpdateCard(card)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Card updated"})
}

func (cc *CardController) DeleteCard(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
		return
	}

	err = cc.cardService.DeleteCard(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Card deleted"})
}
