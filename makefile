
dev: 
	nodemon -w . -w main.go -w app/app.go -w assets/index.html -w assets/index.css -w assets/index.js --exec "(go run main.go --port=8080) || exit 2"   --signal SIGTERM

multiplatform: 
	GOOS=linux GOARCH=amd64 go build -o build/salad-notes_linux-amd64
	GOOS=linux GOARCH=arm64 go build -o build/salad-notes_linux-arm64
	GOOS=windows GOARCH=amd64 go build -o build/salad-notes_windows-amd64
	GOOS=darwin GOARCH=arm64 go build -o build/salad-notes_darwin-arm64
	GOOS=darwin GOARCH=amd64 go build -o build/salad-notes_darwin-amd64
