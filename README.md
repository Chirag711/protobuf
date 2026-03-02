# Protocol Buffers with Go (Golang)

This repository demonstrates how to use **Protocol Buffers (Protobuf)** with **Golang**.
It includes a simple example of defining a `.proto` file, generating Go code using `protoc`, and running a Go program that uses the generated protobuf structures.

Protocol Buffers is a **language-neutral, platform-neutral, extensible mechanism for serializing structured data**, developed by Google.

---

# Project Structure

```
protobuf
│
├── go.mod
├── main.go
├── user.proto
└── userpb
    └── user.pb.go
```

| File         | Description                       |
| ------------ | --------------------------------- |
| `user.proto` | Protobuf message definition       |
| `user.pb.go` | Generated Go code from proto file |
| `main.go`    | Example Go program using protobuf |
| `go.mod`     | Go module configuration           |

---

# Prerequisites

Make sure the following tools are installed:

* **Go (1.20 or later)**
* **Protocol Buffers Compiler (`protoc`)**

Check installation:

```
go version
protoc --version
```

---

# Install Go Protobuf Plugin

Install the required plugin for Go:

```
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
```

Add Go binary path to environment if needed.

Example (Windows):

```
set PATH=%PATH%;%USERPROFILE%\go\bin
```

---

# Generate Go Code from Proto

Run the following command:

```
protoc --go_out=. user.proto
```

This generates:

```
user.pb.go
```

inside the `userpb` folder.

---

# Initialize Go Module

If not already initialized:

```
go mod init protobuf
go mod tidy
```

This installs required dependencies.

---

# Run the Application

Run the Go program:

```
go run main.go
```

Example Output:

```
User ID: 1
User Name: Chirag
```

---

# Example Proto File

```
syntax = "proto3";

package user;

option go_package = "protobuf/userpb";

message User {
  int32 id = 1;
  string name = 2;
}
```

---

# Example Go Code

```
package main

import (
	"fmt"
	"protobuf/userpb"
)

func main() {

	user := userpb.User{
		Id:   1,
		Name: "Chirag",
	}

	fmt.Println("User ID:", user.Id)
	fmt.Println("User Name:", user.Name)
}
```

---

# Useful Commands

Generate protobuf file

```
protoc --go_out=. user.proto
```

Install dependencies

```
go mod tidy
```

Run program

```
go run main.go
```

---

# Future Improvements

* Add **gRPC service implementation**
* Create **server and client example**
* Add **JSON ↔ Protobuf conversion**
* Add **Docker support**

---

# Author

**Chirag Raul**

GitHub:
https://github.com/Chirag711

---

# License

This project is for **learning and educational purposes**.
