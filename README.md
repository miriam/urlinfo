This application is written in golang.

1. You will need golang 1.12 or greater installed. To install golang, visit https://golang.org/ and follow the installation instructions for your operating system.

2. This application uses the Gin framework. To install, run:

```
$ go get -u github.com/gin-gonic/gin
```

3. To start the web server, run

```
go run main.go
```

You can specify your own blocklist file with an environment variable:

```
BLOCKLIST_FILENAME=blocklist.txt go run main.go
```

4. To test, run:

```
BLOCKLIST_FILENAME=blocklist-test.txt go test
```

