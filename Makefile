all:
	env GOOS=darwin GOARCH=amd64 go build -o build/wp-audit_macos
	env GOOS=linux GOARCH=amd64 go build -o build/wp-audit_linux
	env GOOS=windows GOARCH=amd64 go build -o build/wp-audit_windows
