dev: 
	nodemon -w . -w main.go --exec go run main.go --signal SIGTERM

multiplatform:
	GOOS=linux GOARCH=amd64 go build -o build/salad-notes_linux-amd64
	GOOS=linux GOARCH=arm64 go build -o build/salad-notes_linux-arm64
	GOOS=windows GOARCH=amd64 go build -o build/salad-notes_windows-amd64
