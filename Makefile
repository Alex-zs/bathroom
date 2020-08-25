build:
	go build -ldflags '-w -s' -o bathroom-osx && upx bathroom-osx
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags '-w -s' -o bathroom-linux && upx bathroom-linux
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags '-w -s' -o bathroom-win.exe && upx bathroom-win.exe
