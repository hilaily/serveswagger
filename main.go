package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/hilaily/kit/pathx"
	"github.com/pkg/browser"
)

var (
	//go:embed dist/*
	dist    embed.FS
	port    = "8123"
	baseURL = "http://127.0.0.1:" + port
)

func main() {
	log.SetFlags(0)
	var file string
	if len(os.Args) > 1 {
		file = os.Args[1]
	}
	for _, v := range []string{"./swagger.json", "./swagger.yaml"} {
		if pathx.IsExist(v) {
			file = v
			break
		}
	}
	if file == "" {
		log.Printf("you should use it like 'serveswagger /opt/swagger.json'")
		return
	}

	sub, _ := fs.Sub(dist, "dist")

	files := http.FS(sub)
	u := file
	if !strings.HasPrefix(file, "http") {
		files = &mfs{
			FileSystem: http.FS(sub),
			path:       file,
		}
		u = baseURL + "/default.json"
	}
	http.Handle("/", http.FileServer(files))
	browser.OpenURL(baseURL + "?file=" + u)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

type mfs struct {
	http.FileSystem
	path string
}

func (f *mfs) Open(name string) (http.File, error) {
	if name == "/default.json" {
		return os.OpenFile(f.path, os.O_RDONLY, 0777)
	}
	return f.FileSystem.Open(name)
}
