package beamcmd

import (
	rest "graphdbcli/cmd/beamcmd/rest"
	"graphdbcli/internal/tool_configurations/logging"
	"io"
	"log"
	"net/http"
	"net/url"
)

// StartBeamProxyServer starts an HTTP server on listenPort that proxies requests to targetAddr.
// targetAddr should be in the form protocol://host:port (e.g., http://localhost:7200)
// listenPort should be just the port number (e.g., "8080").
// You can add request/response manipulation in the handler as needed.
func StartBeamProxyServer(targetAddr, listenPort string) error {
	targetURL, err := url.Parse(targetAddr)
	if err != nil {
		return err
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		// Build the proxied request URL
		proxyURL := *targetURL
		proxyURL.Path = r.URL.Path
		proxyURL.RawQuery = r.URL.RawQuery

		// Create new request to target
		req, err := http.NewRequest(r.Method, proxyURL.String(), r.Body)
		if err != nil {
			http.Error(w, "Bad Gateway", http.StatusBadGateway)
			return
		}
		// Copy headers
		for k, v := range r.Header {
			req.Header[k] = v
		}

		rest.ChangeRequest(req)

		// --- Place for request manipulation ---
		// Example: log.Printf("Proxying %s to %s", r.URL.Path, proxyURL.String())

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			logging.LOGGER.Error("Failed to connect to the end server")
			http.Error(w, "Bad Gateway", http.StatusBadGateway)
			return
		}
		defer resp.Body.Close()

		// Copy response headers
		for k, v := range resp.Header {
			w.Header()[k] = v
		}
		w.WriteHeader(resp.StatusCode)

		// --- Place for response manipulation ---
		io.Copy(w, resp.Body)
	}

	http.HandleFunc("/", handler)
	log.Printf("Proxy server listening on :%s, forwarding to %s", listenPort, targetAddr)
	return http.ListenAndServe(":"+listenPort, nil)
}
