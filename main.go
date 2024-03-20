package main 

import (
    "fmt"
    "log"
    "net/http"
)

func Handler(respon http.ResponseWriter, reqest *http.Request) {
    title := reqest.URL.Path[len("/"):]
    fmt.Println(title)
    filename := "./" + title + ".html" // Assuming templates directory
    fmt.Println(filename)

    // Serve the file using http.ServeFile
    http.ServeFile(respon, reqest , filename)
}


// NOTE.....
func main() {
    fmt.Println("Running...")
    http.HandleFunc("/", Handler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
