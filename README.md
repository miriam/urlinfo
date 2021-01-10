This application is written in golang.

1. You will need golang >=1.12 and redis installed. To install golang, visit https://golang.org/ and follow the installation instructions for your operating system.
1. To install redis:
```
wget http://download.redis.io/redis-stable.tar.gz
tar xvzf redis-stable.tar.gz
cd redis-stable
make && make install
```
1. This application uses the Gin framework. To install, run:

```
> go get -u github.com/gin-gonic/gin
```

3. To start the redis and web server, run

```
> redis-server & 
> go run main.go
```

You can specify your own blocklist file with an environment variable:

```
> BLOCKLIST_FILENAME=blocklist.txt go run main.go
```

4. To test, run:

```
> redis-server & 
> BLOCKLIST_FILENAME=blocklist-test.txt go test
```

