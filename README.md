# fnsanitizer
Sanitize your file name to prevent Directory Traversing attack

## Install

```bash
go get -u github.com/codenoid/fnsanitizer
```

## Real World Usage

```go
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

import (
	"github.com/codenoid/fnsanitizer"
)

func ReceiveFile(w http.ResponseWriter, r *http.Request) {
	maxBytes := int64(10 << 20)
	r.ParseMultipartForm(maxBytes)
	defer r.MultipartForm.RemoveAll()

	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Bad Request", 400)
		return
	}
	defer file.Close()

	fmt.Println("[POST] Received File : " + handler.Filename)
	filename := fnsanitizer.Sanitize(handler.Filename)

	f, err := os.OpenFile("./upload/" + filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, "File creation failed", 500)
		return
	}
	defer f.Close()

	io.Copy(f, file)

	fmt.Fprintf(w, "Successfully upload file : %v \n", handler.Filename)
}

func main() {
	http.HandleFunc("/upload", ReceiveFile)
	http.ListenAndServe(":8080", nil)
}

```
