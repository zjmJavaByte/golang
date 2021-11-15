package main

import (
	"io/ioutil"
	"net/http"
	"regexp"
	"text/template"
)

const nPath = len("/view/")

var titleValidator = regexp.MustCompile("^[a-zA-Z0-9]+$")

var templates = make(map[string]*template.Template)

var err error

type Page struct {
	Title string
	Body  []byte
}

func init() {
	for _, temp := range []string{"view"} {
		templates[temp] = template.Must(template.ParseFiles(temp + ".html"))
	}
}

func main() {

	http.Handle("/view", makeHandler(viewHandler))
	http.ListenAndServe(":8080", nil)

}

func makeHandler(f func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		title := request.URL.Path[:nPath]
		if !titleValidator.MatchString(title) {
			http.NotFound(writer, request)
		}
		f(writer, request, title)
	}
}

func viewHandler(w http.ResponseWriter, re *http.Request, title string) {
	p, err := load(title)
	if err != nil {
		http.Redirect(w, re, "/edit/", http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)

}

func load(title string) (*Page, error) {
	fileName := title + ".txt"
	body, err := ioutil.ReadFile(fileName)
	if err != nil {

		return nil, err

	}
	return &Page{title, body}, nil
}

func renderTemplate(w http.ResponseWriter, teml string, p *Page) {
	err := templates[teml].Execute(w, p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
