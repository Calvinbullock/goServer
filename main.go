package main 

import (
    "fmt"
    "log"
    "strings"
    "net/http"
)

// strips everything that occurs before the first occurrence of "target" from "text"
//      returns text with out the strip content
func subStringStriper(text string, target string) string {
    // Find the index of the first occurrence of target
    index := strings.Index(text, target)

    // If "scripts/" is found, extract the substring after it
    if index != -1 {
        return text[index:]
    } else {
        // No "scripts/" found, result remains the original string
        return text
    }
}

// test function for subStringStriper
func test_SubStringStriper() {
    // TODO testing
    ans := "pop"
    text := "helppop"
    retu := subStringStriper(text, ans)
    if ans != retu {
        log.Fatal("test 1 failed", ans)
    }

    // TODO add more test cases
}

// Recives request and sends back apropreate response
func Handler(respon http.ResponseWriter, reqest *http.Request) {
    // quick validation
    if reqest.URL.Path == "/" {
        http.NotFound(respon, reqest)
        return
    }

    var filename string

    title := reqest.URL.Path[len("/"):]
    // Remove trailing "/" 
    title = strings.TrimSuffix(title, "/")
    

    if strings.Contains(title, "scripts") {
        // is scripts in the path
        title := subStringStriper(title, "scripts")
        filename := "./" + title
        fmt.Println("js: ", filename)

    } else if strings.Contains(title, "styles") {
        // is styles in path
        title := subStringStriper(title,  "styles")
        filename := "./" + title
        fmt.Println("css: ", filename)

    } else {
        filename := "./" + title + ".html"
        fmt.Println("html: ", filename)
    }

    // Serve the file using http.ServeFile
    http.ServeFile(respon, reqest , filename)
}

// runs all testing functions created
func testRunner() {
    test_SubStringStriper()
    fmt.Println("testing done")
}

// NOTE.....
func main() {
    testRunner()

    fmt.Println("Running...")
    http.HandleFunc("/", Handler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
