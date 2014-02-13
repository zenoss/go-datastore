package datastore

import (
	"errors"
)

type Storeable interface {
	Type() string
	Key() string
	Validate(context StorageContext) error
}

type StorageContext interface {
	Get(Storeable) error
	Put(Storeable) error
	Delete(Storeable) error
}

type DataStore struct {
	schema string
}

var (
	ErrUnimplemented = errors.New("unimplemented")
	ErrNotFound      = errors.New("not found")
)

// Get()s a value from the datastore; s is a pointer to a storable object who's Type()
// and Key() methods return the appropiate type and primary key, respectively.
// The retrieved value is stored in s.
func Get(context StorageContext, s Storeable) error {
	return context.Get(s)
}

// Put()s a value from the datastore; s is a pointer to a storable object who's Type()
// and Key() methods return the appropiate type and primary key, respectively
func Put(context StorageContext, s Storeable) error {
	if err := s.Validate(context); err != nil {
		return err
	}
	return context.Put(s)
}

// Delete()s an object from the datastore; s is a pointer to a storable object who's Type()
// and Key() methods return the appropiate type and primary key, respectively
func Delete(context StorageContext, s Storeable) error {
	return context.Delete(s)
}

// exists && search
