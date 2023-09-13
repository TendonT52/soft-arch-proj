# Software Architecture Project - User Service

## Prerequisites

Before you begin, ensure you have the following software/tools installed:

- [Docker Desktop](https://www.docker.com/products/docker-desktop)
- [Homebrew](https://brew.sh/)

## Installation

### Install Kubernetes on Docker Desktop

1. Make sure you have Docker Desktop installed.
2. Enable Kubernetes in Docker Desktop settings.
3. Start Kubernetes using Docker Desktop.

### Install kubectl using Homebrew

```bash
brew install kubectl
```

### Install golang-migrate

```bash
go get -u -d github.com/golang-migrate/migrate/cmd/migrate
cd $GOPATH/src/github.com/golang-migrate/migrate/cmd/migrate
```

### Install Telepresence

#### Intel

```bash
brew install datawire/blackbird/telepresence
```

#### Apple Silicon

```bash
brew install datawire/blackbird/telepresence-arm64
```

### Setup Kubernetes

Run the following script to set up Kubernetes:

```bash
make k8s-run-dev
```

### Setup Database

Run the following script to set up the database:

```bash
make migrate-up
```

## Connecting to gRPC

To connect to gRPC using your application, use the following host and port:

- Host: localhost
- Port: 8080

## Generating gRPC Code
Run the following script to generate gRPC code:

```bash
make gen-proto
```

## Running the Application
Run the following script to run the application:

```bash
go run .
```

## Testing the Application
Run the following script to test the application:

```bash
make test-all
```

## Database Details
Postgresql
- Host: user-service-postgresql.default
- Port: 5432
- Username: postgres
- Password: password

Redis
- Host: user-service-redis-master.default
- Port: 6379
- Password: password

Memphis
- Host: memphis.default
- Port: 9000
- username: root
- password: password
