package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"

	"xgofind/util"
)

func main() {
	util.PrintBanner()
	// Sets the program parameters
	port := flag.String("port", "8080", "Port for the proxy to listen")
	destination := flag.String("dest", "http://example.com", "Destination URL where the proxy will redirect requests")
	help := flag.Bool("help", false, "Displays program help")
	logRequests := flag.Bool("log", false, "Enables logging of intercepted requests")
	flag.Parse()

	// Displays help and exits if the --help parameter is passed
	if *help {
		fmt.Println("Use:")
		fmt.Println("  xgofind [Options]")
		fmt.Println("Parameters:")
		flag.PrintDefaults()
		os.Exit(0)
	}

	// Validates the destination URL
	targetURL, err := url.Parse(*destination)
	if err != nil {
		log.Fatalf("Error parsing destination URL: %v", err)
	}

	// Displays information about the current configuration
	fmt.Printf("Starting the proxy at http://localhost:%s\n", *port)
	fmt.Printf("Redirecting Requests to %s\n", *destination)

	// Configures the proxy handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if *logRequests {
			// pkg.InitLogger()
			fmt.Printf("Intercepted: %s %s\n", r.Method, r.URL.Path)
		}

		// Creates the reverse proxy
		proxy := httputil.NewSingleHostReverseProxy(targetURL)

		// Intercepta a resposta (se necessário)
		proxy.ModifyResponse = func(resp *http.Response) error {
			if *logRequests {
				fmt.Printf("Response with status: %d\n", resp.StatusCode)
			}
			return nil
		}

		// Sends the request to the destination
		proxy.ServeHTTP(w, r)
	})

	// Starts the HTTP server
	err = http.ListenAndServe(":"+*port, nil)
	if err != nil {
		log.Fatalf("Error starting the proxy: %v", err)
	}
}
