package webserver

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
)

const keyServerAddr = "serverAddr"

func writeJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")

	// Encode the response struct to JSON and write it to the response
	if err := json.NewEncoder(w).Encode(data); err != nil {
		// If there is an error, log it and set the HTTP status to 500
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func healthz(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	fmt.Printf("%s: got /healthz request\n", ctx.Value(keyServerAddr))
	io.WriteString(w, "This is my website!\n")
}

func ready(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	fmt.Printf("%s: got /ready request\n", ctx.Value(keyServerAddr))
	io.WriteString(w, "This is my website!\n")
}

// Define a struct to hold the response data
type Response struct {
	Message   string `json:"message"`
	Status    int    `json:"status"`
	ReceiptId string `json:"receipt_id"`
}

// Receipts
// will receive a POST request with a JSON payload containing a receipt
// will respond with a JSON payload containing the receipt ID
func createReceipt(w http.ResponseWriter, r *http.Request) {
	// log the request
	fmt.Printf("Received request, method: %s uri %s \n", r.Method, r.RequestURI)

	// Create an instance of the Response struct with some data
	response := Response{
		Message:   "Hello, World!",
		Status:    200,
		ReceiptId: "12345",
	}

	// write json response
	writeJSON(w, response)
}

type Receipt struct {
	ID string `json:"id"`
}

func getReceipt(w http.ResponseWriter, r *http.Request) {
	// log the request
	fmt.Printf("Received request, method: %s uri %s \n", r.Method, r.RequestURI)

	re := regexp.MustCompile(`^/receipts/(\d+)$`)
	matches := re.FindStringSubmatch(r.URL.Path)

	if len(matches) == 0 {
		http.NotFound(w, r)
		return
	}

	id := matches[1]
	// Create an instance of the Response struct with some data
	response := Receipt{
		ID: id,
	}

	// write json response
	writeJSON(w, response)
}

func updateReceipt(w http.ResponseWriter, r *http.Request) {
	// log the request
	fmt.Printf("Received request, method: %s uri %s \n", r.Method, r.RequestURI)

	re := regexp.MustCompile(`^/receipts/(\d+)$`)
	matches := re.FindStringSubmatch(r.URL.Path)

	if len(matches) == 0 {
		http.NotFound(w, r)
		return
	}

	id := matches[1]

	response := Receipt{
		ID: id,
	}

	// write json response
	writeJSON(w, response)
}

func deleteReceipt(w http.ResponseWriter, r *http.Request) {
	// log the request
	fmt.Printf("Received request, method: %s uri %s \n", r.Method, r.RequestURI)

	re := regexp.MustCompile(`^/receipts/(\d+)$`)
	matches := re.FindStringSubmatch(r.URL.Path)

	if len(matches) == 0 {
		http.NotFound(w, r)
		return
	}

	id := matches[1]

	response := Receipt{
		ID: id,
	}

	// write json response
	writeJSON(w, response)
}
