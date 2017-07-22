package boltstore

import (
	"github.com/timshannon/bolthold"
)

// BoltStore stores all the information about the store
type BoltStore struct {
	store *bolthold.Store
}

// NewStore returns a new bolthold store
func NewStore(filename string) (*BoltStore, error) {
	b := &BoltStore{}
	store, err := bolthold.Open(filename, 0666, nil)
	if err != nil {
		return nil, err
	}
	b.store = store
	return b, nil
}

// Get returns data from the store
func (b *BoltStore) Get(key string, val interface{}) error {
	if err := b.store.Get(key, val); err != nil {
		return err
	}
	return nil
}

// Upsert inserts or updates data
func (b *BoltStore) Upsert(key string, val interface{}) error {
	if err := b.store.Upsert(key, val); err != nil {
		return err
	}
	return nil
}

// Find searches the database
func (b *BoltStore) Find(field string, equal string, val interface{}) error {
	if err := b.store.Find(val, bolthold.Where(field).Eq(equal)); err != nil {
		return err
	}
	return nil
}
