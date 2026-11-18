````markdown
# go-top

**`go-top`** is a lightweight, modular, and cross-platform system monitor written in **Go**, built around a decoupled **Pub/Sub (Hub & Spoke)** architecture that cleanly separates system metrics collection from presentation layers.

---

## Architecture

At the core of the application is an in-memory **Central Hub** that periodically samples host resources (CPU, RAM, disk, network, and processes) and broadcasts updates in real-time via Go channels (`chan`) to all subscribed clients.

```text
                  +---------------------+
                  |  System Telemetry   |
                  |     (gopsutil)      |
                  +----------+----------+
                             |
                             v
                  +---------------------+
                  |     Central Hub     |
                  |  (Pub/Sub Broker)   |
                  +--+-------+-------+--+
                     |       |       |
          +----------+       |       +----------+
          v                  v                  v
+-------------------+ +-------------+ +-------------------+
|    Text Client    | |  REST API   | |  Web UI (HTMX)    |
|   (ANSI stdout)   | |   (JSON)    | | (net/http + SSE)  |
+-------------------+ +-------------+ +-------------------+
```
````

### Key Features

- **Zero Coupling:** The collection core has zero dependency on UI frameworks. Each client acts as an independent subscriber.
- **Single Binary Deployment:** Compiles into a static, self-contained executable ready for Linux and macOS environments with zero external runtime dependencies.
- **Multiple Operating Modes:**
- **Text Client (`text`):** Minimalist terminal output with clean in-place updates using ANSI escape sequences.
- **TUI (`tui`):** Interactive terminal user interface for rich system inspection.
- **Web GUI & REST API (`web`):** Native HTTP server (`net/http` Go 1.22+) serving JSON endpoints and an HTMX-powered reactive dashboard.

---

## Project Structure

```text
go-top/
├── cmd/
│   └── go-top/
│       └── main.go            # Entry point, CLI flag parsing, and orchestration
├── internal/
│   ├── collector/
│   │   ├── hub.go             # Pub/Sub core: sampling and broadcasting
│   │   └── sysinfo.go         # Specific metric collection modules
│   ├── client/
│   │   └── text.go            # Pure text ANSI client
│   ├── api/
│   │   ├── handlers.go        # REST JSON endpoints and HTML partials
│   │   └── server.go          # net/http server configuration
│   └── web/
│       ├── static/            # Static assets (CSS/Tailwind)
│       └── templates/         # Go HTML templates for HTMX
├── go.mod
└── go.sum

```

---

## Requirements and Installation

- **Go** 1.22 or higher

### Building from Source

```bash
git clone [https://github.com/your-username/go-top.git](https://github.com/your-username/go-top.git)
cd go-top
go mod tidy
go build -o go-top ./cmd/go-top
```

---

## Usage

Run the binary by specifying the desired mode and sampling interval:

```bash
# Run in text mode (default, 1s interval)
./go-top

# Run with a custom sampling interval
./go-top --mode text --interval 500ms

# Run the Web Server / REST API
./go-top --mode web --interval 1s
```

### CLI Options

| Flag         | Default | Description                           |
| ------------ | ------- | ------------------------------------- |
| `--mode`     | `text`  | Operating mode (`text`, `tui`, `web`) |
| `--interval` | `1s`    | Metrics sampling frequency            |

---

## Roadmap

- [x] Pub/Sub Hub Core (`internal/collector`)
- [x] CLI argument parsing and ANSI text client (`text`)
- [ ] Network interfaces and process list telemetry extension
- [ ] Native HTTP server with REST endpoints (`GET /api/stats`)
- [ ] Web dashboard with HTMX and real-time updates
- [ ] Advanced TUI interface
