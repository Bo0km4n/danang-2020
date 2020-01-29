.PHONY: release
release:
	GOOS=linux GOARCH=amd64 go build -o bin/danang_linux_amd64
	GOOS=darwin GOARCH=amd64 go build -o bin/danang_darwin_amd64
	GOOS=windows GOARCH=amd64 go build -o bin/danang_windows_amd64