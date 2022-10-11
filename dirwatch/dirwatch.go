package dirwatch

import (
	"log"
	"os"

	"github.com/fsnotify/fsnotify"
)

type DirWatcher struct {
	watcher  *fsnotify.Watcher
	callback func(fsnotify.Event)
}

func CreateDirWatcher(rootDir string) (*DirWatcher, error) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}

	watcher.Add(rootDir)
	watcher.Add("./misc")

	return &DirWatcher{
		watcher,
		func(fsnotify.Event) {},
	}, nil
}

func (dw *DirWatcher) Listen() {
	go func() {
		for {
			select {
			case event, ok := <-dw.watcher.Events:
				if !ok {
					return
				}

				// Need to add newly created directories to watch list
				if event.Op == fsnotify.Create {
					isdir, err := IsDir(event.Name)
					if err != nil {
						log.Printf("ERROR: %v", err)
						continue
					}

					if isdir {
						dw.watcher.Add(event.Name)
						log.Printf("Watching new directory: %s", event.Name)
					}
				}

				dw.callback(event)

			case err, ok := <-dw.watcher.Errors:
				if !ok {
					return
				}
				log.Println("Error:", err)
			}
		}
	}()
}

func IsDir(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false, err
	}

	return fileInfo.IsDir(), nil
}

func (dw *DirWatcher) SetCallback(callback func(fsnotify.Event)) {
	dw.callback = callback
}
