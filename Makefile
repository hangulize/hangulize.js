hangulize.js: ../hangulize/hangulize-packr.go
	go get github.com/gopherjs/gopherjs
	gopherjs build -o hangulize.js

../hangulize/hangulize-packr.go:
	go get github.com/gobuffalo/packr/...
	packr -i ../hangulize

.PHONY: clean
clean:
	rm -f hangulize.js hangulize.js.map
	(cd ../hangulize && packr clean)