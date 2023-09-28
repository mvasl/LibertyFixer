@echo off
go version
SET GOARCH=amd64
SET GOOS=windows
go build -o "LibertyFixer.exe" -ldflags "-H windowsgui -w -s" .\cmd\main.go
