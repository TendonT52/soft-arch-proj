# Setup Guide

This document provides step-by-step instructions on how to install necessary packages and tools for your system.

## Before You Begin

- Ensure [Homebrew](https://brew.sh/) is installed on your system. If not, set it up with the following command:
    ```bash
    /bin/bash -c "$(curl -fsSL <https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh>)"
    ```

## Installation Procedures

1. **Set Up Parallel**
   
    ```bash
    brew install parallel
    ```

3. **Set Up Kubernetes CLI**
   
    ```bash
    brew install kubernetes-cli
    ```

5. **Set Up Helm**
   
    ```bash
    brew install helm
    ```

7. **Set Up Golang Migrate**
   
    ```bash
    brew install golang-migrate
    ```

9. **Set Up psql**

    > Note: If you already have postgresql on your system, you can bypass this step.

    Execute the following to force brew to link the binary. By default, brew avoids linking binaries if a link already exists. In this context, libpq (a CLI-only client encompassing psql, pg_dump) is linked similarly to the comprehensive postgresql package, which includes both the server and client. This might lead to complications if postgresql is pre-installed.
   
    ```bash
    brew link --force libpq
    ```

11. **Set Up Telepresence**

Please avoid from using brew for this installation since it provides the non-open source version, which may prompt for trial expiration. Adhere to the steps below for Telepresence installation:

- Retrieve and position telepresence within `/usr/local/bin/`:
  
     ```bash
     sudo curl -fL https://app.getambassador.io/download/tel2oss/releases/download/v2.15.1/telepresence-darwin-arm64 -o /usr/local/bin/telepresence
     ```
- Grant the required permissions:
  
     ```bash
     sudo chmod a+x /usr/local/bin/telepresence
     ```

## Set Up Golang and Associated Tools

1. **Install Golang**:
    First, you'll need to ensure you have Golang installed. If not, install it via Homebrew:

    ```bash
    brew install go
    ```

    After installation, verify that Go is installed correctly:

    ```bash
    go version
    ```

2. **Install Protobuf**:

    ```bash
    brew install protobuf
    ```

3. **Install Protobuf Go tools and GRPC-Gateway**:

    ```bash
    go install \
        google.golang.org/protobuf/cmd/protoc-gen-go@latest \
        github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
        github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
        google.golang.org/protobuf/cmd/protoc-gen-go \
        google.golang.org/grpc/cmd/protoc-gen-go-grpc
    ```

4. **Install Statik**:

    ```bash
    go get github.com/rakyll/statik
    ```

Remember to follow each of the steps sequentially for a seamless setup.

## Confirming Your Installation

Post installation, validate the setups with these commands:

- `parallel --version`
- `kubectl version`
- `helm version`
- `golang-migrate version`
- `telepresence version`
