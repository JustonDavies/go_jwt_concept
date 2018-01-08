build:
#	dep ensure
	rm -f bin/*

	env GOOS=linux go build -ldflags="-s -w" -o bin/go_jwt_concept  src/main.go

	chmod -R 700 bin