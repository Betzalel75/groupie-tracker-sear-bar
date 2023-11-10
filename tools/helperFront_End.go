package tools

import (
	"bytes"
	"html/template"
	"net/http"
	"path/filepath"
)

func RenderTemplate(w http.ResponseWriter, tmplName string, tmplDir string, data interface{}) {

	templateCache, err := createTemplateCache(tmplDir)

	if err != nil {
		panic(err)
	}
	tmpl, ok := templateCache[tmplName+".page.tmpl"]

	if !ok {
		ErrorHandler(w, http.StatusInternalServerError, "Internal Server Error", "Sorry, an error occurred while retrieving the artist.")
		return
	}

	buffer := new(bytes.Buffer)
	tmpl.Execute(buffer, data)
	buffer.WriteTo(w)
}

func createTemplateCache(tmplDir string) (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}
	pages, err := filepath.Glob("./template/pages/" + tmplDir + "/*.page.tmpl")
	if err != nil {
		return cache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		tmpl := template.Must(template.ParseFiles(page))

		layouts, err := filepath.Glob("./template/layout/*.layout.tmpl")
		if err != nil {
			return cache, err
		}
		if len(layouts) > 0 {
			tmpl.ParseGlob("./template/layout/*.layout.tmpl")
		}
		cache[name] = tmpl
	}
	return cache, nil
}
