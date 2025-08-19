package notestorage

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	_ "modernc.org/sqlite"
)

type NotesDatabase struct {
	db         *sqlx.DB
	storageDir string
}

type Item struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	Extension    string `json:"extension"`
	CreationDate string `json:"creationDate" db:"creationDate"`
}

type Tag struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

const INITSQL = `
DROP TABLE IF EXISTS items;
DROP TABLE IF EXISTS tags;
DROP TABLE IF EXISTS tagmap;

CREATE TABLE items (
	id TEXT NOT NULL,
	name TEXT NOT NULL,
	extension TEXT NOT NULL,
	creationDate TEXT NOT NULL
);

CREATE TABLE tags (
	id TEXT NOT NULL,
	name TEXT NOT NULL
);

CREATE TABLE tagmap (
	tagId TEXT NOT NULL,
	itemId TEXT NOT NULL
);`

func CreateNotesDatabase(filename string) (*NotesDatabase, error) {
	runInit := false
	if _, err := os.Stat(filename); errors.Is(err, os.ErrNotExist) {
		runInit = true
	}

	db, err := sqlx.Open("sqlite", filename)
	if err != nil {
		return nil, err
	}

	if runInit {
		_, err := db.Exec(INITSQL)
		if err != nil {
			panic(err)
		}
	}

	os.MkdirAll("notesbowl", os.ModePerm)

	return &NotesDatabase{
		db,
		"notesbowl",
	}, nil
}

// Extension should be in the format ".extension"
func (db *NotesDatabase) CreateItem(name string, extension string, content []byte) (Item, error) {
	id := uuid.New()
	date := time.Now().Local().Format("2006-01-02")

	_, err := db.db.NamedExec("INSERT INTO items (id, name, extension, creationDate) VALUES (:id, :name, :extension, :creationDate)", Item{id.String(), name, extension, date})
	fmt.Println(id)

	if err != nil {
		return Item{}, err
	}

	return Item{
		id.String(),
		name,
		extension,
		date,
	}, os.WriteFile(db.storageDir+"/"+id.String(), content, 0666)
}

func (db *NotesDatabase) DeleteItem(itemId string) error {
	_, err := db.db.Exec("DELETE FROM items WHERE id=?; DELETE FROM tagmap WHERE itemId=?", itemId, itemId)
	return err
}

func (db *NotesDatabase) GetItemMeta(uuid string) (Item, error) {
	item := Item{}
	err := db.db.Get(&item, "SELECT * FROM items WHERE id=?", uuid)
	return item, err
}

func (db *NotesDatabase) GetItemContent(uuid string) ([]byte, error) {
	return os.ReadFile(db.storageDir + "/" + uuid)
}

func (db *NotesDatabase) GetAllItems() ([]Item, error) {
	items := []Item{}
	err := db.db.Select(&items, "SELECT * FROM items")
	return items, err
}

// Returns items where name contains search term
func (db *NotesDatabase) GetItemsByName(fragment string) ([]Item, error) {
	items := []Item{}
	err := db.db.Select(&items, "SELECT * FROM items WHERE instr(lower(name), lower(?))", fragment)

	return items, err
}

func (db *NotesDatabase) GetItemsByTag(tagId string) ([]Item, error) {
	items := []Item{}
	err := db.db.Select(&items, "SELECT * FROM items WHERE items.id IN (SELECT itemId FROM tagmap WHERE tagId = ?)", tagId)
	return items, err
}

func (db *NotesDatabase) UpdateItemContent(uuid string, content []byte) error {
	return os.WriteFile(db.storageDir+"/"+uuid, content, 0666)
}

func (db *NotesDatabase) CreateTag(name string) (Tag, error) {
	id := uuid.New()
	_, err := db.db.Exec("INSERT INTO tags (id, name) VALUES (?, ?)", id.String(), name)
	return Tag{id.String(), name}, err
}

func (db *NotesDatabase) DeleteTag(uuid string) error {
	_, err := db.db.Exec("DELETE FROM tags WHERE id=?", uuid)
	return err
}

func (db *NotesDatabase) AddTagToItem(tagId string, itemId string) error {
	_, err := db.db.Exec("INSERT INTO tagmap (tagId, itemId) VALUES (?, ?)", tagId, itemId)
	return err
}

func (db *NotesDatabase) RemoveTagFromItem(tagId string, itemId string) error {
	_, err := db.db.Exec("DELETE FROM tagmap WHERE itemId = ? AND tagId = ?", itemId, tagId)
	return err
}

func (db *NotesDatabase) GetAllTags() ([]Tag, error) {
	tags := []Tag{}
	err := db.db.Select(&tags, "SELECT * FROM tags")
	return tags, err
}

/*
func (db *NotesDatabase) GetItemsByDate(string) {
	items := []Item{}

	err := db.db.Select(&items, "SELECT * FROM items WHERE creationDate=?", "2024-07-19")
	fmt.Println(err, items)
}
*/
