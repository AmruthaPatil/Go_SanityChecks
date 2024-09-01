# Go Database Sanity Checks

A Go-based application to perform comprehensive sanity checks for multiple database and storage services, including ArangoDB, PostgreSQL, and Redis. This application ensures that each service is running correctly, accessible, and performing optimally, while also verifying data integrity, availability, and performance.

## Table of Contents

- [Introduction](#introduction)
- [Features](#features)
- [Technologies Used](#technologies-used)
- [Setup and Installation](#setup-and-installation)
- [Running the Application](#running-the-application)

## Introduction

This project aims to provide a robust mechanism to ensure the health and proper functioning of various database and storage services deployed within a Docker environment. The application is written in Go and uses standard libraries along with database-specific clients to interact with ArangoDB, PostgreSQL, and Redis.

It verifies that each service is:

- Available and running correctly
- Properly configured
- Free of data integrity issues
- Performing within acceptable limits

## Features

- **Service Availability Check**: Confirms that the required services are running and accessible.
- **Connection Validation**: Ensures that the application can connect to each database service using the correct credentials.
- **Data Integrity Check**: Verifies the presence and correctness of specific data within the databases.
- **Performance Monitoring**: Monitors query response times to ensure optimal performance.
- **Retry Logic**: Implements retry mechanisms to handle intermittent connectivity issues.

## Technologies Used

- **Go (Golang)**: Main programming language used for building the application.
- **Docker & Docker Compose**: For containerizing and orchestrating the database services.
- **ArangoDB Go Driver**: Library for connecting to and interacting with ArangoDB.
- **PostgreSQL Go Driver (`lib/pq`)**: Library for connecting to and interacting with PostgreSQL.
- **Redis Go Client (`go-redis/redis/v8`)**: Library for connecting to and interacting with Redis.

## Setup and Installation

### Prerequisites

- Go 1.20 or later
- Docker and Docker Compose

### Installation Steps

1. **Clone the Repository:**

    ```bash
    git clone https://github.com/yourusername/go-database-sanity-checks.git
    cd go-database-sanity-checks
    ```

2. **Build Docker Containers:**

    Ensure Docker is running, and then build the containers:

    ```bash
    docker-compose up --build
    ```

3. **Run the Go Application:**

    The Go application will automatically run as part of the Docker Compose setup. You can view the logs to monitor progress:

    ```bash
    docker-compose logs -f sanity-check
    ```

## Running the Application

The application runs as part of a Docker Compose setup. It performs a series of sanity checks on the following services:

- **ArangoDB**: A NoSQL multi-model database.
- **PostgreSQL**: A powerful, open-source relational database.
- **Redis**: An in-memory key-value store used for caching.


Thank you for exploring this repository!