Urlinfo is a web service that returns blocklist status for URLs.

To see if a URL is blocklisted, request the following endpoint, specifying hostname and optionally, a port, path and query string:
```
http://0.0.0.0:8080/urlinfo/1/:hostnameAndPort/:path/:maybeMorePath?some=OptionalQueryString
```

Successful json responses are of the form HTTP 200 with body:

```
{\"blocklisted\":true}
```

or 

```
{\"blocklisted\":false}
```

Requests that do not adhere to the above format will return an HTTP 404.

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

The blocklist is loaded at start-up from a file. The default file is `blocklist.txt`. You can specify your own blocklist file with an environment variable:

```
> BLOCKLIST_FILENAME=blocklist.txt go run main.go urlinfo_controller.go urlinfo_db.go
```

To test, please specify the testing blocklist by running:

```
> redis-server & 
> BLOCKLIST_FILENAME=blocklist-test.txt go test
```

