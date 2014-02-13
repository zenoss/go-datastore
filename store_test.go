package datastore

import (
	"errors"
	"testing"
)

type mockObject struct {
	Name      string
	Attribute string
}

func (m mockObject) Key() string {
	return m.Name
}

func (m mockObject) Type() string {
	return "MockObject"
}

func (m mockObject) Validate(context StorageContext) error {
	if len(m.Name) == 0 {
		return errors.New("mockObject must have a name")
	}
	return nil
}

func TestStores(t *testing.T) {
	storageContext, _ := NewMemoryStorageContext()

	ob1 := mockObject{
		Name:      "test",
		Attribute: "foo",
	}

	if err := Put(storageContext, ob1); err != nil {
		t.Fatalf("Could not store object %s: %s", ob1, err)
	}

	ob2 := mockObject{Name: "test"}
	if err := Get(storageContext, ob2); err != nil {
		t.Fatalf("Could not get object %s: %s", ob1, err)
	}

	notfindable := mockObject{Name: "does not exist"}
	if err := Get(storageContext, notfindable); err == nil {
		t.Fatalf("Should have not found %s", notfindable)
	}

}