run:
	-go run cmd/main.go

dev:
	nodemon --signal SIGHUP --exec "make run" -e "go"

build_darwin:
	-env GOOS=darwin GOARCH=amd64 go build -v -o build/cmd/project cmd/main.go
build_arm:
	-env GOOS=linux GOARCH=arm64 go build -v -o build/cmd/project_arm64 cmd/main.go
build: build_darwin build_arm