package repository

import (
	"database/sql"
	"flashcards/model"
	"fmt"

	"github.com/google/uuid"
)

type CardRepository struct {
	connection *sql.DB
}

func NewCardRepository(db *sql.DB) *CardRepository {
	return &CardRepository{
		connection: db,
	}
}

func (cr *CardRepository) CreateCard(card model.Card) (uuid.UUID, error) {
	card.ID = uuid.New()

	query := `
		INSERT INTO cards (id, name, front, back, audio_url, image_url, deck_id, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
		RETURNING id
	`

	var newCardID uuid.UUID
	err := cr.connection.QueryRow(query, card.ID, card.Name, card.Front, card.Back, card.AudioURL, card.ImageURL, card.DeckID).Scan(&newCardID)
	if err != nil {
		fmt.Println("Error creating card:", err)
		return uuid.Nil, err
	}

	return newCardID, nil
}

func (cr *CardRepository) GetCards() ([]model.Card, error) {
	query := `
		SELECT id, name, front, back, audio_url, image_url, deck_id, created_at, updated_at
		FROM cards
	`

	rows, err := cr.connection.Query(query)
	if err != nil {
		fmt.Println("Error getting cards:", err)
		return []model.Card{}, err
	}

	var cards []model.Card

	for rows.Next() {
		var card model.Card
		err = rows.Scan(&card.ID, &card.Name, &card.Front, &card.Back, &card.AudioURL, &card.ImageURL, &card.DeckID, &card.CreatedAt, &card.UpdatedAt)
		if err != nil {
			fmt.Println("Error scanning card:", err)
			return []model.Card{}, err
		}

		cards = append(cards, card)
	}

	rows.Close()

	return cards, nil
}

func (cr *CardRepository) GetCardByID(id uuid.UUID) (model.Card, error) {
	query := `
		SELECT id, name, front, back, audio_url, image_url, deck_id, created_at, updated_at
		FROM cards
		WHERE id = $1
	`

	var card model.Card
	err := cr.connection.QueryRow(query, id).Scan(&card.ID, &card.Name, &card.Front, &card.Back, &card.AudioURL, &card.ImageURL, &card.DeckID, &card.CreatedAt, &card.UpdatedAt)
	if err != nil {
		fmt.Println("Error getting card by ID:", err)
		return model.Card{}, err
	}

	return card, nil
}

func (cr *CardRepository) UpdateCard(card model.Card) error {
	query := `
		UPDATE cards
		SET name = $1, front = $2, back = $3, audio_url = $4, image_url = $5, deck_id = $6, updated_at = CURRENT_TIMESTAMP
		WHERE id = $7
	`

	_, err := cr.connection.Exec(query, card.Name, card.Front, card.Back, card.AudioURL, card.ImageURL, card.DeckID, card.ID)
	if err != nil {
		fmt.Println("Error updating card:", err)
		return err
	}

	return nil
}

func (cr *CardRepository) DeleteCard(id uuid.UUID) error {
	query := `
		DELETE FROM cards
		WHERE id = $1
	`

	_, err := cr.connection.Exec(query, id)
	if err != nil {
		fmt.Println("Error deleting card:", err)
		return err
	}

	return nil
}
