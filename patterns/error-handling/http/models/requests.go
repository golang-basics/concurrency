package models

// CreateNoteRequest represents the request for creating a note
type CreateNoteRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

// GetNoteRequest represents the request for fetching a note
type GetNoteRequest struct {
	ID string `json:"id"`
}
