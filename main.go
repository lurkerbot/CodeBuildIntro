package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    // fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
    fmt.Fprintf(w, "Greetings CodeDeploy :)")
}

/**
 * Simple port 8080 server to allow demonstration of minimal CodeBuild/CodePipeline
 * CI deployment
 */
func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":1234", nil)
}
