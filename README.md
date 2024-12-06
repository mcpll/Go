## Exercice 1

countwords


## Exercice 2


Write a program which is a URL shortner, that can shorten URL, and performs HTTP
redirects from the short URL to the original ones.

```
./server host:port
```

1. Endpoint to add URL to the server

Example: If i add http://develer.com to the server, the server returns a short
url, for example `abc012`.

Goal: 

```
curl -X POST http://localhost:8080/shorten -d '{ "url": "http://www.google.com" }'
```

should return the short URL.


`http//localhost:8080/abc012`

documentation: net/http.Server


2. Redirect short URL to the original (long) url

http//localhost:8080/abc012 -> redirects to https://google.com

3. Non-existing short URLs returns HTTP 404

http//localhost:8080/pippo -> HTTP 404
