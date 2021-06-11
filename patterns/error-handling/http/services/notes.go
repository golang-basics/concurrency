package services

import (
	"context"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"

	"github.com/steevehook/http/logging"
	"github.com/steevehook/http/models"
)

type repo interface {
	GetNote(ctx context.Context, id string) (models.Note, error)
	CreateNote(ctx context.Context, note models.Note) error
}

type NotesService struct {
	repo
}

func NewNotes(r repo) NotesService {
	return NotesService{
		repo: r,
	}
}

func (r NotesService) CreateNote(ctx context.Context, req models.CreateNoteRequest) (models.Note, error) {
	logger := logging.Logger()
	id := uuid.New()
	note := models.Note{
		ID:          id.String(),
		Title:       req.Title,
		Description: req.Description,
		CreatedAt:   time.Now().UTC(),
	}

	err := r.repo.CreateNote(ctx, note)
	if err != nil {
		logger.Error("could not create note", zap.Error(err))
		return models.Note{}, err
	}

	return note, nil
}

func (r NotesService) GetNote(ctx context.Context, req models.GetNoteRequest) (models.Note, error) {
	logger := logging.Logger()

	note, err := r.repo.GetNote(ctx, req.ID)
	if err != nil {
		logger.Error("could not fetch note", zap.Error(err))
		return models.Note{}, err
	}

	return note, nil
}
