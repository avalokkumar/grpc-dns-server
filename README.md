# gRPC DNS Server

## Overview
A gRPC-based DNS server with support for DNS-over-HTTPS, caching, load balancing, and Prometheus metrics.

## Features
- DNS resolution using upstream servers.
- DNS-over-HTTPS (DoH) support.
- In-memory caching for faster responses.
- Load balancing across multiple DNS servers.
- Prometheus metrics for monitoring.

## Usage
### Run the Server
```bash
make run
```

### Run Tests
```bash
make test
```

### Build the Server
```bash
make build
```

## Configuration
Modify `config/config.yaml` to update DNS servers, caching policies, and logging levels.

```yaml
dns_servers:
  - "8.8.8.8"
  - "8.8.4.4"
cache:
  enabled: true
  ttl: 300
logging:
  level: "info"
grpc:
  port: 50051
```


---

## To build and run

## **🚀 Prerequisites**
Before proceeding, ensure you have the following installed:

✅ **Go 1.20+** → [Download Go](https://go.dev/dl/)  
✅ **Docker & Docker Compose** (for containerization)  
✅ **Protobuf Compiler (`protoc`)** → [Install Guide](https://grpc.io/docs/protoc-installation/)  
✅ **`protoc-gen-go` & `protoc-gen-go-grpc`** → Installed via Go

```sh
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

---

# **📌 Step 1: Clone the Repository**
```sh
git clone https://github.com/avalokkumar/grpc-dns-server.git
cd grpc-dns-server
```

---

# **📌 Step 2: Generate gRPC Code**
Compile the `dns.proto` file to generate Go code.

```sh
protoc --go_out=pkg/proto --go-grpc_out=pkg/proto \
  --proto_path=pkg/proto dns.proto
```

This will generate:
```
pkg/proto/
├── dns.pb.go        # gRPC message structures
├── dns_grpc.pb.go   # gRPC server interface
```

---

# **📌 Step 3: Build the Application**
Build the gRPC server.

```sh
go build -o bin/dns-server cmd/server/main.go
```

(Optional) Build the client.

```sh
go build -o bin/dns-client cmd/client/main.go
```

---

# **📌 Step 4: Run the gRPC DNS Server**
Start the gRPC server:

```sh
./bin/dns-server
```

By default, the server runs on `localhost:50051`.

---

# **📌 Step 5: Run the gRPC Client**
Use the client to send a DNS query.

```sh
./bin/dns-client --domain example.com
```

---

# **📌 Step 6: Run via Docker (Optional)**
You can run the server in a **Docker container**.

### **Build the Docker image**
```sh
docker build -t grpc-dns-server .
```

### **Run the server**
```sh
docker run -p 50051:50051 grpc-dns-server
```

---

# **📌 Step 7: Run Tests**
Run **unit tests**:

```sh
go test ./test/unit/...
```

Run **integration tests**:

```sh
go test ./test/integration/...
```