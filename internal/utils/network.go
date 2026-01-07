package utils

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

// ParseURL parses a URL and extracts host, port, and path
func ParseURL(rawURL string) (host string, port string, path string, useHTTPS bool, err error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return "", "", "", false, fmt.Errorf("invalid URL: %w", err)
	}

	// Extract host (remove port if present)
	host = u.Hostname()
	if host == "" {
		return "", "", "", false, fmt.Errorf("invalid URL: missing host")
	}

	// Determine port and scheme
	useHTTPS = u.Scheme == "https"
	if u.Port() != "" {
		port = u.Port()
	} else if useHTTPS {
		port = "443"
	} else {
		port = "80"
	}

	// Extract path
	path = u.Path
	if path == "" {
		path = "/"
	}

	return host, port, path, useHTTPS, nil
}

// FormatAddress formats host and port into an address string
func FormatAddress(host, port string) string {
	return host + ":" + port
}

// ValidatePort checks if a port number is valid
func ValidatePort(portStr string) error {
	port, err := strconv.Atoi(portStr)
	if err != nil {
		return fmt.Errorf("invalid port number: %s", portStr)
	}
	if port < 1 || port > 65535 {
		return fmt.Errorf("port number out of range: %d", port)
	}
	return nil
}

// IsValidURL checks if a string is a valid URL
func IsValidURL(rawURL string) bool {
	u, err := url.Parse(rawURL)
	if err != nil {
		return false
	}
	return u.Scheme != "" && u.Host != ""
}

// GetURLParameterSeparator returns the appropriate parameter separator based on URL
func GetURLParameterSeparator(path string) string {
	if strings.Contains(path, "?") {
		return "&"
	}
	return "?"
}