package datastore

import (
	"errors"
	"testing"
	"reflect"
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
	nomore := mockObject{Name:"nomore"}
	if err := Get(context, &nomore); err != ErrNotFound {
		return errors.New("nomore objects allowed because nomore object found")
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
	if err := Get(storageContext, &ob2); err != nil {
		t.Fatalf("Could not get object %s: %s", ob1, err)
	}

	if !reflect.DeepEqual(ob1, ob2) {
		t.Fatalf("ob1 '%v' != ob2 '%v'", ob1, ob2)
	}

	notfindable := mockObject{Name: "does not exist"}
	if err := Get(storageContext, notfindable); err == nil {
		t.Fatalf("Should have not found %s", notfindable)
	}

	nomore := mockObject{Name: "nomore"}
	if err := Put(storageContext, nomore); err != nil {
		t.Fatalf("Could not add 'nomore' object", err)
	}

	if err := Put(storageContext, ob1); err == nil {
		t.Fatalf("expected error after nomore object")
	}

}
