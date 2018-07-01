hangulize.js: main.go ../hangulize/*.go ../hangulize/hangulize-packr.go
	go get github.com/gopherjs/gopherjs
	gopherjs build -o hangulize.js

../hangulize/hangulize-packr.go:
	go generate github.com/hangulize/hangulize

.PHONY: clean
clean:
	rm -f hangulize.js hangulize.js.map
	(cd ../hangulize && packr clean)
