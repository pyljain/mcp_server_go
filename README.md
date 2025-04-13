# MCP Server (Go)

A Go implementation of a Model Context Protocol (MCP) server that provides database querying capabilities through a JSON-RPC interface.

## Features

- JSON-RPC 2.0 protocol implementation
- Server-Sent Events (SSE) for real-time communication
- SQLite database integration
- Tool-based architecture for extensible functionality
- Authentication support

## Available Tools

1. **Query Tool**
   - Execute SQL queries against the database
   - Returns results in tabular format

2. **List Tables Tool**
   - List all available tables in the database

## Getting Started

### Prerequisites

- Go 1.x
- SQLite3

### Installation

1. Clone the repository:
```bash
git clone <repository-url>
cd mcp_server_go
```

2. Install dependencies:
```bash
go mod download
```

### Configuration

The server uses a simple token-based authentication system. The expected token is hardcoded as "abcd" in the server configuration.

### Running the Server

```bash
go run main.go
```

The server will start on port 8777.

## API Documentation

### Authentication

All requests must include an Authorization header with the format:
```
Authorization: Bearer <token>
```

### Endpoints

#### 1. Root Endpoint (`/`)
- Method: GET
- Purpose: Establishes a Server-Sent Events (SSE) connection
- Response: Returns an endpoint URL for message communication

#### 2. Messages Endpoint (`/messages/{sessionID}`)
- Method: POST
- Purpose: Handle client requests
- Supported Methods:
  - `initialize`: Initialize the connection
  - `tools/list`: List available tools
  - `tools/call`: Execute a specific tool

### Example Usage

1. Initialize connection:
```json
{
  "jsonrpc": "2.0",
  "id": 1,
  "method": "initialize",
  "params": {}
}
```

2. List available tools:
```json
{
  "jsonrpc": "2.0",
  "id": 2,
  "method": "tools/list",
  "params": {}
}
```

3. Execute a query:
```json
{
  "jsonrpc": "2.0",
  "id": 3,
  "method": "tools/call",
  "params": {
    "name": "query",
    "arguments": {
      "query": "SELECT * FROM your_table"
    }
  }
}
```

## Project Structure

```
mcp_server_go/
├── main.go              # Main server implementation
├── pkg/
│   ├── messages/        # Message structures
│   ├── methods/         # Request handling methods
│   └── tools/           # Tool implementations
└── mcp.db              # SQLite database file
```
