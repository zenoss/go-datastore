package datastore

import (
	"encoding/json"
)

type datatype []byte

type MemoryContext struct {
	values map[string]map[string]datatype
}

// Assert that MemoryContext implements Context interface
var _ Context = MemoryContext{}

func NewMemoryContext() (context *MemoryContext, err error) {
	context = &MemoryContext{
		values: make(map[string]map[string]datatype),
	}
	return context, nil
}

func (context MemoryContext) Put(storable Storeable) error {

	data, err := json.Marshal(storable)
	if err != nil {
		return err
	}
	type_, ok := context.values[storable.Type()]
	if !ok {
		type_ = make(map[string]datatype)
		context.values[storable.Type()] = type_
	}
	type_[storable.Key()] = data
	return nil
}

func (context MemoryContext) Exists(storable Storeable) (bool, error) {
	type_, ok := context.values[storable.Type()]
	if !ok {
		return ok, nil
	}
	_, ok = type_[storable.Key()]
	return ok, nil
}

func (context MemoryContext) Get(storable Storeable) error {
	type_, ok := context.values[storable.Type()]
	if !ok {
		return ErrNotFound
	}
	val, ok := type_[storable.Key()]
	if !ok {
		return ErrNotFound
	}
	return json.Unmarshal(val, &storable)
}

func (context MemoryContext) Delete(storable Storeable) error {
	type_, ok := context.values[storable.Type()]
	if !ok {
		return nil
	}
	if _, found := type_[storable.Key()]; found {
		delete(type_, storable.Key())
	}
	return nil
}
