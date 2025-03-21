# TinyPath - A Simple URL Shortener

TinyPath is a simple URL shortener service that allows you to convert long URLs into shorter, more manageable links.

## Features

* **URL Shortening:** Converts long URLs into short, unique links.
* **Redirection:** Redirects users from the short URL to the original URL.
* **Custom Short URLs (Optional):** Allows users to specify a custom short URL during creation.
* **Retrieve Short Link Details:** Provides an API to fetch details of a short URL, including the original URL and access count.
* **Delete Short Links:** Enables the deletion of existing short URLs.
* **Update Short Links:** Allows updating the original URL associated with a short URL.
* **Health Check:** Provides an endpoint to check the service's health status.

## Getting Started

### Prerequisites

* **Go:** Make sure you have Go installed on your system (version 1.22 or later is recommended).
* **Dependencies:** The project uses standard Go libraries and some internal packages within the repository. You will need to fetch these if you are setting up the project from scratch.

### Installation

1. **Clone the repository:**
    ```bash
    git clone <repository_url>
    cd tinypath
    ```
2. **Navigate to the project root.**
3. **Build the application:**
    ```bash
    go build ./cmd/api
    ```

### Configuration
#### Manually Create the Database (Recommended)

Run this SQL command before applying migrations:
```bash
psql -U postgres -h localhost -c "CREATE DATABASE url_shortener;"
```
#### Environment Variables

The application requires a .env file to store configuration values. Create a .env file in the root directory with the following contents:
```bash
URLSHORTNER_DB_DSN=postgres://username:password@host:5432/url_shortener?sslmode=disable
```
#### Database Migrations

This project uses Goose for managing database migrations.
Setting Up Migrations
```bash
Install Goose (if not installed):
```
```bash
go install github.com/pressly/goose/v3/cmd/goose@latest
```

1. Apply Migrations:

```bash
goose -dir migrations postgres "$URLSHORTNER_DB_DSN" up
```

2. Rollback Migrations:
```bash
goose -dir migrations postgres "$URLSHORTNER_DB_DSN" down
```

#### Environment Variables for Migrations

Ensure the URLSHORTNER_DB_DSN variable is set before running migrations:
```bash
export URLSHORTNER_DB_DSN="postgres://username:password@host:5432/url_shortener?sslmode=disable"
```

### Running the Application

Once built and configured, you can run the application:

```bash
./api
```

The service will then be accessible on the configured port (default is often 8080).

## API Endpoints

### General

#### `GET /`
* **Description:** Returns a welcome message.
* **Response (JSON):**
  ```json
  {
    "message": "Welcome to Tiny URL!"
  }
  ```

#### `GET /api/v1/healthcheck`
* **Description:** Performs a health check on the service.
* **Response (JSON):**
  ```json
  {
    "status": "ok",
    "message": "service is healthy"
  }
  ```

### Short URL Management

#### `POST /api/v1/short`
* **Description:** Creates a new short URL.
* **Request Body (JSON):**
  ```json
  {
    "original_url": "https://example.com",
    "short_url": "custom123"  // Optional
  }
  ```
* **Response (JSON):**
  ```json
  {
    "id": "unique_id",
    "short_url": "short123",
    "original_url": "https://example.com",
    "created_at": "2025-01-01T12:00:00Z",
    "updated_at": "2025-01-01T12:00:00Z"
  }
  ```

#### `GET /api/v1/short/{short}`
* **Description:** Retrieves details of a short URL.
* **Path Parameter:**
  - `short`: The unique short URL identifier.
* **Response (JSON):**
  ```json
  {
    "id": "unique_id",
    "short_url": "short123",
    "original_url": "https://example.com",
    "access_count": 5,
    "created_at": "2025-01-01T12:00:00Z",
    "updated_at": "2025-01-01T12:00:00Z"
  }
  ```

#### `DELETE /api/v1/short/{short}`
* **Description:** Deletes a short URL.
* **Response (JSON):**
  ```json
  {
    "message": "short url deleted"
  }
  ```

#### `PATCH /api/v1/short/{short}`
* **Description:** Updates the original URL for a given short URL.
* **Request Body (JSON):**
  ```json
  {
    "original_url": "https://new-example.com"
  }
  ```
* **Response (JSON):**
  ```json
  {
    "short_url": "short123",
    "original_url": "https://new-example.com"
  }
  ```

### Short URL Redirection

#### `GET /{short}`
* **Description:** Redirects to the original URL associated with the provided `{short}` URL path.
* **Path Parameter:**
  - `short`: The unique short URL identifier.
* **Status Codes:**
  - `301 Permanent Redirect`: If the short URL is found, redirects to the original URL.
  - `400 Bad Request`: If the short path value is empty.
  - `404 Not Found`: If the provided short URL does not exist.
  - `500 Internal Server Error`: If an unexpected error occurs.

## Logging

The project uses a custom JSON-based logging system (`jsonlog`). Logs are structured and provide clear details about application events.

## Error Handling

* Uses structured error responses with appropriate HTTP status codes.
* Common errors include:
  - `400 Bad Request` for invalid inputs.
  - `404 Not Found` for missing resources.
  - `409 Conflict` for existing short URLs.
  - `500 Internal Server Error` for unexpected issues.

## Contributing

1. Fork the repository.
2. Create a feature branch.
3. Commit your changes.
4. Push your branch and create a pull request.

