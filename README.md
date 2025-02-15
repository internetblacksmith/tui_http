# 🌊 Restless

**The cheeky HTTP client that never sits still!**

Restless is a terminal-based HTTP client built with Go and Bubble Tea, designed to be a fun and lightweight alternative to Postman for developers who love working in the terminal.

## Features

- 🎨 Beautiful TUI interface with syntax highlighting
- 🚀 Support for all HTTP methods (GET, POST, PUT, DELETE, PATCH, HEAD)
- 📝 Request history tracking
- 🎯 Custom headers and request body support
- ⚡ Fast response times and clean JSON formatting
- 🌈 Colorful status code indicators

## Installation

```bash
# Clone the repository
git clone <your-repo-url>
cd tui_http

# Build the application
go build -o restless cmd/restless/main.go

# Run it
./restless
```

## Usage

1. **Switch between tabs** with `Tab`
2. **Navigate inputs** with `↑` and `↓` arrows
3. **Change HTTP method** with `←` and `→` arrows
4. **Send request** with `Enter`
5. **Quit** with `q` or `Ctrl+C`

### Request Tab
- Select HTTP method
- Enter URL
- Add request body (JSON, XML, etc.)
- Add custom headers

### Response Tab
- View status code and response time
- Browse response headers
- Read response body (truncated for readability)

### History Tab
- Browse previous requests
- Quick access to recently used endpoints

## Why "Restless"?

Because it's always making REST calls and never sits still! 😄

## Contributing

Feel free to contribute to make Restless even more cheeky and useful!

## License

MIT License - Go forth and REST!