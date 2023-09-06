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
# go build -tags 'postgres' -o $GOPATH/bin/migrate github.com/golang-migrate/migrate/cmd/migrate
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
./k8s-setup.sh
```

### Setup Database

Run the following script to set up the database:

```bash
./migration.sh
```

## Database UI

Access the database UI using the following URL:

- URL: [http://todo-pgadmin.default/](http://todo-pgadmin.default/)
- Username: admin@admin.com
- Password: mypassword

## Connecting to gRPC

To connect to gRPC using your application, use the following host and port:

- Host: localhost
- Port: 8080

## Generating gRPC Code

Use the `compile-protos.sh` script to generate gRPC code:

```bash
./compile-protos.sh
```

## Database Details

- Database Host: user-postgres.default
- Port: 5432
- Username: root
- Password: mypassword