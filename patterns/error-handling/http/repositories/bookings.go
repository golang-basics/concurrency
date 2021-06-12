package repositories

import (
	"context"
	"encoding/json"
	"time"

	"github.com/boltdb/bolt"
	"go.uber.org/zap"

	"github.com/steevehook/http/logging"
	"github.com/steevehook/http/models"
)

const (
	bookingsBucket = "bookings"
	roomsBucket    = "rooms"
)

type db interface {
	View(func(tx *bolt.Tx) error) error
	Update(func(tx *bolt.Tx) error) error
}

// NewBookings creates a new instance of BookingsRepository
func NewBookings(db db) BookingsRepository {
	return BookingsRepository{
		db:    db,
		rooms: map[string]map[int]bool{},
	}
}

// BookingsRepository represents the Bookings repository that will interact with the database
type BookingsRepository struct {
	db    db
	rooms map[string]map[int]bool
}

func (r BookingsRepository) Init(numberOfRooms int) error {
	logger := logging.Logger()
	rooms, err := r.getRooms(models.DefaultHotelID, numberOfRooms)
	if err != nil {
		logger.Error("could not fetch rooms", zap.Error(err))
		return err
	}

	r.rooms[models.DefaultHotelID] = rooms
	return nil
}

// GetBooking fetches a booking from the database
func (r BookingsRepository) GetBooking(ctx context.Context, id string) (models.Booking, error) {
	logger := logging.Logger()
	booking := models.Booking{}
	err := r.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bookingsBucket))
		notFoundErr := models.ResourceNotFoundError{
			Message: "could not find booking with id: " + id,
		}
		if bucket == nil {
			return notFoundErr
		}

		bs := bucket.Get([]byte(id))
		if len(bs) == 0 {
			return notFoundErr
		}
		return json.Unmarshal(bs, &booking)
	})
	if err != nil {
		logger.Error("could not fetch booking", zap.Error(err))
		return models.Booking{}, err
	}

	return booking, nil
}

// CreateBooking creates and saves a booking inside the database
func (r BookingsRepository) CreateBooking(ctx context.Context, booking models.Booking) (models.Booking, error) {
	logger := logging.Logger()
	if booking.HotelID != models.DefaultHotelID {
		return models.Booking{}, r.hotelNotFoundError(booking.HotelID)
	}

	isFull := true
	for roomNumber, isFree := range r.rooms[booking.HotelID] {
		if isFree {
			continue
		}
		isFull = false
		r.rooms[booking.HotelID][roomNumber] = true
		booking.RoomNumber = roomNumber
		break
	}
	if isFull {
		return models.Booking{}, models.HotelFullError{HotelID: booking.HotelID}
	}

	err := r.db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(bookingsBucket))
		if err != nil {
			logger.Error("could not create bookings bucket", zap.Error(err))
			return err
		}
		err = r.save(bucket, booking.ID, booking)
		if err != nil {
			return err
		}

		bucket, err = tx.CreateBucketIfNotExists([]byte(roomsBucket))
		if err != nil {
			logger.Error("could not create rooms bucket", zap.Error(err))
			return err
		}
		return r.save(bucket, booking.HotelID, r.rooms[booking.HotelID])
	})
	if err != nil {
		return models.Booking{}, err
	}

	return booking, nil
}

func (r BookingsRepository) DeleteExpiredBookings(ctx context.Context) (int, error) {
	logger := logging.Logger()
	bookings := make([]models.Booking, 0)
	err := r.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bookingsBucket))
		if bucket == nil {
			return nil
		}
		err := bucket.ForEach(func(k, v []byte) error {
			var booking models.Booking
			err := json.Unmarshal(v, &booking)
			if err != nil {
				return err
			}

			if time.Now().UTC().Sub(booking.EndsAt) > 0 {
				bookings = append(bookings, booking)
				err := bucket.Delete(k)
				if err != nil {
					return err
				}
			}

			return nil
		})
		if err != nil {
			return err
		}

		bucket = tx.Bucket([]byte(roomsBucket))
		if bucket == nil {
			return nil
		}
		for _, booking := range bookings {
			r.rooms[booking.HotelID][booking.RoomNumber] = false
			err := r.save(bucket, booking.HotelID, r.rooms[booking.HotelID])
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		logger.Error("could not delete expired bookings", zap.Error(err))
		return 0, err
	}
	return len(bookings), nil
}

func (r BookingsRepository) getRooms(hotelID string, numberOfRooms int) (map[int]bool, error) {
	logger := logging.Logger()
	rooms := map[int]bool{}
	for i := 0; i < numberOfRooms; i++ {
		rooms[i+1] = false
	}

	err := r.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(roomsBucket))
		if bucket == nil {
			return nil
		}

		bs := bucket.Get([]byte(hotelID))
		if len(bs) == 0 {
			return r.hotelNotFoundError(hotelID)
		}
		return json.Unmarshal(bs, &rooms)
	})
	if err != nil {
		logger.Error("could not fetch rooms", zap.Error(err))
		return map[int]bool{}, err
	}

	return rooms, nil
}

func (r BookingsRepository) save(bucket *bolt.Bucket, key string, value interface{}) error {
	logger := logging.Logger()
	bs, err := json.Marshal(value)
	if err != nil {
		logger.Error("could not marshal data", zap.Error(err))
		return err
	}
	err = bucket.Put([]byte(key), bs)
	if err != nil {
		logger.Error("could not put data inside bolt db", zap.Error(err))
		return err
	}
	return nil
}

func (r BookingsRepository) hotelNotFoundError(hotelID string) error {
	return models.ResourceNotFoundError{
		Message: "could not find hotel with id: " + hotelID,
	}
}
