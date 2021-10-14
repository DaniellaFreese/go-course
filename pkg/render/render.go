package render

import (
	"fmt"
	"net/http"
	"path/filepath"
	"text/template"
)

var functions = template.FuncMap{}

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	_, err := RenderTemplateTest(w)
	if err != nil {
		fmt.Println("error getting template cache:", err)
	}

	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err = parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template:", err)
		return
	}
}

func RenderTemplateTest(w http.ResponseWriter) (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.go.html")

	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		fmt.Println("currently building a template")
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)

		if err != nil {
			fmt.Println("error on creating template")
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.go.html")
		if err != nil {
			fmt.Println("error on getting matches")
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.go.html")
			if err != nil {
				fmt.Println("error on getting the layout")
				return myCache, err
			}
		}

		myCache[name] = ts
	}
	return myCache, nil
}
