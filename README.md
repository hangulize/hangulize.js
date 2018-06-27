# js

For JavaScript build of Hangulize.

## Installing Build Tools

```console
$ go get -u github.com/gopherjs/gopherjs
$ go get -u github.com/gobuffalo/packr/...
```

## Building

```console
$ go get
$ packr -i "$GOPATH/src/github.com/hangulize/hangulize"
$ gopherjs build -o hangulize.js
```

## Trying It

```console
$ gopherjs serve --http ":8080"
```

Now connect to "localhost:8080" in your browser than open a JavaScript debug
console.

```js
 > hangulize("ita", "gloria");
<- "글로리아"
 > hangulize("aze", "Rəşid Behbudov");
<- "레시트 베흐부도프"
```
