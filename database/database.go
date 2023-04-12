package database

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"sync"

	"github.com/rs/zerolog/log"
)

type DocType string

const (
	MD  DocType = "md"
	PNG         = "png"
	SVG         = "svg"
	JPG         = "jpg"
)

type Document struct {
	UID      string   `json:"uid"`
	Name     string   `json:"name"`
	FileType DocType  `json:"fileType"`
	Tags     []string `json:"tags"`
}

type DatabaseHandle interface {
	AddItem(name string, contentType DocType, tags []string, content []byte) (Document, error)
	DeleteItem(uid string) error
	WriteItem(uid string, content []byte) error
	ReadFile(uid string) ([]byte, error)
	UpdateItemMeta(uid string, name string, tags []string) error
	GetItemByUID(UID string) (Document, error)
	GetItemsByTags(tags []string) []Document
	GetItemsByName(nameFrag string) []Document
	Save()
}

var DocTypeExt = map[DocType]string{
	MD:  ".md",
	PNG: ".png",
	SVG: ".svg",
	JPG: ".png",
}

var DocTypeMimeType = map[DocType]string{
	MD:  "text/markdown",
	PNG: "image/png",
	SVG: "image/svg+xml",
	JPG: "image/jpeg",
}

func DocTypeToMime(d DocType) string {
	return DocTypeMimeType[d]
}

func GenerateUID() string {
	const letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	b := make([]byte, 32)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}

type SaladDB struct {
	sync.RWMutex
	Name       string              `json:"-"`
	DataDir    string              `json:"dataDir"`
	ContentMap map[string]Document `json:"contentMap"`
}

func CreateSaladDB(name string, dataDir string) *SaladDB {
	return &SaladDB{
		sync.RWMutex{},
		name,
		dataDir,
		make(map[string]Document),
	}
}

func (db *SaladDB) Print() {
	db.RLock()
	defer db.RUnlock()
	b, err := json.MarshalIndent(db.ContentMap, "", "    ")
	log.Printf("%v %v", string(b), err)
}

// This spawns a goroutine that will save once the
// caller exits and releases the mutex
func (db *SaladDB) Save() {
	go func() {
		db.Lock()
		defer db.Unlock()

		jsonStr, err := json.MarshalIndent(db, "", "\t")
		if err != nil {
			log.Error().
				Msg("ERROR: Unable to save database meta file. File relations are no longer being saved")
			return
		}

		filename := fmt.Sprintf("%s/saladbowl.json", db.DataDir)
		log.Printf("Saving database meta information to %s", filename)
		os.WriteFile(filename, jsonStr, 0644)
	}()
}

func (db *SaladDB) AddItem(name string, contentType DocType, tags []string, content []byte) (Document, error) {
	db.Lock()
	defer db.Unlock()

	// Generate unique UID by creating and checking
	// against DB
	uid := ""
	for {
		uid = GenerateUID()
		if _, ok := db.ContentMap[uid]; !ok {
			break
		}
	}

	_, ok := DocTypeExt[contentType]
	if !ok {
		return Document{}, fmt.Errorf("Invalid doctype")
	}

	err := os.WriteFile(fmt.Sprintf("%s/%s", db.DataDir, uid),
		content, 0644)
	if err != nil {
		return Document{}, err
	}

	newDoc := Document{uid, name, contentType, tags}

	db.ContentMap[uid] = newDoc

	db.Save()
	return newDoc, err
}

func (db *SaladDB) DeleteItem(uid string) error {
	db.Lock()
	defer db.Unlock()

	doc, ok := db.ContentMap[uid]
	if !ok {
		return fmt.Errorf("UID not found in database")
	}

	err := os.Remove(fmt.Sprintf("%s/%s", db.DataDir, doc.UID))
	if err != nil {
		return fmt.Errorf("Unable to remove entry - %v", err)
	}

	delete(db.ContentMap, uid)
	db.Save()

	return nil
}

func (db *SaladDB) WriteItem(uid string, content []byte) error {
	db.Lock()
	defer db.Unlock()

	doc, ok := db.ContentMap[uid]
	if !ok {
		return fmt.Errorf("UID not found in database")
	}

	return os.WriteFile(fmt.Sprintf("%s/%s", db.DataDir, doc.UID), content, 0644)
}

func (db *SaladDB) ReadFile(uid string) ([]byte, error) {
	return os.ReadFile(fmt.Sprintf("%s/%s", db.DataDir, uid))
}

func (db *SaladDB) UpdateItemMeta(uid string, name string, tags []string) error {
	db.Lock()
	defer db.Unlock()

	doc, ok := db.ContentMap[uid]
	if !ok {
		return fmt.Errorf("UID not found in database")
	}

	db.ContentMap[uid] = Document{
		uid,
		name,
		doc.FileType,
		tags,
	}

	return nil
}

func (db *SaladDB) GetItemByUID(uid string) (Document, error) {
	db.RLock()
	defer db.RUnlock()

	entry, ok := db.ContentMap[uid]
	if !ok {
		return Document{}, fmt.Errorf("UID not found in database")
	}

	return entry, nil
}

// This search function does not look for an exact match
// Instead, it matches against anyname that contains the
// specified name
func (db *SaladDB) GetItemsByName(name string) []Document {
	db.RLock()
	defer db.RUnlock()

	entries := make([]Document, 0)
	for _, doc := range db.ContentMap {
		if strings.Contains(strings.ToLower(doc.Name), strings.ToLower(name)) {
			entries = append(entries, doc)
		}
	}

	return entries
}

func (db *SaladDB) GetItemsByTags(tags []string) []Document {
	db.RLock()
	defer db.RUnlock()

	entries := make([]Document, 0)

	for _, doc := range db.ContentMap {
		for _, t := range tags {
			for _, dt := range doc.Tags {
				if strings.ToLower(dt) == strings.ToLower(t) {
					entries = append(entries, doc)
				}
			}
		}
	}

	return entries
}
