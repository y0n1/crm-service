package store

import (
	"iter"
	"maps"
	"sync"

	"github.com/google/uuid"
	"github.com/y0n1/crm-service/internal/models/aggregates"
	"github.com/y0n1/crm-service/pkg/collections"
)

type MemoryStore struct {
	customers collections.Map[string, *aggregates.CustomerAggregate]
	mutex     sync.RWMutex
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		customers: collections.NewMap[string, *aggregates.CustomerAggregate](0),
		mutex:     sync.RWMutex{},
	}
}

func (r *MemoryStore) Create(aggregate *aggregates.CustomerAggregate) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if aggregate.Customer.ID != uuid.Nil {
		return ErrCustomerAlreadyExists
	}

	emailAlreadyExists := false
	r.customers.ForEach(func(key string, value *aggregates.CustomerAggregate) {
		if aggregate.Customer.Email == value.Customer.Email {
			emailAlreadyExists = true
		}
	})
	if emailAlreadyExists {
		return ErrCustomerEmailAlreadyExists
	}

	newId, err := uuid.NewV7()
	if err != nil {
		return err
	}

	aggregate.Customer.ID = newId
	id := aggregate.Customer.ID.String()
	if r.customers.Has(id) {
		return ErrCustomerAlreadyExists
	}

	r.customers.Set(id, aggregate)
	return nil
}

func (r *MemoryStore) Delete(id uuid.UUID) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if !r.customers.Has(id.String()) {
		return ErrCustomerNotFound
	}

	r.customers.Delete(id.String())
	return nil
}

func (r *MemoryStore) Get(id uuid.UUID) (*aggregates.CustomerAggregate, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	if aggregate, ok := r.customers.Get(id.String()); ok {
		return aggregate, nil
	} else {
		return nil, ErrCustomerNotFound
	}
}

func (r *MemoryStore) List() iter.Seq[*aggregates.CustomerAggregate] {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	return maps.Values(r.customers)
}

func (r *MemoryStore) Update(aggregate *aggregates.CustomerAggregate) error {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	id := aggregate.Customer.ID.String()
	if r.customers.Has(id) {
		r.customers.Set(id, aggregate)
		return nil
	} else {
		return ErrCustomerNotFound
	}
}
