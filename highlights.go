package main

import "time"

type Highlights struct {
	Count          int    `json:"count"`
	NextPageCursor int    `json:"nextPageCursor"`
	Results        []Book `json:"results"`
}

type Book struct {
	UserBookID    int             `json:"user_book_id"`
	Title         string          `json:"title"`
	Author        string          `json:"author"`
	ReadableTitle string          `json:"readable_title"`
	Source        string          `json:"source"`
	CoverImageURL string          `json:"cover_image_url"`
	UniqueURL     string          `json:"unique_url"`
	BookTags      []string        `json:"book_tags"`
	Category      string          `json:"category"`
	DocumentNote  interface{}     `json:"document_note"`
	ReadwiseURL   string          `json:"readwise_url"`
	SourceURL     string          `json:"source_url"`
	Asin          interface{}     `json:"asin"`
	Highlights    []BookHighlight `json:"highlights"`
}

type BookHighlight struct {
	ID            int         `json:"id"`
	Text          string      `json:"text"`
	Location      int         `json:"location"`
	LocationType  string      `json:"location_type"`
	Note          string      `json:"note"`
	Color         string      `json:"color"`
	HighlightedAt time.Time   `json:"highlighted_at"`
	CreatedAt     time.Time   `json:"created_at"`
	UpdatedAt     time.Time   `json:"updated_at"`
	ExternalID    string      `json:"external_id"`
	EndLocation   interface{} `json:"end_location"`
	URL           string      `json:"url"`
	BookID        int         `json:"book_id"`
	Tags          []string    `json:"tags"`
	IsFavorite    bool        `json:"is_favorite"`
	IsDiscard     bool        `json:"is_discard"`
	ReadwiseURL   string      `json:"readwise_url"`
}
