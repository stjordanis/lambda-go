build:
	dep ensure
	env GOOS=linux go build -ldflags="-s -w" -o bin/hello functions/hello/main.go