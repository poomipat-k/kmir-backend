# KMIR Backend

The KMIR Backend is a Go-based backend service designed to support the KMIR (Knowledge Management and International Relations) system. It provides RESTful APIs for managing and retrieving knowledge data efficiently.

## 🚀 Features

- RESTful API: Provides endpoints for CRUD operations on knowledge data.

- Dockerized: Easily deployable using Docker and Docker Compose.

- Modular Architecture: Clean separation of concerns for scalability and maintainability.

## 🛠️ Technologies Used

- Backend: Go

- Database: PostgreSQL

- Containerization: Docker

- Web Server: Nginx

- Development Tools:
  - Docker Compose for local development
  - Makefile for task automation

## ⚙️ Setup Instructions

## Prerequisites

Ensure you have the following installed:

[Go](https://go.dev/doc/install)

[Docker](https://docs.docker.com/engine/install/)

[Docker Compose](https://docs.docker.com/compose/install/)

[Make](https://makefiletutorial.com/)

## Local Development

Clone the repository:

```
git clone https://github.com/poomipat-k/kmir-backend.git
cd kmir-backend
```

Copy the example environment variables:

```
cp .env.example .env
```

Build and start the application using Docker Compose:

```
make up_build_dev
```

Access the application at `http://localhost:8080`
