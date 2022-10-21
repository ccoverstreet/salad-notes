package app

import (
	"fmt"
	"net/http"
	"os"
	"path"

	"github.com/ccoverstreet/salad-notes/pandoc"
	"github.com/ccoverstreet/salad-notes/util"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/rs/zerolog/log"
)

type Client struct {
	conn *websocket.Conn
}

func CreateClient(conn *websocket.Conn) *Client {
	return &Client{conn}
}

type SaladApp struct {
	clients map[string]*Client
	router  *mux.Router
}

func CreateSaladApp() *SaladApp {
	app := &SaladApp{}

	router := mux.NewRouter()

	router.HandleFunc("/", HomeHandler)
	router.HandleFunc("/saladnotes/connectClient", WrapRoute(ConnectClient, app))
	router.HandleFunc("/saladnotes/assets/{file}", AssetHandler)
	router.PathPrefix("/").Handler(http.HandlerFunc(StaticFileHandler))

	app.router = router
	return app
}

func (app *SaladApp) Listen(port int) {
	http.ListenAndServe(fmt.Sprintf(":%d", port), app.router)
}

func (app *SaladApp) AddClient(remoteAddr string, client *Client) error {
	// Check if client with given remoteAddr already exists
	_, ok := app.clients[remoteAddr]
	if ok {
		return fmt.Errorf("Client %s already connected", remoteAddr)
	}

	app.clients[remoteAddr] = client

	return nil
}

func WrapRoute(handler func(http.ResponseWriter, *http.Request, *SaladApp), app *SaladApp) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(w, r, app)
	}
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	b, err := os.ReadFile("assets/index.html")
	if err != nil {
		return
	}

	w.Write(b)
}

var WSUPGRADER = websocket.Upgrader{}

func ConnectClient(w http.ResponseWriter, r *http.Request, app *SaladApp) {
	log.Info().
		Str("remoteAddr", r.RemoteAddr).
		Msg("Client connecting")

	conn, err := WSUPGRADER.Upgrade(w, r, nil)
	if err != nil {
		util.HandleHTTPErrorAndLog(w, 400, err)
		return
	}

	conn.WriteMessage(websocket.TextMessage, []byte(`{"modifiedFile": "test-notebook/NE-551/ion-matter.md"}`))

	err = app.AddClient(r.RemoteAddr, CreateClient(conn))
	if err != nil {
		util.HandleHTTPErrorAndLog(w, 500, err)
		return
	}

	return
}

var ASSETMAP map[string]string = map[string]string{
	"index.js": "assets/index.js",
}

var MIMEMAP map[string]string = map[string]string{
	".js":  "text/javascript",
	".css": "text/css",
}

func AssetHandler(w http.ResponseWriter, r *http.Request) {
	file := mux.Vars(r)["file"]

	filePath, ok := ASSETMAP[file]
	if !ok {
		util.HandleHTTPErrorAndLog(w, 400, fmt.Errorf("Asset '%s' not found", file))
		return
	}

	ext := path.Ext(file)

	mime, ok := MIMEMAP[ext]
	if !ok {
		// TODO: Proper logging
		panic("TODO")
	}

	b, err := os.ReadFile(filePath)
	if err != nil {
		//TODO
		panic("TODO")
	}

	w.Header().Set("Content-Type", mime)
	w.Write(b)
}

func StaticFileHandler(w http.ResponseWriter, r *http.Request) {
	ext := path.Ext(r.RequestURI)

	log.Info().
		Str("URI", r.RequestURI).
		Msg("Static file requested")

	// Use pandoc to create HTML content of Markdown file
	if ext == ".md" {
		b, err := pandoc.RunPandoc("."+r.RequestURI, []string{"--mathjax"})
		log.Printf("%s %v", b, err)
		w.Write(b)
	} else { // Send binary version of any other file
		b, err := os.ReadFile("." + r.RequestURI)
		if err != nil {
			log.Error().
				Err(err).
				Msg("File not found")
			return
		}

		w.Write(b)
	}
}
