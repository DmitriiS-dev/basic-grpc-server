# TaskFlow

A distributed task management system built with Go and gRPC.

## Components

### Backend Server
Provides a gRPC API for task management.

Features:
- Create tasks
- Get tasks
- List tasks
- Complete tasks

### CLI Client
A Cobra-powered command line client that communicates with the backend over gRPC.

Features:
- Add tasks
- Get tasks
- List tasks
- Complete tasks

---

# Prerequisites

Install:

- Go 1.24+
- Protocol Buffers Compiler (protoc)

Verify installation:

```bash
go version
protoc --version
```

---

# Installing Protocol Buffer Tooling

Install Go protobuf generators:

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

Verify:

```bash
protoc-gen-go --version
protoc-gen-go-grpc --version
```

Ensure your Go bin directory is on PATH.

Windows default:

```text
C:\Users\<username>\go\bin
```

---

# Installing Cobra CLI

Install Cobra scaffolding tool:

```bash
go install github.com/spf13/cobra-cli@latest
```

Verify:

```bash
cobra-cli --help
```

---

# Backend Setup

Navigate to backend project:

```bash
cd basic-backend
```

Install dependencies:

```bash
go mod tidy
```

Generate protobuf code:

```bash
protoc --go_out=. --go-grpc_out=. proto/task.proto
```

Run server:

```bash
go run ./cmd/project
```

Server listens on:

```text
localhost:50051
```

---

# CLI Setup

Navigate to CLI project:

```bash
cd example-client
```

Install dependencies:

```bash
go mod tidy
```

Run CLI:

```bash
go run . --help
```

Example:

```bash
go run . list
```

Connect to custom server:

```bash
go run . list --server localhost:50051
```

# Development Workflow

1. Update `task.proto`
2. Regenerate protobuf files

```bash
protoc --go_out=. --go-grpc_out=. proto/task.proto
```

3. Restart server
4. Test with CLI

---
