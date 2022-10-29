VERSION:=$(shell git describe --abbrev=0 --tags)
BUILD_FLAGS=-ldflags="-X main.Version=$(VERSION)"

changelog-darwin-amd64:
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=1 go build $(BUILD_FLAGS) -o build/changelog-darwin-amd64 cmd/main.go
	zip -r build/changelog-darwin-amd64.zip build/changelog-darwin-amd64
	rm build/changelog-darwin-amd64

changelog-darwin-arm64:
	GOOS=darwin GOARCH=arm64 CGO_ENABLED=1 go build $(BUILD_FLAGS) -o build/changelog-darwin-arm64 cmd/main.go
	zip -r build/changelog-darwin-arm64.zip build/changelog-darwin-arm64
	rm build/changelog-darwin-arm64

changelog-linux-amd64:
	GOOS=linux GOARCH=amd64 go build $(BUILD_FLAGS) -o build/changelog-linux-amd64 cmd/main.go
	zip -r build/changelog-linux-amd64.zip build/changelog-linux-amd64
	rm build/changelog-linux-amd64

changelog-linux-arm64:
	GOOS=linux GOARCH=arm64 go build $(BUILD_FLAGS) -o build/changelog-linux-arm64 cmd/main.go
	zip -r build/changelog-linux-arm64.zip build/changelog-linux-arm64
	rm build/changelog-linux-arm64

changelog-windows-amd64:
	GOOS=windows GOARCH=amd64 go build $(BUILD_FLAGS) -o build/changelog-windows-amd64.exe cmd/main.go
	zip -r build/changelog-windows-amd64 build/changelog-windows-amd64.exe
	rm build/changelog-windows-amd64.exe

release: changelog-darwin-amd64 changelog-darwin-arm64 changelog-linux-amd64 changelog-linux-arm64 changelog-windows-amd64
