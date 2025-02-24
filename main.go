package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"

	"github.com/atotto/clipboard"
)

/**
*	Function to handle sanitizing URL's
*   @param: rawURL string | the URL to sanitize
*   @return: string, error | the sanitized URL and any errors
**/

func DefangUrl(rawURL string) (string, error) {
	hasScheme := strings.Contains(rawURL, "://")
	tempURL := rawURL
	if !hasScheme {
		tempURL = "temp://" + rawURL
	}

	parsedURL, err := url.Parse(tempURL)
	if err != nil {
		return "", fmt.Errorf("error parsing URL '%s': %w", rawURL, err)
	}

	if hasScheme {
		scheme := strings.ToLower(parsedURL.Scheme)
		switch scheme {
		case "https":
			parsedURL.Scheme = "hxxps"
		case "http":
			parsedURL.Scheme = "hxxp"
		}
	}

	host := strings.ToLower(parsedURL.Host)
	host = strings.ReplaceAll(host, ".", "[dot]")
	parsedURL.Host = host

	result := parsedURL.String()

	if !hasScheme {
		result = strings.TrimPrefix(result, "temp://")
	}

	return result, nil
}

/**
 * Main function
**/

func main() {

	fmt.Println(`🐍🐍 Thanks for using DURL! 🐍🐍`)

	fmt.Println(`🦷🦷 Pulling teeth... 🦷🦷`)

	// Parse flags, declare flag for multi-URL mode, which allows for multiple URL's to be handled at a time

	multiMode := flag.Bool("m", false, "Enable multi-URL mode")
	flag.Parse()

	if flag.NArg() == 0 {
		fmt.Println("⚠️⚠️ Please provide a URL ⚠️⚠️")
		os.Exit(1)
	}

	var sanitizedURLs []string

	for _, rawURL := range flag.Args() {
		sanitized, err := DefangUrl(rawURL)
		if err != nil {
			log.Printf("⚠️⚠️ Error sanitizing URL %s: %v ⚠️⚠️", rawURL, err)
			continue
		}
		sanitizedURLs = append(sanitizedURLs, sanitized)
	}

	if len(sanitizedURLs) == 0 {
		fmt.Println("⚠️⚠️ No valid URLs to process ⚠️⚠️")
		os.Exit(1)
	}

	// Join URLs if in multi-mode, otherwise take first URL
	var result string
	if *multiMode {
		result = strings.Join(sanitizedURLs, "\n")
	} else {
		result = sanitizedURLs[0]
	}

	// Write to clipboard
	if err := clipboard.WriteAll(result); err != nil {
		log.Fatalf("⚠️⚠️ Failed to write to clipboard: %v ⚠️⚠️", err)
	}

	fmt.Println("📎📎 Defanged URL(s) copied to clipboard: 📎📎")
	fmt.Println(result)
}
