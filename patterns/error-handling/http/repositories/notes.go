package repositories

import (
	"context"
	"encoding/json"

	"github.com/boltdb/bolt"
	"go.uber.org/zap"

	"github.com/steevehook/http/logging"
	"github.com/steevehook/http/models"
)

const bucketName = "notes"

type db interface {
	View(func(tx *bolt.Tx) error) error
	Update(func(tx *bolt.Tx) error) error
}

// NewNotes creates a new instance of Notes repository
func NewNotes(db db) NotesRepository {
	return NotesRepository{
		db: db,
	}
}

// NotesRepository represents the Notes repository that will interact with the database
type NotesRepository struct {
	db db
}

// CreateNote creates and saves a note inside the database
func (r NotesRepository) CreateNote(ctx context.Context, note models.Note) error {
	logger := logging.Logger()
	return r.db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(bucketName))
		if err != nil {
			logger.Error("could not create bucket", zap.Error(err))
			return err
		}

		bs, err := json.Marshal(note)
		if err != nil {
			logger.Error("could not marshal note", zap.Error(err))
			return err
		}

		return bucket.Put([]byte(note.ID), bs)
	})
}

// GetNote fetches a note from the database
func (r NotesRepository) GetNote(ctx context.Context, id string) (models.Note, error) {
	logger := logging.Logger()
	note := models.Note{}
	err := r.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))
		notFoundErr := models.ResourceNotFoundError{
			Message: "could not find note with id: " + id,
		}
		if bucket == nil {
			return notFoundErr
		}

		bs := bucket.Get([]byte(id))
		if len(bs) == 0 {
			return notFoundErr
		}

		err := json.Unmarshal(bs, &note)
		if err != nil {
			logger.Error("could not unmarshal data", zap.Error(err))
			return err
		}

		return nil
	})
	if err != nil {
		logger.Error("could not fetch note", zap.Error(err))
		return models.Note{}, err
	}

	return note, nil
}
