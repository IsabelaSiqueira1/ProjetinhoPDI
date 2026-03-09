package linkedlist

import "testing"

type Professional struct {
	ID   string
	Name string
}

func TestInsertAndLen(t *testing.T) {
	list := New[Professional]()

	list.Insert(Professional{ID: "p1", Name: "Ana"})
	list.Insert(Professional{ID: "p2", Name: "Bruno"})
	list.Insert(Professional{ID: "p3", Name: "Carla"})

	if got := list.Len(); got != 3 {
		t.Fatalf("expected len 3, got %d", got)
	}
}

func TestFindFoundAndNotFound(t *testing.T) {
	list := New[Professional]()
	list.Insert(Professional{ID: "p1", Name: "Ana"})
	list.Insert(Professional{ID: "p2", Name: "Bruno"})

	found, ok := list.Find(func(p Professional) bool { return p.ID == "p2" })
	if !ok {
		t.Fatal("expected to find professional p2")
	}
	if found.Name != "Bruno" {
		t.Fatalf("expected Bruno, got %s", found.Name)
	}

	_, ok = list.Find(func(p Professional) bool { return p.ID == "p9" })
	if ok {
		t.Fatal("did not expect to find professional p9")
	}
}

func TestUpdateFoundAndNotFound(t *testing.T) {
	list := New[Professional]()
	list.Insert(Professional{ID: "p1", Name: "Ana"})
	list.Insert(Professional{ID: "p2", Name: "Bruno"})

	updated, ok := list.Update(
		func(p Professional) bool { return p.ID == "p2" },
		func(p *Professional) { p.Name = "Bruno Silva" },
	)
	if !ok {
		t.Fatal("expected update to find professional p2")
	}
	if updated.Name != "Bruno Silva" {
		t.Fatalf("expected updated name Bruno Silva, got %s", updated.Name)
	}

	check, ok := list.Find(func(p Professional) bool { return p.ID == "p2" })
	if !ok || check.Name != "Bruno Silva" {
		t.Fatalf("expected persisted update Bruno Silva, got %+v (ok=%v)", check, ok)
	}

	_, ok = list.Update(
		func(p Professional) bool { return p.ID == "p9" },
		func(p *Professional) { p.Name = "Nao existe" },
	)
	if ok {
		t.Fatal("did not expect update to find professional p9")
	}
}

func TestForEachPreservesInsertionOrder(t *testing.T) {
	list := New[int]()
	list.Insert(10)
	list.Insert(20)
	list.Insert(30)

	var got []int
	list.ForEach(func(v int) {
		got = append(got, v)
	})

	expected := []int{10, 20, 30}
	if len(got) != len(expected) {
		t.Fatalf("expected %d items, got %d", len(expected), len(got))
	}

	for i := range expected {
		if got[i] != expected[i] {
			t.Fatalf("expected value %d at index %d, got %d", expected[i], i, got[i])
		}
	}
}
