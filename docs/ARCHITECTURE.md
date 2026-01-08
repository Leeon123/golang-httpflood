# Architecture Documentation

## Overview

golang-httpflood is a modular HTTP flood testing tool written in Go. The project follows a clean architecture pattern with clear separation of concerns.

## Project Structure

```
golang-httpflood/
├── cmd/httpflood/          # Application entry point
├── internal/               # Private packages
│   ├── config/            # Configuration management
│   ├── flood/             # Core flooding logic
│   ├── useragents/        # User agent generation
│   ├── headers/           # HTTP header management
│   ├── stats/             # Statistics and metrics
│   └── utils/             # Utility functions
├── pkg/                   # Public packages
│   ├── proxy/            # Proxy support (planned)
│   └── cookies/          # Cookie management (planned)
└── configs/               # Configuration files
```

## Core Components

### 1. Configuration (`internal/config`)

**Purpose**: Manages application configuration from command-line arguments

**Key Components**:
- `Config struct`: Holds all configuration parameters
- `ParseFromArgs()`: Parses CLI arguments into Config
- `Validate()`: Validates configuration parameters

**Flow**:
1. Parse URL and extract host, port, path
2. Parse threads, duration, mode
3. Parse custom header file path
4. Validate all parameters

### 2. User Agent Generation (`internal/useragents`)

**Purpose**: Generates realistic random user agents

**Key Components**:
- `Generator`: Main user agent generator
- Component data: Platforms, OS versions, browser versions

**Flow**:
1. Select random platform (Macintosh, Windows, X11)
2. Select random OS version for platform
3. Select random browser type
4. Generate appropriate user agent string

### 3. Header Management (`internal/headers`)

**Purpose**: Manages HTTP header generation and loading

**Key Components**:
- `Manager`: Header management
- `acceptHeaders[]`: Random Accept header variations
- `referers[]`: Random Referer URLs

**Flow**:
1. Check if custom headers should be used
2. Load custom headers from file if specified
3. Otherwise generate randomized default headers
4. Include User-Agent, Accept, Referer

### 4. Flood Engine (`internal/flood`)

**Purpose**: Core HTTP flooding logic

**Key Components**:
- `Flood`: Main flood coordinator
- `RequestBuilder`: Constructs HTTP requests
- Worker goroutines: Concurrent request senders

**Flow**:
1. Spawn N worker goroutines
2. Workers wait for start signal
3. Each worker:
   - Establishes TCP/TLS connection
   - Sends burst of requests (100 per connection)
   - Updates metrics
   - Closes and reconnects
4. Continues until stop signal or timeout

### 5. Statistics (`internal/stats`)

**Purpose**: Tracks and displays operation statistics

**Key Components**:
- `Metrics`: Thread-safe metrics tracking
- `MetricsSnapshot`: Read-only metrics snapshot
- Display functions: Console output

**Metrics Tracked**:
- Total requests
- Successful/failed requests
- Bytes sent
- Active connections
- Requests per second
- Response times
- Bandwidth

### 6. Utilities (`internal/utils`)

**Purpose**: Shared utility functions

**Key Components**:
- Random number generation
- String operations
- Network utilities

## Data Flow

```
User Input (CLI)
    ↓
Config Parser
    ↓
Component Initialization
    ├─→ User Agent Generator
    ├─→ Header Manager
    └─→ Metrics Collector
    ↓
Flood Operation
    ├─→ Worker Goroutines
    ├─→ Request Builder
    └─→ Connection Manager
    ↓
Statistics Display
    ↓
Summary Report
```

## Concurrency Model

The application uses Go's goroutines for concurrent operations:

1. **Main Goroutine**: Coordinates the flood operation
2. **Worker Goroutines**: N concurrent request senders (configurable)
3. **Stats Display Goroutine**: Periodic statistics updates

**Synchronization**:
- `startChan` channel coordinates worker start
- `stopChan` channel coordinates graceful shutdown
- `sync.WaitGroup` ensures all workers complete
- Atomic operations for thread-safe metrics

## Request Lifecycle

1. **Connection Establishment**:
   - TCP connection for HTTP
   - TLS handshake for HTTPS
   - Connection metrics updated

2. **Request Building**:
   - GET: Path + random parameters
   - POST: Path + headers + data
   - Headers randomized

3. **Request Sending**:
   - Write to connection
   - Metrics updated
   - Response time recorded

4. **Connection Reuse**:
   - 100 requests per connection
   - Connection closed
   - New connection established

## Error Handling

- Connection errors: Logged, connection retried
- Header file errors: Fatal exit with message
- Configuration errors: Fatal exit with message
- Write errors: Failure count incremented

## Performance Considerations

1. **Goroutine Efficiency**: Go's lightweight goroutines allow thousands of concurrent workers
2. **Connection Pooling**: 100 requests per connection reduces overhead
3. **Atomic Operations**: Lock-free metrics for minimal contention
4. **Microsecond Delays**: Staggered worker startup to prevent thundering herd

## Extensibility

The modular design allows easy addition of:
- Proxy support (pkg/proxy)
- Cookie management (pkg/cookies)
- HTTP/2 support
- Rate limiting
- Configuration file support
- Multiple target support