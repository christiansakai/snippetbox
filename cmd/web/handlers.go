package main

import (
  "net/http"
  "strconv"
  "fmt"
  "html/template"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
  if r.URL.Path != "/" {
    http.NotFound(w, r)
    return
  }

  files := []string{
    "./ui/html/home.page.tmpl",
    "./ui/html/base.layout.tmpl",
    "./ui/html/footer.partial.tmpl",
  }

  ts, err := template.ParseFiles(files...)
  if err != nil {
    app.serverError(w, err)
    return
  }

  err = ts.Execute(w, nil)
  if err != nil {
    app.serverError(w, err)
  }
}

func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
  id, err := strconv.Atoi(r.URL.Query().Get("id"))
  if err != nil || id < 1 {
    app.notFound(w)
    return
  }

  fmt.Fprint(w, "Display a specific snippet with ID %d...", id)
}

func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
  if r.Method != http.MethodPost {
    w.Header().Set("Allow", http.MethodPost)

    app.clientError(w, http.StatusMethodNotAllowed)
    return
  }

  w.Write([]byte("Create a new snippet..."))
}
