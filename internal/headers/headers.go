package headers

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/bonbonbon1414/golang-httpflood/internal/utils"
)

// Manager handles HTTP header generation and management
type Manager struct {
	userAgentGenerator interface {
		Generate() string
	}
}

// NewManager creates a new header manager
func NewManager(uaGen interface{ Generate() string }) *Manager {
	return &Manager{
		userAgentGenerator: uaGen,
	}
}

// GetRandomAccept returns a random Accept header
func (m *Manager) GetRandomAccept() string {
	return utils.String(acceptHeaders)
}

// GetRandomReferer returns a random Referer URL
func (m *Manager) GetRandomReferer() string {
	return utils.String(referers)
}

// GenerateDefaultHeaders generates default randomized headers
func (m *Manager) GenerateDefaultHeaders() string {
	headers := ""
	headers += "Connection: Keep-Alive\r\n"
	headers += "Cache-Control: max-age=0\r\n"
	headers += "User-Agent: " + m.userAgentGenerator.Generate() + "\r\n"
	headers += m.GetRandomAccept()
	headers += m.GetRandomReferer() + "\r\n"
	return headers
}

// LoadCustomHeaders loads custom headers from a file
func (m *Manager) LoadCustomHeaders(filepath string) (string, error) {
	if filepath == "nil" {
		return "", fmt.Errorf("no custom header file specified")
	}

	file, err := os.Open(filepath)
	if err != nil {
		return "", fmt.Errorf("failed to open header file: %w", err)
	}
	defer file.Close()

	var headers string
	scanner := bufio.NewReader(file)
	for {
		line, isPrefix, err := scanner.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			return "", fmt.Errorf("error reading header file: %w", err)
		}
		if !isPrefix {
			headers += string(line) + "\r\n"
		}
	}

	return headers, nil
}

// LoadPostData loads POST data from a file
func (m *Manager) LoadPostData(filepath string) (string, error) {
	if filepath == "nil" {
		return "f", nil // Default data
	}

	file, err := os.Open(filepath)
	if err != nil {
		return "", fmt.Errorf("failed to open POST data file: %w", err)
	}
	defer file.Close()

	var data string
	scanner := bufio.NewReader(file)
	for {
		line, isPrefix, err := scanner.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			return "", fmt.Errorf("error reading POST data file: %w", err)
		}
		if !isPrefix {
			data += string(line) + "\r\n"
		}
	}

	return data, nil
}

// BuildGETRequestHeaders builds headers for a GET request
func (m *Manager) BuildGETRequestHeaders(useCustomHeaders bool, customHeaderFile string) (string, error) {
	if useCustomHeaders {
		headers, err := m.LoadCustomHeaders(customHeaderFile)
		if err != nil {
			return "", err
		}
		return headers, nil
	}
	return m.GenerateDefaultHeaders(), nil
}

// BuildPOSTRequestHeaders builds headers for a POST request
func (m *Manager) BuildPOSTRequestHeaders(useCustomHeaders bool, customHeaderFile string) (string, string, error) {
	var headers string
	var data string
	var err error

	if useCustomHeaders {
		headers, err = m.LoadCustomHeaders(customHeaderFile)
		if err != nil {
			return "", "", err
		}
		data, err = m.LoadPostData(customHeaderFile)
		if err != nil {
			return "", "", err
		}
	} else {
		data = "f"
		headers = "Connection: Keep-Alive\r\n"
		headers += "Content-Type: x-www-form-urlencoded\r\n"
		headers += fmt.Sprintf("Content-Length: %d\r\n", len(data))
		headers += "Accept-Encoding: gzip, deflate\r\n"
	}

	return headers, data, nil
}