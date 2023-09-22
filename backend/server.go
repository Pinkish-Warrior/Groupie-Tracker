package backend

import (
	"html/template"
	"io"
	"net/http"
)

type PageData struct {
	Myartists []Artists
	ErrorNum  int
	ErrorText string
}

var (
	d                  = PageData{}
	tpl                *template.Template
	Error400, Error500 = false, false
)

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html")) //.gohtml is a standard html in go. both is ok.
}

func MainpageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		d.ErrorNum = 405
		d.ErrorText = "Method Not Allowed"
		errorHandler(w, r, &d)
		return
	}
	if r.URL.Path != "/" {
		d.ErrorNum = 404
		d.ErrorText = "Page Not Found"
		errorHandler(w, r, &d)
		return
	} else {
		d.Myartists = Myartists
		//Not a good way but can whrite data into some blackwhole to avoid it show on html page.
		err := tpl.ExecuteTemplate(io.Discard, "index.html", &d)
		if err != nil {
			d.ErrorNum = 500
			d.ErrorText = "Internal Server Error"
			errorHandler(w, r, &d)
			return
		} else {
			tpl.ExecuteTemplate(w, "index.html", &d)
		}
	}
}

// Just temeraly show error.html
func errorHandler(w http.ResponseWriter, r *http.Request, d *PageData) {
	w.WriteHeader(d.ErrorNum)
	tpl.ExecuteTemplate(w, "error.html", &d)
}
