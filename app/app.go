package app

import (
	"fmt"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"sync"

	"github.com/ccoverstreet/salad-notes/assets"
	"github.com/ccoverstreet/salad-notes/dirwatch"
	"github.com/ccoverstreet/salad-notes/pandoc"
	"github.com/ccoverstreet/salad-notes/util"
	"github.com/fsnotify/fsnotify"
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
	sync.RWMutex
	clients  map[string]*Client
	router   *mux.Router
	watchdir string
	watcher  *dirwatch.DirWatcher
}

func CreateSaladApp(watchdir string) *SaladApp {
	app := &SaladApp{}

	router := mux.NewRouter()

	router.HandleFunc("/", HomeHandler)
	router.HandleFunc("/saladnotes/connectClient", WrapRoute(ConnectClient, app))
	router.HandleFunc("/saladnotes/assets/{file}", AssetHandler)
	router.HandleFunc("/saladnotes/listDir", ListDirHandler)
	router.PathPrefix("/").Handler(http.HandlerFunc(StaticFileHandler))

	app.router = router
	app.watchdir = watchdir

	newWatcher, err := dirwatch.CreateDirWatcher(watchdir)
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("Unable to create directory watcher")
	}

	app.watcher = newWatcher
	app.watcher.SetCallback(app.DWCallback)
	app.clients = make(map[string]*Client)

	return app
}

func (app *SaladApp) Listen(port int) {
	app.watcher.Listen() // Spawn listener
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

// This is an intermediate function used as a callback
// to dirwatch.DirWatcher. This function only notifies clients
// upon file changes for Markdown files
func (app *SaladApp) DWCallback(event fsnotify.Event) {
	ext := path.Ext(event.Name)

	// Only interested in write events
	if event.Op != fsnotify.Write {
		return
	}

	// Ignore files that aren't Markdown files
	if ext != ".md" {
		return
	}

	app.NotifyClients(event.Name)
}

func (app *SaladApp) NotifyClients(modifiedFile string) {
	data := struct {
		ModifiedFile string `json:"modifiedFile"`
	}{modifiedFile}

	errString := ""

	for remoteAddr, client := range app.clients {
		log.Info().
			Str("remoteAddr", remoteAddr).
			Msg("Notifying client")

		err := client.conn.WriteJSON(data)
		if err != nil {
			delete(app.clients, remoteAddr)
			errString += err.Error() + ";"
		}
	}

	if len(errString) > 0 {
		log.Error().
			Err(fmt.Errorf("%s", errString)).
			Msg("Errors occured while notifying clients.")
	}
}

func WrapRoute(handler func(http.ResponseWriter, *http.Request, *SaladApp), app *SaladApp) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(w, r, app)
	}
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fileData, err := assets.GetAsset("index.html")
	if err != nil {
		return
	}

	w.Write(fileData.Bytes)
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

	err = app.AddClient(r.RemoteAddr, CreateClient(conn))
	if err != nil {
		util.HandleHTTPErrorAndLog(w, 500, err)
		return
	}

	return
}

func AssetHandler(w http.ResponseWriter, r *http.Request) {
	file := mux.Vars(r)["file"]

	fileData, err := assets.GetAsset(file)
	if err != nil {
		util.HandleHTTPErrorAndLog(w, 400, fmt.Errorf("Asset '%s' not found", file))
		return
	}

	w.Header().Set("Content-Type", fileData.MimeType)
	w.Write(fileData.Bytes)
}

func StaticFileHandler(w http.ResponseWriter, r *http.Request) {
	ext := path.Ext(r.RequestURI)

	log.Info().
		Str("URI", r.RequestURI).
		Msg("Static file requested")

	// Use pandoc to create HTML content of Markdown file
	if ext == ".md" {
		b, err := pandoc.RunPandoc("."+r.RequestURI, []string{"--mathjax"})
		if err != nil {
			log.Error().
				Err(err).
				Str("file", r.RequestURI).
				Msg("Error while using pandoc")
		}
		//log.Printf("%s %v", b, err)
		w.Write(b)
	} else { // Send binary version of any other file
		b, err := os.ReadFile("." + r.RequestURI)
		if err != nil {
			log.Error().
				Err(err).
				Msg("File not found")
			return
		}

		// Add MimeType if it exists in map
		// Otherwise, print a warning
		mimeType, ok := assets.MIMEMAP[ext]
		if !ok {
			log.Warn().
				Str("file", r.RequestURI).
				Msg("Valid MimeType not found")
		}

		if ok {
			w.Header().Set("Content-Type", mimeType)
		}

		w.Write(b)
	}
}

func ListDirHandler(w http.ResponseWriter, r *http.Request) {
	reqData := struct {
		DirName string `json:"dirName"`
	}{}

	err := util.UnmarshalJSONBody(r, &reqData)
	if err != nil {
		util.HandleHTTPErrorAndLog(w, 400, err)
		return
	}

	sanitizedDir := sanitizeDirName(reqData.DirName)

	files, err := util.ListDir(sanitizedDir)
	if err != nil {
		util.HandleHTTPErrorAndLog(w, 500, err)
		return
	}

	util.SendJSONResponse(w, files)
}

func sanitizeDirName(dirName string) string {
	out := filepath.Clean(dirName)
	if string(out[0]) == "/" {
		return out[1:]
	}

	return out
}
