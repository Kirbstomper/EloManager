package handlers

import (
	"html/template"
	"net/http"
	"path"
)

const indexFilename = "index.html"
const viewsRootDir = "./views"

func (s defaultServer) serveIndexPage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		indexTemplate := template.Must(template.New(indexFilename).
			ParseFiles(path.Join(viewsRootDir, indexFilename)))
		if err := indexTemplate.ExecuteTemplate(w, indexFilename, nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	}
}
