package service

import (
	"flashcards/model"
	"flashcards/repository"

	"github.com/google/uuid"
)

type DeckService struct {
	repository *repository.DeckRepository
}

func NewDeckService(repo *repository.DeckRepository) DeckService {
	return DeckService{
		repository: repo,
	}
}

func (ds *DeckService) CreateDeck(deck model.Deck) (uuid.UUID, error) {
	return ds.repository.CreateDeck(deck)
}

func (ds *DeckService) GetDecks() ([]model.Deck, error) {
	return ds.repository.GetDecks()
}

func (ds *DeckService) GetDeckByID(id uuid.UUID) (model.Deck, error) {
	return ds.repository.GetDeckByID(id)
}

func (ds *DeckService) UpdateDeck(deck model.Deck) error {
	return ds.repository.UpdateDeck(deck)
}

func (ds *DeckService) DeleteDeck(id uuid.UUID) error {
	return ds.repository.DeleteDeck(id)
}
