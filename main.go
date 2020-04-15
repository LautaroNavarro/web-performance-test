package main

import (
	"net/http"
	"io"
	"encoding/json"
	"fmt"
)

func readBody(body io.ReadCloser, message *bodyMessage, l *logger) error {
	l.Printf("[DEBUG] Parsing request\n")

    buffer := make([]byte, 32768)
    read_len, _ := body.Read(buffer)
    return json.Unmarshal(buffer[:read_len], &message)
}

func performanceTestHandler (w http.ResponseWriter, req *http.Request) {

	l := getLogger()
	l.Printf("Request received.\n")

    message := bodyMessage{}

    error := readBody(req.Body, &message, l)

	l.Printf("[DEBUG] Request body: %+v\n", message)
    if error != nil {
    	l.Printf("Invalid request\n")
    	fmt.Fprintf(w, "{\"error\": \"There are errors with the provided body\"}\n")
    }

    results := executePerformanceTest(message.Website, message.TimesToHit, l)

    l.Println("[DEBUG] Performance test finsihed succesfully")

    fmt.Fprintf(w, "{\"elapsed_time\": \"%v %v\"}\n", results, "ms")

}

func main(){

	port := ":8090"

    l := getLogger()
    l.Printf("Listening at port %v\n", port)

    http.HandleFunc("/performance-test", performanceTestHandler)

    http.ListenAndServe(port, nil)
}
