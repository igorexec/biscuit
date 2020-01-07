windows:
	GOOS=windows GOARCH=amd64 go build -o dist/biscuit.exe app/main.go

macos:
	GOOS=darwin GOARCH=amd64 go build -o dist/biscuit app/main.go