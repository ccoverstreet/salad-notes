all:

dev:
	nodemon -x "go run ." -w core -w notestorage -w markdown -w main.go --signal SIGTERM -e go

test:
	cd notestorage; go test; cd ..;

init_db: delete_db
	sqlite3 storage.db < init.sql

delete_db:
	rm notesbowl/*
	rm storage.db
