
build:
	go build -o "LibertyFixer.exe" -ldflags "-H windowsgui -w -s" ./cmd/main.go