package datastore

type MemoryStorageContext struct {
	values map[string]map[string]Storeable
}

// Assert that MemoryStorageContext implements StorageContext interface
var _ StorageContext = MemoryStorageContext{}

func NewMemoryStorageContext() (context *MemoryStorageContext, err error) {
	context = &MemoryStorageContext{
		values: make(map[string]map[string]Storeable),
	}
	return context, nil
}

func (context MemoryStorageContext) Put(storable Storeable) error {
	type_, ok := context.values[storable.Type()]
	if !ok {
		type_ = make(map[string]Storeable)
		context.values[storable.Type()] = type_
	}
	type_[storable.Key()] = storable
	return nil
}

func (context MemoryStorageContext) Get(storable Storeable) error {
	type_, ok := context.values[storable.Type()]
	if !ok {
		return ErrNotFound
	}
	val, ok := type_[storable.Key()]
	if ok {
		storable = val
	} else {
		return ErrNotFound
	}
	return nil
}

func (context MemoryStorageContext) Delete(storable Storeable) error {
	type_, ok := context.values[storable.Type()]
	if !ok {
		return nil
	}
	if _, found := type_[storable.Key()]; found {
		delete(type_, storable.Key())
	}
	return nil
}
