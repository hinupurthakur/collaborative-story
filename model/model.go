package model

import "time"

type WordDTO struct {
	Word string `db:"word" json:"word" validate:"required,excludesall=' '"`
}

type NewWordResponse struct {
	ID              int    `db:"id" json:"id"`
	Title           string `db:"title" json:"title"`
	CurrentSentence string `db:"-" json:"current_sentence`
}

type StoryDTO struct {
	ID         int        `db:"id" json:"id"`
	Title      string     `db:"title" json:"title"`
	CreatedAt  time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt  time.Time  `db:"updated_at" json:"updated_at"`
	IsDeleted  bool       `db:"is_deleted" json:"is_deleted,omitempty"`
	Paragraphs []Sentence `db:"paragraphs" json:"paragraphs,omitempty"`
}

type ParagraphDTO struct {
	ID        int       `db:"id" json:"id"`
	StoryID   int       `db:"story_id" json:"story_id"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
	IsDeleted bool      `db:"is_deleted" json:"is_deleted,omitempty"`
	Sentences []string  `db:"sentences" json:"sentences,omitempty"`
}

type Sentence struct {
	Sentences []string `json:"sentences"`
}

type SentenceDTO struct {
	ID          int       `db:"id" json:"id"`
	Sentence    string    `db:"sentence" json:"sentence"`
	ParagraphID int       `db:"paragraph_id" json:"paragraph_id"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
	IsDeleted   bool      `db:"is_deleted" json:"is_deleted,omitempty"`
}

type Stories struct {
	Limit   uint32     `json:"limit"`
	Offset  uint32     `json:"offset"`
	Count   uint32     `json:"count"`
	Results []StoryDTO `json:"results"`
}
