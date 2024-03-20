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
    var filename string

    title := reqest.URL.Path[len("/"):]

    if strings.Contains(title, "scripts") {
        // is scripts in the path
        title := subStringStriper(title, "scripts") // TODO testing
        filename := "./" + title
        fmt.Println(filename)

    } else if strings.Contains(title, "styles") {
        // is styles in path
        title := subStringStriper(title,  "styles") // TODO testing
        filename := "./" + title
        fmt.Println(filename)

    } else {
        filename := "./" + title + ".html" // Assuming templates directory
        fmt.Println(filename)
    }

    // Serve the file using http.ServeFile
    http.ServeFile(respon, reqest , filename)
}

// runs all testing functions created
func testRunner() {
    test_SubStringStriper()
}

// NOTE.....
func main() {
    testRunner()

    fmt.Println("Running...")
    http.HandleFunc("/", Handler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
