package handlers

import (
	"fmt"
	"gotrix/lib/logger"
	"io"
	"mime"
	"net/http"
	"path/filepath"
)

func StaticFile(l logger.Logger, statusCode int, fs http.FileSystem, path string) http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		fh, err := fs.Open(path)
		if err != nil {
			panic(fmt.Sprintf("couldn't open static file %q: %s", path, err))
		}
		defer fh.Close()

		ext := filepath.Ext(path)
		rw.Header().Set("Content-Type", mime.TypeByExtension(ext))

		rw.WriteHeader(statusCode)
		_, err = io.Copy(rw, fh)
		if err != nil {
			panic(fmt.Sprintf("couldn't copy static file %q: %s", path, err))
		}

		l.Printf("sending static file %s with status code %d", path, statusCode)
	}
}
