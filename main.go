package main

import (
	"fmt"
	"github.com/google/uuid"
	"log"
	"net/http"
	"net/url"
	"text/template"
)

// LoadForm renders the HTML page
func LoadForm(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/index.html")
	tmpl.Execute(w, nil)
}

// HandleFormSubmission processes the form data, constructs the URL, and redirects
func HandleFormSubmission(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Generate default values
		workItemID := uuid.New().String() // Automatically generate a UUID for Work Item ID
		branchCode := "800"               // Default Branch Code
		segmentID := "CDS"                // Default Segment ID
		lga := "NA"                       // Default LGA
		state := "LA"                     // Default State
		requestID := ""                   // Default Request ID as empty string
		rmCode := ""                      // Default RM Code as empty string
		createdBy := "SME_APP_WEB"        // Default Created By

		// Collect the required form data
		customerName := r.FormValue("customerName")
		address := r.FormValue("address")
		landmark := r.FormValue("landmark")
		latitude := r.FormValue("latitude")
		longitude := r.FormValue("longitude")
		phoneNumber := r.FormValue("phoneNumber")

		// Print out the values being submitted
		fmt.Println("Values being submitted:")
		fmt.Printf("createdBy - %s\n", createdBy)
		fmt.Printf("workItemID - %s\n", workItemID)
		fmt.Printf("branchCode - %s\n", branchCode)
		fmt.Printf("segmentID - %s\n", segmentID)
		fmt.Printf("lga - %s\n", lga)
		fmt.Printf("state - %s\n", state)
		fmt.Printf("requestID - %s\n", requestID)
		fmt.Printf("rmCode - %s\n", rmCode)
		fmt.Printf("customerName - %s\n", customerName)
		fmt.Printf("address - %s\n", address)
		fmt.Printf("landmark - %s\n", landmark)
		fmt.Printf("latitude - %s\n", latitude)
		fmt.Printf("longitude - %s\n", longitude)
		fmt.Printf("phoneNumber - %s\n", phoneNumber)
        fmt.Println()

		// Construct the URL with query parameters
		baseURL := "https://ecocomonoreact.azurewebsites.net/customer-details/"
		reqURL, err := url.Parse(baseURL)
		if err != nil {
			http.Error(w, "Invalid URL", http.StatusInternalServerError)
			return
		}

		// Add query parameters
		query := reqURL.Query()
		query.Set("workitemId", workItemID)
		query.Set("customerName", customerName)
		query.Set("branchCode", branchCode)
		query.Set("segmentId", segmentID)
		query.Set("address", address)
		query.Set("landmark", landmark)
		query.Set("state", state)
		query.Set("lga", lga)
		query.Set("createdBy", createdBy)
		query.Set("Latitude", latitude)
		query.Set("Longitude", longitude)
		query.Set("phoneNumber", phoneNumber)
		query.Set("requestId", requestID)
		query.Set("rmCode", rmCode)

		reqURL.RawQuery = query.Encode()

		// Redirect the user to the constructed URL
		http.Redirect(w, r, reqURL.String(), http.StatusFound)
	}
}

func main() {
	http.HandleFunc("/", LoadForm)
	http.HandleFunc("/submit", HandleFormSubmission)

	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
