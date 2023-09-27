# Configuration Instructions

This README provides detailed instructions on how to install the required packages and tools for your system.

## Preparatory Steps

- Make sure [Homebrew](https://brew.sh/) is installed on your system. If not, install it using the following command:

```bash
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh>)"
```

## Steps for Installation

1. **Install Parallel**

```bash
brew install parallel
```

2. **Install Kubernetes CLI**

```bash
brew install kubernetes-cli
```

3. **Install Helm**

```bash
brew install helm
```

4. **Install Golang Migrate**

```bash
brew install golang-migrate
```

5. **Install psql**

>If you have postgresql pre-installed on your system, you can skip this step.

Run the following command to link the binary forcefully since brew avoids binary links if a link already exists. Here, libpq (a CLI-only client inclusive of psql, pg_dump) is linked the same way as the complete postgresql package. This might create issues if postgresql is already installed.

```bash
brew link --force libpq
```

6. **Install Telepresence**

It's recommended not to use brew for this installation as it provides a non-open source version which might ask for trial expiration. Follow the below steps for Telepresence installation:

- Download and place telepresence inside `/usr/local/bin/`:

```bash
sudo curl -fL https://app.getambassador.io/download/tel2oss/releases/download/v2.15.1/telepresence-darwin-arm64 -o /usr/local/bin/telepresence
```

- Provide necessary permissions:

```bash
sudo chmod a+x /usr/local/bin/telepresence
```

## Installation of Golang and its Tools

1. **Golang Installation**:
Make sure Golang is installed on your system. If not, install it with Homebrew:

```bash
brew install go
```

Then, confirm the successful installation of Go:

```bash
go version
```

2. **Protobuf Installation**:

```bash
brew install protobuf
```

3. **Installation of Protobuf Go Tools and GRPC-Gateway**:

```bash
go install \
    google.golang.org/protobuf/cmd/protoc-gen-go@latest \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
    google.golang.org/protobuf/cmd/protoc-gen-go \
    google.golang.org/grpc/cmd/protoc-gen-go-grpc \
    github.com/rakyll/statik
```

Ensure you follow these steps sequentially for a smooth setup process.

## Verification of Installation

After installation, verify the setup using these commands:

- `parallel --version`
- `kubectl version`
- `helm version`
- `golang-migrate version`
- `telepresence version`
