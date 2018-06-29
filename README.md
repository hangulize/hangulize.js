# hangulize.js

The JavaScript build of Hangulize, compiled by
[GopherJS](https://github.com/gopherjs/gopherjs).

## Compiling hangulize.js

You will get `hangulize.js` by:

```console
$ make
```

You can try it with a GopherJS test server:

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
