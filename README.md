# Golang-httpflood ![](https://img.shields.io/badge/Version-2.0-brightgreen.svg) ![](https://img.shields.io/badge/license-MIT-blue.svg) ![](https://img.shields.io/badge/Go-1.21+-00ADD8?logo=go)

Using Golang(net/socket) to HTTP flood - A modular, high-performance HTTP flood testing tool.

**Warning: Please use command "ulimit -n 999999" before use this on Linux**

## Features

### Core Features
- [x] HTTP Get Flood
- [x] HTTP Post Flood
- [x] Random URL parameters (GET requests)
- [x] Custom header file support
- [x] Improved threading control
- [x] High-performance flood engine
- [x] Auto IP resolution from domain
- [x] Fixed for 386 systems
- [x] **Real-time statistics display**
- [x] **Modular architecture**

### Default Header Configuration
- [x] Random user-agents (Chrome, IE, Spider bots)
- [x] Random Accept headers (14 variations)
- [x] Random Referer URLs (15 domains)
- [x] Random POST data

### New in Version 2.0
- 📊 **Real-time statistics** with metrics dashboard
- 🏗️ **Modular architecture** with clean separation of concerns
- 🧩 **Extensible design** for easy feature additions
- 📈 **Performance tracking** (requests/sec, bandwidth, response times)
- 🔧 **Better error handling** and validation
- 📚 **Comprehensive documentation**

## Architecture

This project follows Go best practices with a clean, modular architecture:

```
cmd/httpflood/          # Application entry point
internal/
├── config/            # Configuration management
├── flood/             # Core flooding logic
├── useragents/        # User agent generation
├── headers/           # HTTP header management
├── stats/             # Statistics and metrics
└── utils/             # Utility functions
pkg/                   # Public packages (planned)
configs/               # Configuration files and templates
examples/              # Usage examples
docs/                  # Documentation
```

For detailed architecture information, see [docs/ARCHITECTURE.md](docs/ARCHITECTURE.md)

## Installation

### Prerequisites
- Go 1.21 or higher

### Build from Source

```bash
# Clone the repository
git clone https://github.com/bonbonbon1414/golang-httpflood.git
cd golang-httpflood

# Build the application
go build -o httpflood cmd/httpflood/main.go

# Or use go install
go install github.com/bonbonbon1414/golang-httpflood/cmd/httpflood@latest
```

### Linux Users

Before running on Linux, increase the file descriptor limit:

```bash
ulimit -n 999999
```

## Usage

### Basic Usage

```bash
./httpflood <url> <threads> <get/post> <seconds> <header.txt/nil>
```

### Arguments

| Argument | Description | Example |
|----------|-------------|---------|
| `url` | Target URL | `https://example.com` |
| `threads` | Number of concurrent goroutines | `100` |
| `mode` | HTTP mode: `get` or `post` | `get` |
| `seconds` | Duration of flood in seconds | `60` |
| `header.txt` | Custom header file or `nil` for defaults | `nil` |

### Examples

#### GET Flood with Default Headers
```bash
./httpflood https://example.com 100 get 60 nil
```

#### GET Flood with High Thread Count
```bash
./httpflood https://example.com 500 get 120 nil
```

#### POST Flood with Default Headers
```bash
./httpflood https://example.com 100 post 60 nil
```

#### GET Flood with Custom Headers
```bash
./httpflood https://example.com 100 get 60 headers.txt
```

#### POST Flood with Custom Headers
```bash
./httpflood https://example.com 200 post 60 headers.txt
```

### Custom Header Format

Create a text file with custom headers:

```
Accept: text/html
User-Agent: Mozilla/5.0 (compatible; MyBot/1.0)
Referer: https://google.com
Authorization: Bearer token123
```

If you don't have specific requirements, use `nil` for randomized headers.

## Performance

### Thread Count Guidelines
- **1 Thread** = 1 connection
- **100-300 threads** can overwhelm normal websites in ~10 seconds
- **1000-5000 threads** = easily achievable with Go's goroutines

### Why High Thread Counts?
Go's goroutines are lightweight (2KB stack) compared to OS threads (2MB stack). This allows running thousands of concurrent operations with minimal overhead.

[Learn more about goroutines vs threads](http://tleyden.github.io/blog/2014/10/30/goroutines-vs-threads/)

## Statistics Display

The tool now provides real-time statistics during operation:

```
========== STATISTICS ==========
Total Requests:      50000
Successful:          49800 (99.60%)
Failed:              200 (0.40%)
Requests/sec:        833.33
Active Connections:  100
Bytes Sent:          52428800 (50.00 MB)
Bandwidth:           41.67 MB/s
Avg Response Time:   0.12 ms
Elapsed Time:        60.00 seconds
===============================
```

## Configuration

Configuration can be provided via:
1. Command-line arguments (primary method)
2. YAML configuration files (planned)

See `configs/` directory for example configurations.

## Project Structure

```
golang-httpflood/
├── cmd/httpflood/              # Main application
├── internal/
│   ├── config/                 # Configuration handling
│   ├── flood/                  # Flood engine
│   ├── useragents/             # User agent generation
│   ├── headers/                # Header management
│   ├── stats/                  # Statistics tracking
│   └── utils/                  # Utilities
├── pkg/                        # Public packages
├── configs/                    # Configuration files
├── examples/                   # Usage examples
├── docs/                       # Documentation
│   ├── ARCHITECTURE.md        # Architecture docs
│   ├── API.md                 # API documentation
│   └── TROUBLESHOOTING.md     # Troubleshooting guide
├── go.mod                      # Go module definition
├── go.sum                      # Go module checksums
├── Makefile                    # Build automation
├── Dockerfile                  # Docker image
└── README.md                   # This file
```

## Development

### Building

```bash
# Build for current platform
go build -o httpflood cmd/httpflood/main.go

# Build for multiple platforms
make build-all
```

### Testing

```bash
# Run tests
go test ./...

# Run with coverage
go test -cover ./...
```

### Linting

```bash
# Run linter
make lint
```

## Contributing

Contributions are welcome! Please read [docs/CONTRIBUTING.md](docs/CONTRIBUTING.md) for guidelines.

## Troubleshooting

See [docs/TROUBLESHOOTING.md](docs/TROUBLESHOOTING.md) for common issues and solutions.

## License

MIT License - see [LICENSE](LICENSE) file for details.

## Credits

- Original code by LeeOn123
- Version 2.0 refactoring and improvements

## Disclaimer

This tool is designed for **educational and testing purposes only**. Always ensure you have proper authorization before testing any system. Unauthorized use is illegal and unethical.

## Roadmap

### Planned Features
- [ ] Configuration file support (YAML/TOML)
- [ ] Proxy support
- [ ] HTTP/2 support
- [ ] Cookie management
- [ ] Rate limiting
- [ ] Multiple target support
- [ ] IPv6 support
- [ ] Interactive CLI mode
- [ ] API for programmatic control

See [docs/ROADMAP.md](docs/ROADMAP.md) for detailed plans.

## Links

- [GitHub Repository](https://github.com/bonbonbon1414/golang-httpflood)
- [Issues](https://github.com/bonbonbon1414/golang-httpflood/issues)
- [Documentation](docs/)

---

**Version 2.0** | **Coded by LeeOn123** | **Refactored for improved maintainability**