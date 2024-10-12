package repository

import (
	"database/sql"
	"flashcards/model"
	"fmt"

	"github.com/google/uuid"
)

type DeckRepository struct {
	connection *sql.DB
}

func NewDeckRepository(db *sql.DB) *DeckRepository {
	return &DeckRepository{
		connection: db,
	}
}

func (dr *DeckRepository) CreateDeck(deck model.Deck) (uuid.UUID, error) {
	query := "INSERT INTO decks (name) VALUES ($1) RETURNING id"
	var id uuid.UUID
	err := dr.connection.QueryRow(query, deck.Name).Scan(&id)
	if err != nil {
		fmt.Println(err)
		return uuid.Nil, err
	}

	return id, nil
}

func (dr *DeckRepository) GetDecks() ([]model.Deck, error) {
	query := "SELECT id, name, created_at, updated_at FROM decks"
	rows, err := dr.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return []model.Deck{}, err
	}

	var decks []model.Deck
	var deckObj model.Deck

	for rows.Next() {
		err = rows.Scan(&deckObj.ID, &deckObj.Name, &deckObj.CreatedAt, &deckObj.UpdatedAt)

		if err != nil {
			fmt.Println(err)
			return []model.Deck{}, err
		}

		decks = append(decks, deckObj)
	}

	rows.Close()

	return decks, nil
}

func (dr *DeckRepository) GetDeckByID(id uuid.UUID) (model.Deck, error) {
	query := "SELECT id, name, created_at, updated_at FROM decks WHERE id = $1"
	var deck model.Deck
	err := dr.connection.QueryRow(query, id).Scan(&deck.ID, &deck.Name, &deck.CreatedAt, &deck.UpdatedAt)
	if err != nil {
		fmt.Println(err)
		return model.Deck{}, err
	}

	return deck, nil
}

func (dr *DeckRepository) UpdateDeck(deck model.Deck) error {
	query := "UPDATE decks SET name = $1, updated_at = CURRENT_TIMESTAMP WHERE id = $2"
	result, err := dr.connection.Exec(query, deck.Name, deck.ID)
	if err != nil {
		fmt.Println(err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no rows were updated, deck with id %s not found", deck.ID)
	}

	return nil
}

func (dr *DeckRepository) DeleteDeck(id uuid.UUID) error {
	query := "DELETE FROM decks WHERE id = $1"
	_, err := dr.connection.Exec(query, id)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
