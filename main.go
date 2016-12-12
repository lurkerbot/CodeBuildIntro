package main

import (
    "fmt"
    "net/http"
    "net"
    "log"
)

type SimpleHTTPHandler struct{}

func (*SimpleHTTPHandler) ServeHTTP(httpWriter http.ResponseWriter, httpRequest *http.Request) {
    fmt.Fprintf(httpWriter, "Greetings CodeDeploy :), path: %s", httpRequest.URL.Path[1:])
    log.Print("serving")
}

/**
 * Simple port 1234 server to allow demonstration of minimal CodeBuild/CodePipeline
 * CI deployment
 */
func main() {
    server := &http.Server{Handler: &SimpleHTTPHandler{}}

    listener, error := net.Listen("tcp4", ":1234")
    
    if error != nil {
        log.Fatal(error)
    }

    error = server.Serve(listener)

    if error != nil {
        log.Fatal(error)
    }
}	
