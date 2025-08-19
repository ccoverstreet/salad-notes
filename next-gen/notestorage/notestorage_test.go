package notestorage

import "testing"

func TestMulti(t *testing.T) {
	db, err := CreateNotesDatabase("storage.db")
	if err != nil {
		t.Fatalf("Unable to make storage struct: %v", err)
	}

	itemId, err := db.CreateItem("Test_1", ".md", []byte("Hello"))
	if err != nil {
		t.Fatalf("Unable to create item: %v", err)
	}

	tagId, err := db.CreateTag("Tag_1")
	if err != nil {
		t.Fatalf("Unable to create tag: %v", err)
	}

	err = db.AddTagToItem(tagId.Id, itemId.Id)
	if err != nil {
		t.Fatalf("Unable to add tag to item: %v", err)
	}

	err = db.RemoveTagFromItem(tagId.Id, itemId.Id)
	if err != nil {
		t.Fatalf("Unable to remove tag from item: %v", err)
	}
}
