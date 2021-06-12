package transport

import (
	"encoding/json"
	"net/http"

	"go.uber.org/zap"

	"github.com/steevehook/http/logging"
	"github.com/steevehook/http/models"
)

const (
	dataValidationErrorType   = "data_validation_error"
	formatValidationErrorType = "format_validation_error"
	resourceNotFoundErrorType = "resource_not_found"
	serviceErrorType          = "service_error"
)

// SendHTTPError converts errors into HTTP JSON errors
func SendHTTPError(w http.ResponseWriter, err error) {
	httpError := toHTTPError(err)
	SendJSON(w, httpError.Code, httpError)
}

// SendJSON converts application response into JSON responses
func SendJSON(w http.ResponseWriter, statusCode int, response interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(&response); err != nil {
		logging.Logger().Error("could not encode response", zap.Error(err))
	}
}

func toHTTPError(err error) models.HTTPError {
	switch e := err.(type) {
	case models.HTTPError:
		return e

	case models.InvalidJSONError:
		return models.HTTPError{
			Code:    http.StatusBadRequest,
			Type:    formatValidationErrorType,
			Message: e.Message,
		}

	case models.FormatValidationError:
		return models.HTTPError{
			Code:    http.StatusBadRequest,
			Type:    formatValidationErrorType,
			Message: e.Message,
		}

	case models.DataValidationError:
		return models.HTTPError{
			Code:    http.StatusBadRequest,
			Type:    dataValidationErrorType,
			Message: e.Message,
		}

	case models.ResourceNotFoundError:
		return models.HTTPError{
			Code:    http.StatusNotFound,
			Type:    resourceNotFoundErrorType,
			Message: e.Message,
		}

	case models.HotelFullError:
		return models.HTTPError{
			Code:    http.StatusServiceUnavailable,
			Type:    serviceErrorType,
			Message: e.Error(),
		}

	default:
		return models.HTTPError{
			Code:    http.StatusInternalServerError,
			Type:    serviceErrorType,
			Message: "server was not able to process your request",
		}
	}
}
