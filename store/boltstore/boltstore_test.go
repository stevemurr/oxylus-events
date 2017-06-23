package boltstore

import (
	"testing"
)

type TestStruct struct {
	Name string
}

func TestUpsert(t *testing.T) {
	b, err := NewStore("bolt.db")
	if err != nil {
		t.Error("failed to allocate database")
	}
	testStruct := &TestStruct{Name: "hello"}
	b.Upsert("test", testStruct)
	b.Upsert("test2", testStruct)
	b.Upsert("test3", testStruct)
}

func TestGet(t *testing.T) {
	b, err := NewStore("bolt.db")
	if err != nil {
		t.Error("failed to allocate database")
	}
	var testStruct TestStruct
	if err := b.Get("test", &testStruct); err != nil {
		t.Error("failed to get key")
	}
	if testStruct.Name != "hello" {
		t.Error("name field incorrect")
	}
}

func TestFind(t *testing.T) {
	b, err := NewStore("bolt.db")
	if err != nil {
		t.Error("failed to allocate database")
	}
	var results []TestStruct
	if err := b.Find("Name", "hello", &results); err != nil {
		t.Error("failed to find results with name hello")
	}
	if len(results) != 3 {
		t.Error("length of results should be 3")
	}
}
