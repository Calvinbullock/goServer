package main 

import (
    "fmt"
    "os"
    "log"
    "net/http"
)

/* ================================ \\
            Page Struct
\\ ================================ */

type Page struct {
    Title string
    Body  []byte
}

// saves pages to file systems
func (p *Page) save() error {
    filename := p.Title + ".txt"
    return os.WriteFile(filename, p.Body, 0600)
}

/* ================================ \\
              Main File
\\ ================================ */

// Loads Page files and loads them into memory
func loadPage(title string) (*Page, error) {
    filename := title + ".txt"
    body, err := os.ReadFile(filename)
    if err != nil {
        return nil, err
    }
    return &Page{Title: title, Body: body}, nil
}

// loads the page that matches the url request
func viewHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/view/"):]
    p, _ := loadPage(title)
    fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

// NOTE.....
func main() {
    fmt.Println("Running...")
    http.HandleFunc("/view/", viewHandler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
