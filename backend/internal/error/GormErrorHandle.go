package error

import (
	"errors"

	"gorm.io/gorm"
)

func HandleGormError(err error) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &NotFoundError{Message: err.Error()}
	}

	if errors.Is(err, gorm.ErrDuplicatedKey) {
		return &DuplicateKeyError{Message: err.Error()}
	}

	return err
}
