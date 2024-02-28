package data

import (
	"datavisualisation/internal/core/domain"
	"io"
)

type DogRepository interface {
	GetDogDataFromCSV(data io.Reader, firstLineIsTitle bool) ([]domain.Dog, error)
}

type Repository interface {
	GetDogDataFromCSV(data io.Reader, firstLineIsTitle bool) ([]domain.Dog, error)
}
