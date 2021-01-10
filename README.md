This application is written in golang.

You will need golang >=1.12 and redis installed. To install golang, visit https://golang.org/ and follow the installation instructions for your operating system.

To install redis:
```
wget http://download.redis.io/redis-stable.tar.gz
tar xvzf redis-stable.tar.gz
cd redis-stable
make && make install
```

To start redis and the web server, run

```
> redis-server & 
> go run main.go urlinfo_controller.go urlinfo_db.go
```

You with the web server running, you can make requests to http://0.0.0.0:8080/

To request blocklist status for a url, request http://0.0.0.0:8080/urlinfo/1/:hostnameAndPort/:whatever/:path?some=query

Responses are Json with format:

```
{\"blocklisted\":true}
```

or

```
{\"blocklisted\":false}
```

The blocklist is loaded at start-up from a file. Default file is blocklist.txt. You can specify your own blocklist file with an environment variable:

```
> BLOCKLIST_FILENAME=blocklist.txt go run main.go urlinfo_controller.go urlinfo_db.go
```

To test, please specify the testing blocklist and run:

```
> redis-server & 
> BLOCKLIST_FILENAME=blocklist-test.txt go test
```

