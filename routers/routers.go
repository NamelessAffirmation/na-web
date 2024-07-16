package routers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/namelessaffirmation/na-web/config"
	"github.com/namelessaffirmation/na-web/database"
)

type PageData struct {
	Title   string
	Content string
}

func LoadRouters() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	workDir, _ := os.Getwd()
	filesDir := http.Dir(filepath.Join(workDir, "public"))
	FileServer(r, "/public", filesDir)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		// Parse the HTML template
		tmpl, err := template.ParseFiles("views/index.html")
		if err != nil {
			http.Error(w, "Unable to load template", http.StatusInternalServerError)
			return
		}

		// Data to be passed to the template
		data := PageData{
			Title:   "NamelessAffirmation",
			Content: "This is an example of serving HTML templates in Go.",
		}

		// Execute the template with data
		if err := tmpl.Execute(w, data); err != nil {
			http.Error(w, "Unable to execute template", http.StatusInternalServerError)
		}
	})

	r.Get("/test", func(w http.ResponseWriter, r *http.Request) {
		database.CreateStagedMessage("This is a test message!")
		w.Write([]byte(fmt.Sprintf("welcome %s", config.AppConfiguration.DBConnectionString)))
	})
	return r
}

func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", http.StatusMovedPermanently).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
		fs.ServeHTTP(w, r)
	})
}
