package service

import (
	"flashcards/model"
	"flashcards/repository"

	"github.com/google/uuid"
)

type CardService struct {
	repository *repository.CardRepository
}

func NewCardService(repo *repository.CardRepository) CardService {
	return CardService{
		repository: repo,
	}
}

func (cs *CardService) CreateCard(card model.Card) (uuid.UUID, error) {
	return cs.repository.CreateCard(card)
}

func (cs *CardService) GetCards() ([]model.Card, error) {
	return cs.repository.GetCards()
}

func (cs *CardService) GetCardByID(id uuid.UUID) (model.Card, error) {
	return cs.repository.GetCardByID(id)
}

func (cs *CardService) UpdateCard(card model.Card) error {
	return cs.repository.UpdateCard(card)
}

func (cs *CardService) DeleteCard(id uuid.UUID) error {
	return cs.repository.DeleteCard(id)
}
