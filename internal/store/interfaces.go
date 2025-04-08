package store

import (
	"iter"

	"github.com/google/uuid"
)

type Storable[T any] interface {
	Create(T) error
	Get(uuid.UUID) (T, error)
	List() iter.Seq[T]
	Update(T) error
	Delete(uuid.UUID) error
}

type AsyncStorable[T any] interface {
	Create(T) chan error
	Get(uuid.UUID) (chan T, chan error)
	List() chan iter.Seq[T]
	Update(T) chan error
	Delete(uuid.UUID) chan error
}
