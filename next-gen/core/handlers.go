package core

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/rs/zerolog/log"
	"github.com/go-chi/chi/v5"

	"github.com/ccoverstreet/Salad-Notes/markdown"
)

func WrapRoute(raw func(http.ResponseWriter, *http.Request, *Core), core *Core) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		raw(w, r, core)
	}
}

func HTTPErrorHandler(w http.ResponseWriter, r *http.Request, err error, message string, code int) {
	fmt.Println(err)
	log.Error().
		Str("err", err.Error()).
		Str("route", r.URL.Path).
		Int("statusCode", code).
		Msg(message)
}

func ReadJSONBody(r *http.Request, v interface{}) error {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(b, v)
}

func SendJSONResponse(w http.ResponseWriter, r *http.Request, v interface{}) {
	b, err := json.Marshal(v)
	if err != nil {
		HTTPErrorHandler(w, r, err, "Unable to marshal JSON to string", 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

func GetAllItemsHandler(w http.ResponseWriter, r *http.Request, core *Core) {
	items, err := core.DB.GetAllItems()
	if err != nil {
		HTTPErrorHandler(w, r, err, "Unable to query database for items", 500)
		return
	}

	SendJSONResponse(w, r, items)
}


func GetItemsByNameHandler(w http.ResponseWriter, r *http.Request, core *Core) {
	req := struct {
		Fragment string `json:"fragment"`
	}{}

	err := ReadJSONBody(r, &req)
	if err != nil {
		HTTPErrorHandler(w, r, err, "Unable to parse request body", 400)
		return 
	}

	items, err := core.DB.GetItemsByName(req.Fragment)
	if err != nil {
		HTTPErrorHandler(w, r, err, "Unable to query database for matching names", 500)
		return 
	}

	SendJSONResponse(w, r, items)
}

func GetAllTagsHandler(w http.ResponseWriter, r *http.Request, core *Core) {
	tags, err := core.DB.GetAllTags()
	if err != nil {
		HTTPErrorHandler(w, r, err, "Unable to query database for tags", 500)
		return
	}

	SendJSONResponse(w, r, tags)
}

func GetItemContentHandler(w http.ResponseWriter, r *http.Request, core *Core) {
	uuid := chi.URLParam(r, "itemId")

	b, err := core.DB.GetItemContent(uuid)
	if err != nil {
		HTTPErrorHandler(w, r, err, "Unable to retrieve item", 500)
		return 
	}

	w.Write(b)
}


func CreateItemHandler(w http.ResponseWriter, r *http.Request, core *Core) {
	req := struct {
		Name      string `json:"name"`
		Extension string `json:"extension"`
	}{}

	err := ReadJSONBody(r, &req)
	if err != nil {
		HTTPErrorHandler(w, r, err, "Unable to read JSON body", 400)
		return
	}

	item, err := core.DB.CreateItem(req.Name, req.Extension, []byte{})
	if err != nil {
		HTTPErrorHandler(w, r, err, "Unable to create item", 400)
		return
	}

	SendJSONResponse(w, r, item)
}

func UpdateItemContentHandler(w http.ResponseWriter, r *http.Request, core *Core) {
	err := r.ParseMultipartForm(0)
	if err != nil {
		HTTPErrorHandler(w, r, err, "Unable to parse multipart form", 400)
		return
	}

	fmt.Println(r.Form)
	itemId, ok := r.Form["itemId"]
	if !ok {
		HTTPErrorHandler(w, r, fmt.Errorf("Trying to access non-existent key"), "Item ID not provided", 400)
		return
	}

	// Image or binary content (file)
	fmt.Println(r.Header.Get("Content-Type"))
	fmt.Println("FILE", r.MultipartForm.File)
	files, ok := r.MultipartForm.File["file"]
	if ok {
		// Content is indeed a file in multipart
		if len(files) < 1 {
			HTTPErrorHandler(w, r, err, "No files present", 400)
			return
		}

		fileHeader := files[0]
		file, err := fileHeader.Open()
		if err != nil {
			HTTPErrorHandler(w, r, err, "Unable to open file header", 500)
			return
		}

		b, err := io.ReadAll(file)
		if err != nil {
			HTTPErrorHandler(w, r, err, "Unable to read provided file", 500)
			return
		}

		err = core.DB.UpdateItemContent(string(itemId[0]), b)
		if err != nil {
			HTTPErrorHandler(w, r, err, "Unable to update item content", 500)
			return
		}
	} else {
		fmt.Println("HASDKJASDJ", r.Form["file"])

		err = core.DB.UpdateItemContent(string(itemId[0]), []byte(r.Form["file"][0]))
	}

	

	w.WriteHeader(200)
}

func MarkdownToHTMLHandler(w http.ResponseWriter, r *http.Request, core *Core) {
	uuid := chi.URLParam(r, "itemId")

	item, err := core.DB.GetItemMeta(uuid)
	if err != nil {
		HTTPErrorHandler(w, r, err, "Unable to get item metadata", 500)
		return 
	}

	if item.Extension != ".md"{
		HTTPErrorHandler(w, r, err, "Cannot get HTML version of this file", 500)
		return
	}

		
	b, err := core.DB.GetItemContent(uuid)
	if err != nil {
		HTTPErrorHandler(w, r, err, "Unable to get markdown content", 500)	
		return 
	}

	bout, err := markdown.MarkdownToHTML(b)
	if err != nil {
		HTTPErrorHandler(w, r, err, "Umable to convert markdown to HTML", 500)
	}

	w.Write(bout)
}
