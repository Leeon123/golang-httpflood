package config

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bonbonbon1414/golang-httpflood/internal/utils"
)

// Config holds the configuration for the flood operation
type Config struct {
	URL             string
	Host            string
	Port            string
	Path            string
	Mode            string // "get" or "post"
	Threads         int
	Duration        int // in seconds
	UseHTTPS        bool
	CustomHeaderFile string
	UseCustomHeaders bool
	RequestsPerConnection int
	RandomChars string
}

// NewConfig creates a new configuration with defaults
func NewConfig() *Config {
	return &Config{
		Mode:                  "get",
		Threads:               10,
		Duration:              60,
		UseHTTPS:              false,
		CustomHeaderFile:      "nil",
		UseCustomHeaders:      false,
		RequestsPerConnection: 100,
		RandomChars:           "asdfghjklqwertyuiopzxcvbnmASDFGHJKLQWERTYUIOPZXCVBNM",
	}
}

// ParseFromArgs parses configuration from command-line arguments
func (c *Config) ParseFromArgs(args []string) error {
	if len(args) != 6 {
		return fmt.Errorf("expected 5 arguments, got %d", len(args)-1)
	}

	// Parse URL
	c.URL = args[1]
	host, port, path, useHTTPS, err := utils.ParseURL(c.URL)
	if err != nil {
		return fmt.Errorf("invalid URL: %w", err)
	}
	c.Host = host
	c.Port = port
	c.Path = path
	c.UseHTTPS = useHTTPS

	// Parse threads
	threads, err := strconv.Atoi(args[2])
	if err != nil {
		return fmt.Errorf("invalid threads value: %s", args[2])
	}
	if threads < 1 {
		return fmt.Errorf("threads must be at least 1, got %d", threads)
	}
	c.Threads = threads

	// Parse mode
	mode := strings.ToLower(args[3])
	if mode != "get" && mode != "post" {
		return fmt.Errorf("invalid mode: %s, must be 'get' or 'post'", mode)
	}
	c.Mode = mode

	// Parse duration
	duration, err := strconv.Atoi(args[4])
	if err != nil {
		return fmt.Errorf("invalid duration value: %s", args[4])
	}
	if duration < 1 {
		return fmt.Errorf("duration must be at least 1 second, got %d", duration)
	}
	c.Duration = duration

	// Parse header file
	c.CustomHeaderFile = args[5]
	c.UseCustomHeaders = (c.CustomHeaderFile != "nil")

	return nil
}

// Validate validates the configuration
func (c *Config) Validate() error {
	if c.URL == "" {
		return fmt.Errorf("URL is required")
	}
	if !utils.IsValidURL(c.URL) {
		return fmt.Errorf("invalid URL format: %s", c.URL)
	}
	if c.Host == "" {
		return fmt.Errorf("host is required")
	}
	if c.Port == "" {
		return fmt.Errorf("port is required")
	}
	if c.Mode != "get" && c.Mode != "post" {
		return fmt.Errorf("mode must be 'get' or 'post'")
	}
	if c.Threads < 1 {
		return fmt.Errorf("threads must be at least 1")
	}
	if c.Duration < 1 {
		return fmt.Errorf("duration must be at least 1 second")
	}
	if c.UseCustomHeaders && c.CustomHeaderFile == "nil" {
		return fmt.Errorf("custom header file specified but file is 'nil'")
	}
	return nil
}

// GetAddress returns the formatted address string (host:port)
func (c *Config) GetAddress() string {
	return utils.FormatAddress(c.Host, c.Port)
}

// GetURLParameterSeparator returns the appropriate parameter separator
func (c *Config) GetURLParameterSeparator() string {
	return utils.GetURLParameterSeparator(c.Path)
}

// String returns a string representation of the config
func (c *Config) String() string {
	return fmt.Sprintf(
		"URL: %s\nHost: %s\nPort: %s\nPath: %s\nMode: %s\nThreads: %d\nDuration: %ds\nHTTPS: %t\nCustom Headers: %t",
		c.URL, c.Host, c.Port, c.Path, c.Mode, c.Threads, c.Duration, c.UseHTTPS, c.UseCustomHeaders,
	)
}